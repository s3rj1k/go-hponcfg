package hponcfg

import (
	"encoding/xml"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// GetHealthData - process health output of hponcfg
func GetHealthData(data []byte) (GetEmbeddedHealthData, error) {

	var healthData GetEmbeddedHealthData

	// fix XML tags in firmware section for proper unmarshalling
	data = regexp.MustCompile("(?i)<INDEX_[0-9]+>").ReplaceAllLiteral(data, []byte("<INDEX>"))
	data = regexp.MustCompile("(?i)</INDEX_[0-9]+>").ReplaceAllLiteral(data, []byte("</INDEX>"))

	err := xml.Unmarshal(data, &healthData)
	if err != nil {
		return GetEmbeddedHealthData{}, err
	}

	// process Health At A Glance
	if healthData.HealthAtAGlance != nil {
		for _, fan := range healthData.HealthAtAGlance.Fans {
			if len(strings.TrimSpace(fan.Status)) > 0 {
				healthData.HealthAtAGlance.Fan.Status = fan.Status
			}
			if len(strings.TrimSpace(fan.Redundancy)) > 0 {
				healthData.HealthAtAGlance.Fan.Redundancy = fan.Redundancy
			}
		}
		for _, ps := range healthData.HealthAtAGlance.PowerSupplies {
			if len(strings.TrimSpace(ps.Status)) > 0 {
				healthData.HealthAtAGlance.PowerSupply.Status = ps.Status
			}
			if len(strings.TrimSpace(ps.Redundancy)) > 0 {
				healthData.HealthAtAGlance.PowerSupply.Redundancy = ps.Redundancy
			}
		}
	}

	// make CPU names as CPU tags in XML
	if healthData.Processors.Processor != nil {
		for _, cpu := range healthData.Processors.Processor {
			cpu.Label.Value = regexp.MustCompile(`(?i)Proc `).ReplaceAllString(cpu.Label.Value, "CPU_")
		}
	}

	// process fan sensors
	if healthData.Fans.Fan != nil {
		for _, fan := range healthData.Fans.Fan {

			// set bool fan status
			if strings.EqualFold(strings.TrimSpace(fan.Status.Value), "OK") {
				fan.Status.Ok = true
			}

			// set int64 fan speed
			fan.Speed.IntValue, err = strconv.ParseInt(fan.Speed.Value, 10, 64)
			if err != nil {
				fan.Speed.IntValue = 0
			}

			// set default units for fan speed
			if len(strings.TrimSpace(fan.Speed.Unit)) == 0 {
				fan.Speed.Unit = "Percentage"
			}

			// fail if fan speed is 0
			if fan.Speed.IntValue == 0 {
				fan.Status.Ok = false
				fan.Status.Value = "FAIL"
			}

			// fail if fan speed is 100 and units are Percentage
			if strings.EqualFold(strings.TrimSpace(fan.Speed.Unit), "Percentage") {
				if fan.Speed.IntValue == 100 {
					fan.Status.Ok = false
					fan.Status.Value = "FAIL"
				}
			}
		}
	}

	// process temperature sensors
	if healthData.Temperature.Temp != nil {
		for _, sensor := range healthData.Temperature.Temp {

			var err error

			// set bool sensor status
			if strings.EqualFold(strings.TrimSpace(sensor.Status.Value), "OK") {
				sensor.Status.Ok = true
			}

			// set bool installed status
			if strings.EqualFold(strings.TrimSpace(sensor.Status.Value), "Not Installed") {
				sensor.Status.Installed = false
				// skip this sensor if not installed
				continue
			} else {
				sensor.Status.Installed = true
			}

			// set int64 current sensor reading
			sensor.CurrentReading.IntValue, err = strconv.ParseInt(sensor.CurrentReading.Value, 10, 64)
			if err != nil {
				sensor.CurrentReading.IntValue = math.MaxInt64
			}

			// set default units for current sensor reading
			if len(strings.TrimSpace(sensor.CurrentReading.Unit)) == 0 {
				sensor.CurrentReading.Unit = "Celsius"
			}

			// set int64 sensor caution value
			sensor.Caution.IntValue, err = strconv.ParseInt(sensor.Caution.Value, 10, 64)
			if err != nil {
				sensor.Caution.IntValue = math.MaxInt64
			}

			// set default units for caution sensor value
			if len(strings.TrimSpace(sensor.Caution.Unit)) == 0 {
				sensor.Caution.Unit = "Celsius"
			}

			// set int64 sensor caution value
			sensor.Critical.IntValue, err = strconv.ParseInt(sensor.Critical.Value, 10, 64)
			if err != nil {
				sensor.Critical.IntValue = math.MaxInt64
			}

			// set default units for critical sensor value
			if len(strings.TrimSpace(sensor.Critical.Unit)) == 0 {
				sensor.Critical.Unit = "Celsius"
			}

			// set fail status of temperature reading are not ok
			if (sensor.CurrentReading.IntValue >= sensor.Caution.IntValue) ||
				(sensor.CurrentReading.IntValue >= sensor.Critical.IntValue) {
				sensor.Status.Ok = false
				sensor.Status.Value = "FAIL"
			}

		}
	}

	// process power supply summary
	if healthData.PowerSupplies.PowerSupplySummary != nil {
		// set as not available if reading is 0
		if strings.EqualFold(strings.TrimSpace(healthData.PowerSupplies.PowerSupplySummary.PresentPowerReading.Value), "0 Watts") {
			healthData.PowerSupplies.PowerSupplySummary.PresentPowerReading.Available = false
		} else {
			healthData.PowerSupplies.PowerSupplySummary.PresentPowerReading.Available = true
		}
	}

	// process ilo
	if healthData.NICInformation.ILO != nil {
		healthData.NICInformation.ILO.ILO.Value = true
		// add to merged slice
		healthData.NICInformation.Merged = append(healthData.NICInformation.Merged, healthData.NICInformation.ILO)
	}

	// process nics
	if healthData.NICInformation.NICs != nil {
		for _, nic := range healthData.NICInformation.NICs {
			if strings.Contains(strings.ToLower(nic.PortDescription.Value), "ilo") {
				nic.ILO.Value = true
			}

			// add to merged slice
			healthData.NICInformation.Merged = append(healthData.NICInformation.Merged, nic)
		}
	}

	// process memory summary for CPUs
	if healthData.Memory.MemoryDetailsSummary.CPU1 != nil {
		healthData.Memory.MemoryDetailsSummary.IsCPU1 = true
	}
	if healthData.Memory.MemoryDetailsSummary.CPU2 != nil {
		healthData.Memory.MemoryDetailsSummary.IsCPU2 = true
	}
	if healthData.Memory.MemoryDetailsSummary.CPU3 != nil {
		healthData.Memory.MemoryDetailsSummary.IsCPU3 = true
	}
	if healthData.Memory.MemoryDetailsSummary.CPU4 != nil {
		healthData.Memory.MemoryDetailsSummary.IsCPU4 = true
	}

	// process memory info on CPUs
	if healthData.Memory.MemoryDetails.CPU1 != nil {
		for _, mem := range healthData.Memory.MemoryDetails.CPU1 {
			if strings.EqualFold(strings.TrimSpace(mem.Status.Value), "Not Present") {
				mem.Status.Installed = false
			} else {
				mem.Status.Installed = true
				healthData.Memory.MemoryDetails.IsCPU1 = true
			}
		}
	}
	if healthData.Memory.MemoryDetails.CPU2 != nil {
		healthData.Memory.MemoryDetails.IsCPU2 = true
		for _, mem := range healthData.Memory.MemoryDetails.CPU2 {
			if strings.EqualFold(strings.TrimSpace(mem.Status.Value), "Not Present") {
				mem.Status.Installed = false
			} else {
				mem.Status.Installed = true
				healthData.Memory.MemoryDetails.IsCPU2 = true
			}
		}
	}
	if healthData.Memory.MemoryDetails.CPU3 != nil {
		for _, mem := range healthData.Memory.MemoryDetails.CPU3 {
			if strings.EqualFold(strings.TrimSpace(mem.Status.Value), "Not Present") {
				mem.Status.Installed = false
			} else {
				mem.Status.Installed = true
				healthData.Memory.MemoryDetails.IsCPU3 = true
			}
		}
	}
	if healthData.Memory.MemoryDetails.CPU4 != nil {
		for _, mem := range healthData.Memory.MemoryDetails.CPU4 {
			if strings.EqualFold(strings.TrimSpace(mem.Status.Value), "Not Present") {
				mem.Status.Installed = false
			} else {
				mem.Status.Installed = true
				healthData.Memory.MemoryDetails.IsCPU4 = true
			}
		}
	}

	return healthData, nil
}
