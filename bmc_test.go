package hponcfg

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestGetBMCData(t *testing.T) {

	// data to process
	data := []byte(`
		THIS IS TEST GARBAGE
		<GET_FW_VERSION
			FIRMWARE_VERSION = "2.55"
			FIRMWARE_DATE = "Aug 16 2017"
			MANAGEMENT_PROCESSOR = "iLO4"
			LICENSE_TYPE = "iLO Advanced"
		/>
	`)

	data = append(data, byte(0x13), byte(0x14), byte(0xF0), byte(0x9F), byte(0x98), byte(0x82))

	// parse data
	bmc1, err := GetBMCData(data)
	if err != nil {
		t.Errorf("Failed to parse BMC1 data: %s", err.Error())
	}

	// marshal bmc1 data
	xml, err := xml.MarshalIndent(bmc1, "", "\t")
	if err != nil {
		t.Errorf("Failed to marshal BMC1 data: %s", err.Error())
	}

	// parse data
	bmc2, err := GetBMCData(xml)
	if err != nil {
		t.Errorf("Failed to parse BMC2 data: %s", err.Error())
	}

	if !reflect.DeepEqual(bmc1, bmc2) {
		t.Errorf("Unmarshaled BMC data does not equal Re-Unmarshaled, must be equal")
	}

}
