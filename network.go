package hponcfg

import (
	"encoding/xml"
	"regexp"
)

// GetNetworkData - process network output of hponcfg
func GetNetworkData(data []byte) (GetNetworkSettings, error) {

	var networkData GetNetworkSettings

	// replace non-ASCII symbols with space
	data = regexp.MustCompile("[[:^ascii:]]").ReplaceAllLiteral(data, []byte(" "))

	err := xml.Unmarshal(data, &networkData)
	if err != nil {
		return GetNetworkSettings{}, err
	}

	return networkData, nil
}
