package hponcfg

import (
	"encoding/xml"
)

// GetNetworkData - process network output of hponcfg
func GetNetworkData(data []byte) (GetNetworkSettings, error) {

	var networkData GetNetworkSettings

	err := xml.Unmarshal(data, &networkData)
	if err != nil {
		return GetNetworkSettings{}, err
	}

	return networkData, nil
}
