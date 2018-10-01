package hponcfg

import (
	"encoding/xml"
)

// GetNetworkData - process network output of hponcfg
func GetNetworkData(data []byte) (GetNetworkSettings, error) {

	var networkData GetNetworkSettings

	clean, err := SanitizeHPOnCfgXML(data)
	if err != nil {
		return GetNetworkSettings{}, err
	}

	err = xml.Unmarshal(clean, &networkData)
	if err != nil {
		return GetNetworkSettings{}, err
	}

	return networkData, nil
}
