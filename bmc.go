package hponcfg

import (
	"encoding/xml"
)

// GetBMCData - process BMC(GET_FW_VERSION) output of hponcfg
func GetBMCData(data []byte) (GetFWVersion, error) {

	var bmcData GetFWVersion

	clean, err := SanitizeHPOnCfgXML(data)
	if err != nil {
		return GetFWVersion{}, err
	}

	err = xml.Unmarshal(clean, &bmcData)
	if err != nil {
		return GetFWVersion{}, err
	}

	return bmcData, nil
}
