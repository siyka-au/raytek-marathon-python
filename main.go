package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"./raytek"

	"github.com/tarm/serial"
)

type ValueItem struct {
	Item  raytek.DataItem
	Value interface{}
}

func extractData(raw string) (*ValueItem, error) {
	var item raytek.DataItem
	var val interface{}

	if len(raw) < 2 {
		return nil, errors.New("Can't extract data with length less than 2")
	}

	dataItemCode := raw[0:1]
	rawData := raw[1:]

	if raw[0:1] == "X" {
		dataItemCode = raw[0:2]
		rawData = raw[2:]
	}

	item = raytek.DataItem(dataItemCode)

	switch dataItemCode {
	// String values
	case "$", "X$", "XV", "XU":
		// BurstStringFormat              DataItem = "$"
		// BurstStringContents            DataItem = "X$"
		// IdentifyUnit                   DataItem = "XU"
		// Serial                         DataItem = "XV"
		val = rawData

	// Integer values
	case "A", "B", "C", "H", "I", "L", "N", "T", "W", "Y", "Z", "XA", "XB", "XD", "XE", "XP", "XS", "XY":
		// AmbientRadiationCorrection     DataItem = "A"
		// Attenuation                    DataItem = "B"
		// AdvancedHoldThreshold          DataItem = "C"
		// TopOfMilliampRange             DataItem = "H"
		// InternalAmbientTemperature     DataItem = "I"
		// BottomOfMilliampRange          DataItem = "L"
		// TemperatureNarrow              DataItem = "N"
		// Temperature                    DataItem = "T"
		// TemperatureWide                DataItem = "W"
		// MultidropAddress               DataItem = "XA"
		// LowTemperatureLimit            DataItem = "XB"
		// Deadband                       DataItem = "XD"
		// DecayRate                      DataItem = "XE"
		// HighTemperatureLimit           DataItem = "XH"
		// SecondSetpoint                 DataItem = "XP"
		// SetpointRelayFunction          DataItem = "XS"
		// HysteresisAdvancedHold         DataItem = "XY"
		// AttentuationForRelayActivation DataItem = "Y"
		// AttentuationForFailsafe        DataItem = "Z"
		val, _ = strconv.Atoi(rawData)

	// Float values
	case "E", "F", "G", "P", "Q", "R", "S":
		// Emissivity                     DataItem = "E"
		// ValleyHoldTime                 DataItem = "F"
		// AverageTime                    DataItem = "G"
		// PeakHoldTime                   DataItem = "P"
		// PowerWide                      DataItem = "Q"
		// PowerNarrow                    DataItem = "R"
		// Slope                          DataItem = "S"
		val, _ = strconv.ParseFloat(rawData, 64)

	case "D": // BaudRate                       DataItem = "D"
		baudRate, _ := strconv.Atoi(rawData)
		baudRate = baudRate * 100
		val = baudRate

	case "J": // SwitchPanelLock                DataItem = "J"
		val = raytek.PanelLockState(rawData[0])

	case "M":
		// Mode                           DataItem = "M"
		mode, _ := strconv.Atoi(rawData)
		val = raytek.PyrometerMode(mode)

	case "U":
		// TemperatureUnits               DataItem = "U"
		val = raytek.TemperatureUnit(rawData[0])

	// case "XI":
	// 	// Initialisation                 DataItem = "XI"
	// 	return ValueItem{raytek.Initialisation, raw[2:]}

	case "XL":
		// Laser                          DataItem = "XL"
		val = raytek.LaserStatus(rawData[0])

	case "XM":
		// ModelType                      DataItem = "XM"
		val = raytek.SensorModel(rawData[0])

	case "XO":
		// OutputCurrentType              DataItem = "XO"
		outputCurrentRange, _ := strconv.Atoi(rawData)
		val = raytek.OutputCurrentRange(outputCurrentRange)

	// case "XR":
	// 	// Revision                       DataItem = "XR"
	// 	return ValueItem{raytek.Revision, raw[2:]}

	// case "XT":
	// 	// Trigger                        DataItem = "XT"

	// 	return ValueItem{raytek.Trigger, raw[2:]}

	default:
		return nil, errors.New("Couldn't parse the data item " + dataItemCode)
	}

	return &ValueItem{item, val}, nil
}

func readLoop(rw *bufio.ReadWriter) (<-chan ValueItem, <-chan ValueItem, <-chan ValueItem, <-chan error) {
	responses := make(chan ValueItem, 1)
	burstData := make(chan ValueItem, 1)
	events := make(chan ValueItem, 1)
	errors := make(chan error, 1)

	go func() {
		defer close(responses)
		defer close(burstData)
		defer close(events)
		defer close(errors)

		for {
			r, err := rw.ReadBytes('\n')
			if err != nil {
				errors <- err
			}
			s := string(bytes.Trim(r, " \r\n"))
			if len(s) > 0 {
				dataTypeCode := s[0]
				data := strings.TrimSpace(s[1:])
				switch dataTypeCode {
				// A response was received
				case '!':
					vi, err := extractData(data)
					if err != nil {
						errors <- err
					} else {
						responses <- *vi
					}

				// An event was received
				case '#':
					vi, err := extractData(data)
					if err != nil {
						errors <- err
					} else {
						events <- *vi
					}

					// If it starts with a letter, assume we have burst data, otherwise assume we have partial data and skip
				default:
					if strings.ContainsAny(s[0:1], "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
						parts := strings.Split(s, " ")
						for _, v := range parts {
							vi, err := extractData(v)
							if err != nil {
								errors <- err
							} else {
								burstData <- *vi
							}
						}
					}
				}
			}
		}
	}()

	return responses, burstData, events, errors
}

func merge(cs ...<-chan ValueItem) <-chan ValueItem {
	out := make(chan ValueItem)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan ValueItem) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func __write(rw *bufio.ReadWriter, command string) error {
	_, err := rw.Write([]byte(command + "\r"))
	if err != nil {
		return err
	}

	return rw.Flush()
}

func query(rw *bufio.ReadWriter, di raytek.DataItem) error {
	command := fmt.Sprintf("?%s", di)
	return __write(rw, command)
}

func set(rw *bufio.ReadWriter, di raytek.DataItem, val string) error {
	command := fmt.Sprintf("%s=%s", di, val)
	return __write(rw, command)
}

func main() {

	serialPortPtr := flag.String("serial-port", "/dev/ttyUSB0", "serial port device")
	baudRatePtr := flag.Uint("baud-rate", 38400, "baud rate")
	piControlPtr := flag.String("pi-control", "/dev/piControl0", "RevPi piControl device")
	readAddressPtr := flag.Uint("read-addr-base", 11, "base address for reading")
	writeAddressPtr := flag.Uint("write-addr-base", 523, "base address for writing")

	readAddressPtr = readAddressPtr

	flag.Parse()

	c := &serial.Config{Name: *serialPortPtr, Baud: int(*baudRatePtr)}
	ser, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	piCtl, err := os.OpenFile(*piControlPtr, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	rw := bufio.NewReadWriter(bufio.NewReader(ser), bufio.NewWriter(ser))

	responses, burstData, events, _ := readLoop(rw)

	dataUpdates := merge(responses, burstData, events)

	go func() {
		var offset int64 = int64(*writeAddressPtr)

		memoryMap := map[raytek.DataItem]int64{
			raytek.Temperature:       0,
			raytek.TemperatureWide:   2,
			raytek.TemperatureNarrow: 4,
			raytek.PowerWide:         6,
			raytek.PowerNarrow:       14,
			raytek.Emissivity:        22,
			raytek.Slope:             30,
			raytek.Attenuation:       38,
		}

		for {
			var buf []byte

			r := <-dataUpdates
			piCtl.Seek(offset+memoryMap[r.Item], os.SEEK_SET)
			fmt.Println(r)
			switch r.Item {
			case raytek.Temperature, raytek.TemperatureWide, raytek.TemperatureNarrow, raytek.Attenuation:
				intVal := uint16(r.Value.(int))
				buf = make([]byte, 2)
				binary.LittleEndian.PutUint16(buf[:], intVal)

			case raytek.Emissivity, raytek.Slope, raytek.PowerWide, raytek.PowerNarrow:
				floatVal := r.Value.(float64)
				buf = make([]byte, 8)
				binary.LittleEndian.PutUint64(buf[:], math.Float64bits(floatVal))
			}
			piCtl.Write(buf)
		}
	}()

	// set(rw, raytek.BurstStringFormat,
	// 	string(raytek.Temperature)+
	// 		string(raytek.TemperatureNarrow)+
	// 		string(raytek.TemperatureWide)+
	// 		string(raytek.PowerNarrow)+
	// 		string(raytek.PowerWide)+
	// 		string(raytek.Attenuation)+
	// 		string(raytek.InternalAmbientTemperature))
	// query(rw, raytek.Emissivity)
	// query(rw, raytek.Slope)
	set(rw, raytek.PollBurstMode, string(raytek.BurstDataMode))

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	fmt.Println("Shutting down...")
	os.Exit(1)
}
