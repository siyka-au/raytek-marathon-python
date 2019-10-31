package raytek

type LaserStatus byte

const (
	LaserOff      LaserStatus = '0'
	LaserOn       LaserStatus = '1'
	LaserOverheat LaserStatus = 'H'
)

type SensorModel byte

const (
	SensorModelA SensorModel = 'A'
	SensorModelB SensorModel = 'B'
	SensorModelC SensorModel = 'C'
)

type OutputCurrentMode int

const (
	OutputCurrentControlledByUnit   OutputCurrentMode = 1
	OutputCurrentControlledManually OutputCurrentMode = 2
)

type OutputCurrentRange int

const (
	OutputCurrentRange0To20 OutputCurrentRange = 0
	OutputCurrentRange4To20 OutputCurrentRange = 4
)

type RelayAlarmOutputControl int

const (
	RelayAlarmOutputOff            RelayAlarmOutputControl = 0
	RelayAlarmOutputOn             RelayAlarmOutputControl = 1
	RelayAlarmOutputNormallyOpen   RelayAlarmOutputControl = 2
	RelayAlarmOutputNormallyClosed RelayAlarmOutputControl = 3
)

type PyrometerMode int

const (
	SingleColourMode PyrometerMode = 1
	TwoColourMode    PyrometerMode = 2
	SingleColorMode  PyrometerMode = 1
	TwoColorMode     PyrometerMode = 2
)

type TemperatureUnit byte

const (
	Celcius    TemperatureUnit = 'C'
	Fahrenheit TemperatureUnit = 'F'
)

type DataMode byte

const (
	PolledDataMode DataMode = 'P'
	BurstDataMode  DataMode = 'B'
)

type PanelLockState byte

const (
	PanelLocked   PanelLockState = 'L'
	PanelUnlocked PanelLockState = 'U'
)

type TriggerState string

const (
	TriggerInactive TriggerState = "XT0"
	TriggerActive   TriggerState = "XT1"
)

type DataItem string

const (
	BurstStringFormat              DataItem = "$"
	AmbientRadiationCorrection     DataItem = "A"
	Attenuation                    DataItem = "B"
	AdvancedHoldThreshold          DataItem = "C"
	BaudRate                       DataItem = "D"
	Emissivity                     DataItem = "E"
	ValleyHoldTime                 DataItem = "F"
	AverageTime                    DataItem = "G"
	TopOfMilliampRange             DataItem = "H"
	InternalAmbientTemperature     DataItem = "I"
	SwitchPanelLock                DataItem = "J"
	RelayAlarmOutput               DataItem = "K"
	BottomOfMilliampRange          DataItem = "L"
	Mode                           DataItem = "M"
	TemperatureNarrow              DataItem = "N"
	OutputCurrent                  DataItem = "O"
	PeakHoldTime                   DataItem = "P"
	PowerWide                      DataItem = "Q"
	PowerNarrow                    DataItem = "R"
	Slope                          DataItem = "S"
	Temperature                    DataItem = "T"
	TemperatureUnits               DataItem = "U"
	PollBurstMode                  DataItem = "V"
	TemperatureWide                DataItem = "W"
	BurstStringContents            DataItem = "X$"
	MultidropAddress               DataItem = "XA"
	LowTemperatureLimit            DataItem = "XB"
	Deadband                       DataItem = "XD"
	DecayRate                      DataItem = "XE"
	RestoreFactoryDefaults         DataItem = "XF"
	HighTemperatureLimit           DataItem = "XH"
	Initialisation                 DataItem = "XI"
	Initialization                 DataItem = "XI"
	Laser                          DataItem = "XL"
	ModelType                      DataItem = "XM"
	OutputCurrentType              DataItem = "XO"
	SecondSetpoint                 DataItem = "XP"
	Revision                       DataItem = "XR"
	SetpointRelayFunction          DataItem = "XS"
	Trigger                        DataItem = "XT"
	IdentifyUnit                   DataItem = "XU"
	Serial                         DataItem = "XV"
	HysteresisAdvancedHold         DataItem = "XY"
	AttentuationForRelayActivation DataItem = "Y"
	AttentuationForFailsafe        DataItem = "Z"
	Unknown                        DataItem = ""
)

var validBurstItems = [...]DataItem{
	Attenuation,
	Emissivity,
	AverageTime,
	TopOfMilliampRange,
	InternalAmbientTemperature,
	BottomOfMilliampRange,
	Mode,
	TemperatureNarrow,
	OutputCurrent,
	PeakHoldTime,
	PowerWide,
	PowerNarrow,
	Slope,
	Temperature,
	TemperatureUnits,
	TemperatureWide,
	MultidropAddress,
	Initialisation,
	Initialization,
	Trigger,
	AttentuationForRelayActivation,
	AttentuationForFailsafe,
}
