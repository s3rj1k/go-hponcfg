package hponcfg

import (
	"encoding/xml"
	"regexp"
)

// GetBMCData - process BMC(GET_FW_VERSION) output of hponcfg
func GetBMCData(data []byte) (GetFWVersion, error) {

	var bmcData GetFWVersion

	// replace non-ASCII characters with space
	data = regexp.MustCompile("[[:^ascii:]]").ReplaceAllLiteral(data, []byte(" "))

	// replace control characters with space
	data = regexp.MustCompile("[[:cntrl:]]").ReplaceAllLiteral(data, []byte(" "))

	err := xml.Unmarshal(data, &bmcData)
	if err != nil {
		return GetFWVersion{}, err
	}

	return bmcData, nil
}
