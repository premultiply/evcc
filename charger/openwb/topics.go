package openwb

// predefined openWB topic names
const (
	// alive
	TimestampTopic = "Timestamp"

	// status
	PluggedTopic    = "boolPlugStat"
	ChargingTopic   = "boolChargeStat"
	ConfiguredTopic = "boolChargePointConfigured"

	// getter/setter
	EnabledTopic    = "ChargePointEnabled"
	MaxCurrentTopic = "DirectChargeAmps"

	// charge power
	ChargePowerTopic       = "W"
	ChargeTotalEnergyTopic = "kWhCounter"

	// general measurements
	PowerTopic   = "W"
	SoCTopic     = "%Soc"
	CurrentTopic = "APhase" // 1..3

	// configuration
	PvConfigured      = "boolPVConfigured"
	BatteryConfigured = "boolHouseBatteryConfigured"
)
