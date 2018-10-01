package hponcfg

import (
	"encoding/xml"
	"reflect"
	"regexp"
	"testing"
)

func TestGetNetworkData(t *testing.T) {

	// data to process
	data := []byte(`
		<GET_NETWORK_SETTINGS>
			<ENABLE_NIC VALUE="Y"/>
			<SHARED_NETWORK_PORT VALUE="LOM"/>
			<VLAN_ENABLED VALUE="Y"/>
			<VLAN_ID VALUE="88"/>
			<SPEED_AUTOSELECT VALUE="Y"/>
			<NIC_SPEED VALUE="Automatic"/>
			<FULL_DUPLEX VALUE="Automatic"/>
			<DHCP_ENABLE VALUE="N"/>
			<DHCP_GATEWAY VALUE="N"/>
			<DHCP_DNS_SERVER VALUE="N"/>
			<DHCP_WINS_SERVER VALUE="N"/>
			<DHCP_STATIC_ROUTE VALUE="N"/>
			<DHCP_DOMAIN_NAME VALUE="N"/>
			<DHCP_SNTP_SETTINGS VALUE="N"/>
			<REG_WINS_SERVER VALUE="N"/>
			<REG_DDNS_SERVER VALUE="N"/>
			<PING_GATEWAY VALUE="N"/>
			<MAC_ADDRESS VALUE="d0:67:26:9c:66:d3"/>
			<IP_ADDRESS VALUE="192.168.92.18"/>
			<SUBNET_MASK VALUE="255.255.252.0"/>
			<GATEWAY_IP_ADDRESS VALUE="192.168.92.1"/>
			<DNS_NAME VALUE="ILOCZ8813099V"/>
			<DOMAIN_NAME VALUE="local.net"/>
			<PRIM_DNS_SERVER VALUE="88.55.44.33"/>
			<SEC_DNS_SERVER VALUE="0.0.0.0"/>
			<TER_DNS_SERVER VALUE="0.0.0.0"/>
			<PRIM_WINS_SERVER VALUE="0.0.0.0"/>
			<SEC_WINS_SERVER VALUE="0.0.0.0"/>
			<SNTP_SERVER1 VALUE="17.77.77.77"/>
			<SNTP_SERVER2 VALUE=""/>
			<TIMEZONE VALUE="Europe/Kiev"/>
			<STATIC_ROUTE_1
				DEST="0.0.0.0"
				MASK="0.0.0.0"
				GATEWAY="0.0.0.0"
			/>
			<STATIC_ROUTE_2
				DEST="0.0.0.0"
				MASK="0.0.0.0"
				GATEWAY="0.0.0.0"
			/>
			<STATIC_ROUTE_3
				DEST="0.0.0.0"
				MASK="0.0.0.0"
				GATEWAY="0.0.0.0"
			/>
			<IPV6_ADDRESS
				VALUE="fe80::9e8e:99ff:fe30:a550"
				PREFIXLEN="64"
				ADDR_SOURCE="SLAAC"
				ADDR_STATUS="ACTIVE"
			/>
			<IPV6_ADDRESS
				VALUE="fd0e:2d:ba9:221a:9e8e:99ff:fe30:a550"
				PREFIXLEN="64"
				ADDR_SOURCE="SLAAC"
				ADDR_STATUS="ACTIVE"
			/>
			<IPV6_STATIC_ROUTE_1
				IPV6_DEST="::"
				PREFIXLEN="0"
				IPV6_GATEWAY="::"
				ADDR_STATUS="INACTIVE"
			/>
			<IPV6_STATIC_ROUTE_2
				IPV6_DEST="::"
				PREFIXLEN="0"
				IPV6_GATEWAY="::"
				ADDR_STATUS="INACTIVE"
			/>
			<IPV6_STATIC_ROUTE_3
				IPV6_DEST="::"
				PREFIXLEN="0"
				IPV6_GATEWAY="::"
				ADDR_STATUS="INACTIVE"
			/>
			<IPV6_PRIM_DNS_SERVER VALUE="::"/>
			<IPV6_SEC_DNS_SERVER VALUE="::"/>
			<IPV6_TER_DNS_SERVER VALUE="::"/>
			<IPV6_DEFAULT_GATEWAY VALUE="::"/>
			<IPV6_PREFERRED_PROTOCOL VALUE="Y"/>
			<IPV6_ADDR_AUTOCFG VALUE="Y"/>
			<IPV6_REG_DDNS_SERVER VALUE="Y"/>
			<DHCPV6_STATELESS_ENABLE VALUE="Y"/>
			<DHCPV6_STATEFUL_ENABLE VALUE="Y"/>
			<DHCPV6_RAPID_COMMIT VALUE="N"/>
			<DHCPV6_DOMAIN_NAME VALUE="N"/>
			<DHCPV6_SNTP_SETTINGS VALUE="N"/>
			<DHCPV6_DNS_SERVER VALUE="Y"/>
			<ILO_NIC_AUTO_SELECT VALUE="DISABLED"/>
			<ILO_NIC_AUTO_SNP_SCAN VALUE="0"/>
			<ILO_NIC_AUTO_DELAY VALUE="90"/>
			<ILO_NIC_FAIL_OVER VALUE="DISABLED"/>
			<ILO_NIC_FAIL_OVER_DELAY VALUE="300"/>
			<SNP_PORT VALUE="1"/>
		</GET_NETWORK_SETTINGS>
	`)

	data = append(data, byte(0x13), byte(0x14), byte(0xF0), byte(0x9F), byte(0x98), byte(0x82))

	// parse data
	network1, err := GetNetworkData(data)
	if err != nil {
		t.Errorf("Failed to parse network1 data: %s", err.Error())
	}

	// marshal network1 data
	xml, err := xml.MarshalIndent(network1, "", "\t")
	if err != nil {
		t.Errorf("Failed to marshal network1 data: %s", err.Error())
	}

	// fix self-closing tags
	selfClosingXML := regexp.MustCompile("></([a-zA-Z0-9-_]+)>").ReplaceAll(xml, []byte("/>"))

	// parse data
	network2, err := GetNetworkData(selfClosingXML)
	if err != nil {
		t.Errorf("Failed to parse network2 data: %s", err.Error())
	}

	if !reflect.DeepEqual(network1, network2) {
		t.Errorf("Unmarshaled network data does not equal Re-Unmarshaled, must be equal")
	}

}
