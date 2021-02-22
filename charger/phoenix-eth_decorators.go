package charger

// Code generated by github.com/andig/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/andig/evcc/api"
)

func decoratePhoenixEth(base *PhoenixEth, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil:
		return base

	case meter != nil && meterCurrent == nil && meterEnergy == nil:
		return &struct {
			*PhoenixEth
			api.Meter
		}{
			PhoenixEth: base,
			Meter: &decoratePhoenixEthMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*PhoenixEth
			api.MeterEnergy
		}{
			PhoenixEth: base,
			MeterEnergy: &decoratePhoenixEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*PhoenixEth
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEth: base,
			Meter: &decoratePhoenixEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*PhoenixEth
			api.MeterCurrent
		}{
			PhoenixEth: base,
			MeterCurrent: &decoratePhoenixEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*PhoenixEth
			api.Meter
			api.MeterCurrent
		}{
			PhoenixEth: base,
			Meter: &decoratePhoenixEthMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*PhoenixEth
			api.MeterCurrent
			api.MeterEnergy
		}{
			PhoenixEth: base,
			MeterCurrent: &decoratePhoenixEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*PhoenixEth
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			PhoenixEth: base,
			Meter: &decoratePhoenixEthMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratePhoenixEthMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEthMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEthMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEthMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decoratePhoenixEthMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixEthMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
