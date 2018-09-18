package hponcfg

import (
	"encoding/xml"
)

// GetBMCData - process BMC(GET_FW_VERSION) output of hponcfg
func GetBMCData(data []byte) (GetFWVersion, error) {

	var bmcData GetFWVersion

	err := xml.Unmarshal(data, &bmcData)
	if err != nil {
		return GetFWVersion{}, err
	}

	return bmcData, nil
}
