package raytek

import (
	"bufio"
	"fmt"

	"github.com/robaho/fixed"
)

type Marathon struct {
	Temperature       fixed.Fixed
	TemperatureWide   float64
	TemperatureNarrow float64

	Emissivity float64
	Slope      float64

	rw *bufio.ReadWriter
}

func NewMarathon(rw *bufio.ReadWriter) {

}

func __write(rw *bufio.ReadWriter, command string) error {
	_, err := rw.Write([]byte(command + "\r"))
	if err != nil {
		return err
	}

	return rw.Flush()
}

func (m *Marathon) query(rw *bufio.ReadWriter, dataItem DataItem) error {
	command := fmt.Sprintf("?%s", di)
	return __write(rw, command)
}

func (m *Marathon) set(rw *bufio.ReadWriter, di dataItem, val string) error {
	command := fmt.Sprintf("%s=%s", di, val)
	return __write(rw, command)
}

// func (p RaytekMarathon)

// 	def __init__(self, reader, writer):
// 		self._dataMode: DataMode = DataMode.POLL

// 		self.r = reader
// 		self.w = writer
// 		self.f = None
// 		self.txlock = asyncio.Lock()
// 		asyncio.ensure_future(self.__process())
// 		#asyncio.new_event_loop().run_until_complete(self.__init())

// 	async def __init(self):
// 		await self.__send(b'X=TWNQRBI')
// 		await self.__send(b'?E')
// 		await self.__send(b'V=P')

// 	@property
// 	async def attenuation(self) -> float:
// 		return await self.__queryFloat(DataItem.ATTENUATION)

// 	@property
// 	async def ambientRadiationCorrection(self) -> int:
// 		return await self.__queryFloat(DataItem.ATTENUATION)

// 	@ambientRadiationCorrection.setter
// 	async def ambientRadiationCorrection(self, x):
// 		pass

// 	@property
// 	async def advancedHoldThreshold(self) -> int:
// 		return await self.__queryFloat(DataItem.ADVANCED_HOLD_THRESHOLD)

// 	@advancedHoldThreshold.setter
// 	async def advancedHoldThreshold(self, x):
// 		pass

// 	@property
// 	async def baudRate(self):
// 		pass

// 	@baudRate.setter
// 	async def baudRate(self, br: int):
// 		pass

// 	@property
// 	async def emissivity(self) -> float:
// 		return await self.__queryFloat(DataItem.EMISSIVITY)

// 	@emissivity.setter
// 	async def emissivity(self, x: float):
// 		pass

// 	@property
// 	async def valleyHoldTime(self) -> float:
// 		return await self.__queryFloat(DataItem.VALLEY_HOLD_TIME)

// 	@valleyHoldTime.setter
// 	async def valleyHoldTime(self, x: float):
// 		pass

// 	@property
// 	async def averageTime(self) -> float:
// 		return await self.__queryFloat(DataItem.AVERAGE_TIME)

// 	@averageTime.setter
// 	async def averageTime(self, x: float):
// 		pass

// 	@property
// 	async def topOfMaRange(self) -> float:
// 		return await self.__queryFloat(DataItem.TOP_OF_MA_RANGE)

// 	@topOfMaRange.setter
// 	async def topOfMaRange(self, x: float):
// 		pass

// 	@property
// 	async def internalAmbientTemperature(self) -> float:
// 		return await self.__queryFloat(DataItem.INTERNAL_AMBIENT_TEMPERATURE)

// 	@property
// 	async def panelLock(self) -> PanelLockState:
// 		return await self.__queryEnum(DataItem.SWITCH_PANEL_LOCK, PanelLockState)

// 	@panelLock.setter
// 	async def panelLock(self, state: PanelLockState):
// 		pass

// 	@property
// 	async def relayAlarmOutputControl(self):
// 		pass

// 	@relayAlarmOutputControl.setter
// 	async def relayAlarmOutputControl(self, x: RelayAlarmOutputControl):
// 		# return await self.__queryEnum(DataItem.RELAY_ALARM_OUTPUT_CONTROL, RelayAlarmOutputControl)
// 		pass

// 	@property
// 	async def bottomOfMaRange(self) -> float:
// 		return await self.__queryFloat(DataItem.BOTTOM_OF_MA_RANGE)

// 	@bottomOfMaRange.setter
// 	async def bottomOfMaRange(self, x: float):
// 		pass

// 	@property
// 	async def mode(self) -> PyrometerMode:
// 		return await self.__queryEnum(DataItem.MODE, PyrometerMode)

// 	@mode.setter
// 	def mode(self, mode: PyrometerMode):
// 		pass

// 	@property
// 	async def temperatureNarrow(self) -> float:
// 		return await self.__queryFloat(DataItem.TEMPERATURE_NARROW)

// 	@property
// 	async def outputCurrent(self) -> float:
// 		pass

// 	@outputCurrent.setter
// 	def outputCurrent(self, mode: OutputCurrentMode):
// 		pass

// 	@property
// 	async def peakHoldTime(self) -> float:
// 		return await self.__queryFloat(DataItem.PEAK_HOLD_TIME)

// 	@peakHoldTime.setter
// 	async def peakHoldTime(self, x: float):
// 		pass

// 	@property
// 	async def powerWide(self) -> float:
// 		return await self.__queryFloat(DataItem.POWER_WIDE)

// 	@property
// 	async def powerNarrow(self) -> float:
// 		return await self.__queryFloat(DataItem.POWER_NARROW)

// 	@property
// 	async def slope(self) -> float:
// 		return await self.__queryFloat(DataItem.SLOPE)

// 	@slope.setter
// 	async def slope(self, x: float):
// 		pass

// 	@property
// 	async def temperature(self) -> float:
// 		return await self.__queryFloat(DataItem.TEMPERATURE)

// 	@property
// 	async def temperatureUnit(self) -> TemperatureUnit:
// 		return await self.__queryEnum(DataItem.TEMPERATURE_UNIT, TemperatureUnit)

// 	@temperatureUnit.setter
// 	async def temperatureUnit(self, x: TemperatureUnit):
// 		pass

// 	@property
// 	async def dataMode(self) -> DataMode:
// 		return self._dataMode

// 	@dataMode.setter
// 	async def dataMode(self, x: DataMode):
// 		pass

// 	@property
// 	async def temperatureWide(self) -> float:
// 		return await self.__queryFloat(DataItem.TEMPERATURE_WIDE)

// 	@property
// 	async def burstData(self):
// 		return await self.__queryFloat(DataItem.BURST_STRING_CONTENTS)

// 	@property
// 	async def multidropAddress(self) -> int:
// 		return await self.__queryInt(DataItem.MULTIDROP_ADDRESS)

// 	@multidropAddress.setter
// 	async def multidropAddress(self, addr: int):
// 		pass

// 	@property
// 	async def lowTemperatureLimit(self) -> float:
// 		return await self.__queryInt(DataItem.LOW_TEMPERATURE_LIMIT)

// 	@property
// 	async def deadband(self) -> float:
// 		return float(await self.__queryInt(DataItem.DEADBAND))

// 	@deadband.setter
// 	async def deadband(self, addr: float):
// 		pass

// 	@property
// 	async def decayRate(self) -> float:
// 		return float(await self.__queryInt(DataItem.DECAY_RATE))

// 	@decayRate.setter
// 	async def decayRate(self, addr: float):
// 		pass

// 	async def restoreFactoryDefaults(self):
// 		pass

// 	@property
// 	async def highTemperatureLimit(self) -> float:
// 		return float(await self.__queryInt(DataItem.HIGH_TEMPERATURE_LIMIT))

// 	@property
// 	async def sensorInitialisation(self) -> float:
// 		return float(await self.__queryInt(DataItem.DEADBAND))

// 	async def sensorInitialise(self, addr: float):
// 		pass

// 	async def laserOn(self):
// 		pass

// 	async def laserOff(self):
// 		pass

// 	@property
// 	def laserStatus(self) -> LaserStatus:
// 		pass

// 	@property
// 	async def sensorModel(self) -> SensorModel:
// 		return await self.__queryString(DataItem.SERIAL)

// 	@property
// 	async def outputCurrentRange(self) -> OutputCurrentRange:
// 		return await self.__queryInt(DataItem.OUTPUT_CURRENT_MODE)

// 	@outputCurrentRange.setter
// 	async def outputCurrentRange(self, range: OutputCurrentRange):
// 		pass

// 	@property
// 	async def secondSetPoint(self) -> float:
// 		return await self.__queryFloat(DataItem.SECOND_SETPOINT)

// 	@secondSetPoint.setter
// 	async def secondSetPoint(self, setPoint: float):
// 		pass

// 	@property
// 	async def sensorRevision(self) -> str:
// 		return await self.__queryString(DataItem.REVISION)

// 	@property
// 	async def setPoint(self) -> float:
// 		return await self.__queryFloat(DataItem.SETPOINT_RELAY_FUNCTION)

// 	@setPoint.setter
// 	async def setPoint(self, setPoint: float):
// 		pass

// 	@property
// 	async def trigger(self) -> float:
// 		return await self.__queryFloat(DataItem.SETPOINT_RELAY_FUNCTION)

// 	@trigger.setter
// 	async def trigger(self, setPoint: float):
// 		pass

// 	@property
// 	async def identity(self) -> str:
// 		return await self.__queryString(DataItem.IDENTIFY_UNIT)

// 	@property
// 	async def serial(self) -> str:
// 		return await self.__queryString(DataItem.SERIAL)

// 	@property
// 	async def hysteresisAdvancedHold(self) -> float:
// 		return await self.__queryFloat(DataItem.HYSTERESIS_ADVANCED_HOLD)

// 	@hysteresisAdvancedHold.setter
// 	async def hysteresisAdvancedHold(self, setPoint: float):
// 		pass

// 	@property
// 	async def attentuationToActivateRelay(self) -> int:
// 		return await self.__queryInt(DataItem.ATTENUATION_FOR_RELAY_ACTIVATION)

// 	@attentuationToActivateRelay.setter
// 	async def attentuationToActivateRelay(self, setPoint: int):
// 		pass

// 	@property
// 	async def attentuationForFailsafe(self) -> int:
// 		return await self.__queryInt(DataItem.ATTENTUATION_FOR_FAILSAFE)

// 	@attentuationForFailsafe.setter
// 	async def attentuationForFailsafe(self, setPoint: int):
// 		pass

// 	def events(observer):
// 		self.observer = observer

// 	def __isFloat(self, x):
// 		try:
// 			float(x)
// 			return True
// 		except ValueError:
// 			return False

// 	async def __queryString(self, dataItem: DataItem) -> str:
// 		s = await self.__query(dataItem)
// 		_, data = self.__parse(s, lambda x: x.decode('ascii'))
// 		return data

// 	async def __queryInt(self, dataItem: DataItem) -> int:
// 		s = await self.__query(dataItem)
// 		_, data = self.__parse(s, lambda x: int(x))
// 		return data

// 	async def __queryFloat(self, dataItem: DataItem) -> float:
// 		s = await self.__query(dataItem)
// 		_, data = self.__parse(s, lambda x: float(x) if self.__isFloat(x) else float('nan'))
// 		return data

// 	async def __queryEnum(self, dataItem: DataItem, enumClass):
// 		s = await self.__query(dataItem)
// 		_, data = self.__parse(s, lambda x: enumClass(x))
// 		return data

// 	async def __query(self, dataItem: DataItem):
// 		return await self.__send(b'?' + dataItem.value)

// 	async def setParameter(self, dataItem: DataItem, value):
// 		return await self.__send(dataItem.value + b'=' + value)

// 	def __parse(self, frag: bytes, dataConverter):
// 		if frag[0:1] == b'X':
// 			id = frag[0:2]
// 			rawData = frag[2:]
// 		else:
// 			id = frag[0:1]
// 			rawData = frag[1:]
// 		data = dataConverter(rawData)
// 		return id, data

// 	async def __send(self, cmd):
// 		async with self.txlock:
// 			self.f = asyncio.get_running_loop().create_future()
// 			cmd = cmd.rstrip() + b'\r'
// 			self.w.write(cmd)
// 			await self.f
// 			return self.f.result()

// 	async def __process(self):
// 		while True:
// 			# Perform reads
// 			line = await self.r.readuntil(b'\n')
// 			line = line.rstrip()
// 			if len(line) > 0:
// 				if line[0] == ord('!'):
// 					# Response
// 					if self.f != None and self.f.done() == False:
// 						self.f.set_result(line[1:])
// 				elif line[0] == ord('#'):
// 					# Event
// 					print(f'EVENT: {line}')
// 				elif line[0] == ord('*'):
// 					# Error
// 					print(f'ERROR: {line}')
// 					if self.f != None and self.f.done() == False:
// 						self.f.set_exception(line[1:])
// 				else:
// 					# Could be burst data or it could be corrupt data due to communication parameter changes

// 					print(f'BURST: {line}')
