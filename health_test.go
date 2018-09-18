package hponcfg

import (
	"encoding/xml"
	"reflect"
	"regexp"
	"testing"
)

func TestGetHealthData(t *testing.T) {

	// data to process
	data := []byte(`
		<GET_EMBEDDED_HEALTH_DATA>
			<PROCESSORS>
				<PROCESSOR>
					<LABEL VALUE = "Proc 1"/>
					<NAME VALUE = "Intel(R) Xeon(R) Silver 4114 CPU @ 2.20GHz"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "2200 MHz"/>
					<EXECUTION_TECHNOLOGY VALUE = "10/10 cores; 20 threads"/>
					<MEMORY_TECHNOLOGY VALUE = "64-bit Capable"/>
					<INTERNAL_L1_CACHE VALUE = "640 KB"/>
					<INTERNAL_L2_CACHE VALUE = "10240 KB"/>
					<INTERNAL_L3_CACHE VALUE = "14080 KB"/>
				</PROCESSOR>
				<PROCESSOR>
					<LABEL VALUE = "Proc 2"/>
					<NAME VALUE = "Intel(R) Xeon(R) Silver 4114 CPU @ 2.20GHz"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "2200 MHz"/>
					<EXECUTION_TECHNOLOGY VALUE = "10/10 cores; 20 threads"/>
					<MEMORY_TECHNOLOGY VALUE = "64-bit Capable"/>
					<INTERNAL_L1_CACHE VALUE = "640 KB"/>
					<INTERNAL_L2_CACHE VALUE = "10240 KB"/>
					<INTERNAL_L3_CACHE VALUE = "14080 KB"/>
				</PROCESSOR>
			</PROCESSORS>
			<MEMORY>
				<ADVANCED_MEMORY_PROTECTION>
					<AMP_MODE_STATUS VALUE = "Advanced ECC"/>
					<CONFIGURED_AMP_MODE VALUE = "Advanced ECC"/>
					<AVAILABLE_AMP_MODES VALUE = "Advanced ECC, Online Spare (Rank Sparing), Intrasocket Mirroring"/>
				</ADVANCED_MEMORY_PROTECTION>
				<MEMORY_DETAILS_SUMMARY>
					<CPU_1>
						<NUMBER_OF_SOCKETS VALUE = "12"/>
						<TOTAL_MEMORY_SIZE VALUE = "16 GB"/>
						<OPERATING_FREQUENCY VALUE = "2400 MHz"/>
						<OPERATING_VOLTAGE VALUE = "1.20 v"/>
					</CPU_1>
					<CPU_2>
						<NUMBER_OF_SOCKETS VALUE = "12"/>
						<TOTAL_MEMORY_SIZE VALUE = "16 GB"/>
						<OPERATING_FREQUENCY VALUE = "2400 MHz"/>
						<OPERATING_VOLTAGE VALUE = "1.20 v"/>
					</CPU_2>
				</MEMORY_DETAILS_SUMMARY>
				<MEMORY_DETAILS>
					<CPU_1>
						<SOCKET VALUE = "1"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "2"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "3"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "4"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "5"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "6"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "7"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "8"/>
						<STATUS VALUE = "Good, In Use"/>
						<HP_SMART_MEMORY VALUE = "Yes" Type = "Smart"/>
						<PART NUMBER = "840757-091"/>
						<TYPE VALUE = "DIMM DDR4"/>
						<SIZE VALUE = "16384 MB"/>
						<FREQUENCY VALUE = "2666 MHz"/>
						<MINIMUM_VOLTAGE VALUE = "1.20 v"/>
						<RANKS VALUE = "1"/>
						<TECHNOLOGY VALUE = "RDIMM"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "9"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "10"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "11"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_1>
						<SOCKET VALUE = "12"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_1>
					<CPU_2>
						<SOCKET VALUE = "1"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "2"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "3"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "4"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "5"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "6"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "7"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "8"/>
						<STATUS VALUE = "Good, In Use"/>
						<HP_SMART_MEMORY VALUE = "Yes" Type = "Smart"/>
						<PART NUMBER = "840757-091"/>
						<TYPE VALUE = "DIMM DDR4"/>
						<SIZE VALUE = "16384 MB"/>
						<FREQUENCY VALUE = "2666 MHz"/>
						<MINIMUM_VOLTAGE VALUE = "1.20 v"/>
						<RANKS VALUE = "1"/>
						<TECHNOLOGY VALUE = "RDIMM"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "9"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "10"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "11"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
					<CPU_2>
						<SOCKET VALUE = "12"/>
						<STATUS VALUE = "Not Present"/>
						<HP_SMART_MEMORY VALUE = "N/A" Type = "Unknown"/>
						<PART NUMBER = "N/A"/>
						<TYPE VALUE = "N/A"/>
						<SIZE VALUE = "N/A"/>
						<FREQUENCY VALUE = "N/A"/>
						<MINIMUM_VOLTAGE VALUE = "N/A"/>
						<RANKS VALUE = "N/A"/>
						<TECHNOLOGY VALUE = "N/A"/>
					</CPU_2>
				</MEMORY_DETAILS>
			</MEMORY>
			<NIC_INFORMATION>
				<iLO>
					<NETWORK_PORT VALUE = "iLO Dedicated Network Port"/>
					<PORT_DESCRIPTION VALUE = "iLO Dedicated Network Port"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "20:67:7c:e4:87:8a"/>
					<IP_ADDRESS VALUE = "N/A"/>
					<STATUS VALUE = "OK"/>
				</iLO>
				<NIC>
					<NETWORK_PORT VALUE = "Port 2"/>
					<PORT_DESCRIPTION VALUE = "iLO 4"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "d0:67:26:9c:81:d3"/>
					<IP_ADDRESS VALUE = "192.168.92.18"/>
					<STATUS VALUE = "OK"/>
				</NIC>
				<NIC>
					<NETWORK_PORT VALUE = "Port 1"/>
					<PORT_DESCRIPTION VALUE = "Network Controller"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "20:67:7c:d8:59:dc"/>
					<IP_ADDRESS VALUE = ""/>
					<STATUS VALUE = "OK"/>
				</NIC>
				<NIC>
					<NETWORK_PORT VALUE = "Port 2"/>
					<PORT_DESCRIPTION VALUE = "Network Controller"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "20:67:7c:d8:59:dd"/>
					<IP_ADDRESS VALUE = ""/>
					<STATUS VALUE = "Unknown"/>
				</NIC>
				<NIC>
					<NETWORK_PORT VALUE = "Port 3"/>
					<PORT_DESCRIPTION VALUE = "Network Controller"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "20:67:7c:d8:59:de"/>
					<IP_ADDRESS VALUE = ""/>
					<STATUS VALUE = "Unknown"/>
				</NIC>
				<NIC>
					<NETWORK_PORT VALUE = "Port 4"/>
					<PORT_DESCRIPTION VALUE = "Network Controller"/>
					<LOCATION VALUE = "Embedded"/>
					<MAC_ADDRESS VALUE = "20:67:7c:d8:59:df"/>
					<IP_ADDRESS VALUE = ""/>
					<STATUS VALUE = "Unknown"/>
				</NIC>
			</NIC_INFORMATION>
			<FIRMWARE_INFORMATION>
				<INDEX_1>
					<FIRMWARE_NAME VALUE = "iLO 5"/>
					<FIRMWARE_VERSION VALUE = "1.30 May 31 2018"/>
				</INDEX_1>
				<INDEX_2>
					<FIRMWARE_NAME VALUE = "System ROM"/>
					<FIRMWARE_VERSION VALUE = "U32 v1.40 (06/15/2018)"/>
				</INDEX_2>
				<INDEX_3>
					<FIRMWARE_NAME VALUE = "Redundant System ROM"/>
					<FIRMWARE_VERSION VALUE = "U32 v1.36 (02/14/2018)"/>
				</INDEX_3>
				<INDEX_4>
					<FIRMWARE_NAME VALUE = "Intelligent Platform Abstraction Data"/>
					<FIRMWARE_VERSION VALUE = "6.1.0 Build 12"/>
				</INDEX_4>
				<INDEX_5>
					<FIRMWARE_NAME VALUE = "Power Management Controller Firmware"/>
					<FIRMWARE_VERSION VALUE = "1.0.4"/>
					<FIRMWARE_FAMILY VALUE = "19h"/>
				</INDEX_5>
				<INDEX_6>
					<FIRMWARE_NAME VALUE = "Power Management Controller FW Bootloader"/>
					<FIRMWARE_VERSION VALUE = "1.1"/>
				</INDEX_6>
				<INDEX_7>
					<FIRMWARE_NAME VALUE = "System Programmable Logic Device"/>
					<FIRMWARE_VERSION VALUE = "0x2A"/>
				</INDEX_7>
				<INDEX_8>
					<FIRMWARE_NAME VALUE = "Server Platform Services (SPS) Firmware"/>
					<FIRMWARE_VERSION VALUE = "4.0.4.288"/>
				</INDEX_8>
				<INDEX_9>
					<FIRMWARE_NAME VALUE = "Intelligent Provisioning"/>
					<FIRMWARE_VERSION VALUE = "3.20.154"/>
				</INDEX_9>
				<INDEX_10>
					<FIRMWARE_NAME VALUE = "Innovation Engine (IE) Firmware"/>
					<FIRMWARE_VERSION VALUE = "0.1.5.2"/>
				</INDEX_10>
				<INDEX_11>
					<FIRMWARE_NAME VALUE = "Embedded Video Controller"/>
					<FIRMWARE_VERSION VALUE = "2.5"/>
				</INDEX_11>
				<INDEX_12>
					<FIRMWARE_NAME VALUE = "Network Controller"/>
					<FIRMWARE_VERSION VALUE = "20.12.41"/>
				</INDEX_12>
			</FIRMWARE_INFORMATION>
			<FANS>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 1"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "0" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 2"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "35" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 3"/>
					<STATUS VALUE = "Fail"/>
					<SPEED VALUE = "0" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 4"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "35" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 5"/>
					<STATUS VALUE = "Fail"/>
					<SPEED VALUE = "100" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 6"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "35" UNIT="Percentage"/>
				</FAN>
				<FAN>
					<ZONE VALUE = "System"/>
					<LABEL VALUE = "Fan 7"/>
					<STATUS VALUE = "OK"/>
					<SPEED VALUE = "100" UNIT="Percentage"/>
				</FAN>
			</FANS>
			<TEMPERATURE>
				<TEMP>
					<LABEL VALUE = "01-Inlet Ambient"/>
					<LOCATION VALUE = "Ambient"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "50" UNIT="Celsius"/>
					<CAUTION VALUE = "42" UNIT="Celsius"/>
					<CRITICAL VALUE = "47" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "02-CPU 1"/>
					<LOCATION VALUE = "CPU"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "40" UNIT="Celsius"/>
					<CAUTION VALUE = "70" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "03-CPU 2"/>
					<LOCATION VALUE = "CPU"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "75" UNIT="Celsius"/>
					<CAUTION VALUE = "70" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "04-P1 DIMM 1-6"/>
					<LOCATION VALUE = "Memory"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "06-P1 DIMM 7-12"/>
					<LOCATION VALUE = "Memory"/>
					<STATUS VALUE = "Fail"/>
					<CURRENTREADING VALUE = "99" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "08-P2 DIMM 1-6"/>
					<LOCATION VALUE = "Memory"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "10-P2 DIMM 7-12"/>
					<LOCATION VALUE = "Memory"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "17" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "12-HD Max"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "13-Exp Bay Drive"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "14-Stor Batt 1"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "15-Front Ambient"/>
					<LOCATION VALUE = "Ambient"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "17" UNIT="Celsius"/>
					<CAUTION VALUE = "60" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "16-VR P1"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "23" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "17-VR P2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "18-VR P1 Mem 1"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "20" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "19-VR P1 Mem 2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "20-VR P2 Mem 1"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "22" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "21-VR P2 Mem 2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "22" UNIT="Celsius"/>
					<CAUTION VALUE = "115" UNIT="Celsius"/>
					<CRITICAL VALUE = "120" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "22-Chipset"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "31" UNIT="Celsius"/>
					<CAUTION VALUE = "100" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "23-BMC"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "56" UNIT="Celsius"/>
					<CAUTION VALUE = "110" UNIT="Celsius"/>
					<CRITICAL VALUE = "115" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "24-BMC Zone"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "27" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "95" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "25-HD Controller"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "26-HD Cntlr Zone"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "27-LOM"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "40" UNIT="Celsius"/>
					<CAUTION VALUE = "100" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "28-LOM Card"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "29-I/O Zone"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "22" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "95" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "30-PCI 1"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "31-PCI 1 Zone"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "95" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "32-PCI 2"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "33-PCI 2 Zone"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "90" UNIT="Celsius"/>
					<CRITICAL VALUE = "95" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "34-PCI 3"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "35-PCI 3 Zone"/>
					<LOCATION VALUE = "I/O Board"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "37-Rear HD Max"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "38-Battery Zone"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "75" UNIT="Celsius"/>
					<CRITICAL VALUE = "80" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "39-P/S 1 Inlet"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "21" UNIT="Celsius"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "40-P/S 2 Inlet"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "41-P/S 1"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "40" UNIT="Celsius"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "42-P/S 2"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "43-E-Fuse"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "15" UNIT="Celsius"/>
					<CAUTION VALUE = "100" UNIT="Celsius"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "44-P/S 2 Zone"/>
					<LOCATION VALUE = "Power Supply"/>
					<STATUS VALUE = "OK"/>
					<CURRENTREADING VALUE = "19" UNIT="Celsius"/>
					<CAUTION VALUE = "75" UNIT="Celsius"/>
					<CRITICAL VALUE = "80" UNIT="Celsius"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "69-PCI 1 M2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "70-PCI 1 M2 Zn"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "71-PCI 2 M2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "72-PCI 2 M2 Zn"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "73-PCI 3 M2"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
				<TEMP>
					<LABEL VALUE = "74-PCI 3 M2 Zn"/>
					<LOCATION VALUE = "System"/>
					<STATUS VALUE = "Not Installed"/>
					<CURRENTREADING VALUE = "N/A"/>
					<CAUTION VALUE = "N/A"/>
					<CRITICAL VALUE = "N/A"/>
				</TEMP>
			</TEMPERATURE>
			<POWER_SUPPLIES>
				<POWER_SUPPLY_SUMMARY>
					<PRESENT_POWER_READING VALUE = "61 Watts"/>
					<POWER_MANAGEMENT_CONTROLLER_FIRMWARE_VERSION VALUE = "1.0.4"/>
					<POWER_SYSTEM_REDUNDANCY VALUE = "Not Redundant"/>
					<HP_POWER_DISCOVERY_SERVICES_REDUNDANCY_STATUS VALUE = "N/A"/>
					<HIGH_EFFICIENCY_MODE VALUE = "Balanced"/>
				</POWER_SUPPLY_SUMMARY>
				<SUPPLY>
					<LABEL VALUE = "Power Supply 1"/>
					<PRESENT VALUE = "Yes"/>
					<STATUS VALUE = "Good, In Use"/>
					<PDS VALUE = "No"/>
					<HOTPLUG_CAPABLE VALUE = "Yes"/>
					<MODEL VALUE = "865408-B21"/>
					<SPARE VALUE = "866729-001"/>
					<SERIAL_NUMBER VALUE = "5WBXL0C8JAJ1MA"/>
					<CAPACITY VALUE = "500 Watts"/>
					<FIRMWARE_VERSION VALUE = "1.02"/>
				</SUPPLY>
				<SUPPLY>
					<LABEL VALUE = "Power Supply 2"/>
					<PRESENT VALUE = "No"/>
					<STATUS VALUE = "Unknown"/>
					<PDS VALUE = "Other"/>
					<HOTPLUG_CAPABLE VALUE = "Yes"/>
					<MODEL VALUE = "N/A"/>
					<SPARE VALUE = "N/A"/>
					<SERIAL_NUMBER VALUE = "N/A"/>
					<CAPACITY VALUE = "N/A"/>
					<FIRMWARE_VERSION VALUE = "N/A"/>
				</SUPPLY>
				<SUPPLY>
					<LABEL VALUE = "Power Supply 3"/>
					<PRESENT VALUE = "Yes"/>
					<STATUS VALUE = "Good, In Use"/>
					<PDS VALUE = "No"/>
					<HOTPLUG_CAPABLE VALUE = "No"/>
					<MODEL VALUE = "						  "/>
					<SPARE VALUE = "Unknown"/>
					<SERIAL_NUMBER VALUE = "			 "/>
					<CAPACITY VALUE = "0 Watts"/>
					<FIRMWARE_VERSION VALUE = "N/A"/>
				</SUPPLY>
				<POWER_DISCOVERY_SERVICES_IPDU_SUMMARY>
					<IPDU>
						<BAY VALUE="2"/>
						<STATUS VALUE="iPDU Not Redundant"/>
						<PART_NUMBER VALUE="AF522A"/>
						<SERIAL_NUMBER VALUE="2CJ0221672"/>
						<MAC_ADDRESS VALUE="d8:d3:85:6d:36:9c"/>
						<IPDU_LINK VALUE="http://16.85.177.189"/>
					</IPDU>
				</POWER_DISCOVERY_SERVICES_IPDU_SUMMARY>
				<SMART_STORAGE_BATTERY>
					<LABEL VALUE="Battery 1"/>
					<PRESENT VALUE="Yes"/>
					<STATUS VALUE="OK"/>
					<MODEL VALUE="727258-B21"/>
					<SPARE VALUE="750450-001"/>
					<SERIAL_NUMBER VALUE="6EMYC0AWY7X77Q"/>
					<CAPACITY VALUE="96 Watts"/>
					<FIRMWARE_VERSION VALUE="1.1"/>
				</SMART_STORAGE_BATTERY>
			</POWER_SUPPLIES>
			<HEALTH_AT_A_GLANCE>
				<BIOS_HARDWARE STATUS= "OK"/>
				<FANS STATUS= "OK"/>
				<FANS REDUNDANCY= "Redundant"/>
				<TEMPERATURE STATUS= "OK"/>
				<POWER_SUPPLIES STATUS= "OK"/>
				<POWER_SUPPLIES REDUNDANCY= "Not Redundant"/>
				<BATTERY STATUS= "Not Installed"/>
				<PROCESSOR STATUS= "OK"/>
				<MEMORY STATUS= "OK"/>
				<NETWORK STATUS= "OK"/>
				<STORAGE STATUS= "OK"/>
			</HEALTH_AT_A_GLANCE>
			<STORAGE>
				<CONTROLLER>
					<LABEL VALUE="Controller on System Board"/>
					<STATUS VALUE="OK"/>
					<CONTROLLER_STATUS VALUE="OK"/>
					<SERIAL_NUMBER VALUE="PDNLH0BRH8A25C"/>
					<MODEL VALUE="Smart Array P440ar Controller"/>
					<FW_VERSION VALUE="3.52"/>
					<CACHE_MODULE_STATUS VALUE="OK"/>
					<CACHE_MODULE_SERIAL_NUM VALUE="PDNLH0BRH8A25C"/>
					<CACHE_MODULE_MEMORY VALUE="2097152 KB"/>
					<ENCRYPTION_STATUS VALUE="Not Enabled"/>
					<ENCRYPTION_SELF_TEST_STATUS VALUE="OK"/>
					<ENCRYPTION_CSP_STATUS VALUE="OK"/>
					<DRIVE_ENCLOSURE>
						<LABEL VALUE="Port 1I Box 3"/>
						<STATUS VALUE="OK"/>
						<DRIVE_BAY VALUE="04"/>
					</DRIVE_ENCLOSURE>
					<DRIVE_ENCLOSURE>
						<LABEL VALUE="Port 2I Box 0"/>
						<STATUS VALUE="OK"/>
						<DRIVE_BAY VALUE="04"/>
					</DRIVE_ENCLOSURE>
					<LOGICAL_DRIVE>
						<LABEL VALUE="01"/>
						<STATUS VALUE="OK"/>
						<CAPACITY VALUE="231 GiB"/>
						<FAULT_TOLERANCE VALUE="RAID 0"/>
						<LOGICAL_DRIVE_TYPE VALUE="Data LUN"/>
						<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 2"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EGT20000C5236EYR"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 2"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 1"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EJE30000C523FA8T"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 1"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
					</LOGICAL_DRIVE>
					<LOGICAL_DRIVE>
						<LABEL VALUE="02"/>
						<STATUS VALUE="OK"/>
						<CAPACITY VALUE="231 GiB"/>
						<FAULT_TOLERANCE VALUE="RAID 0"/>
						<LOGICAL_DRIVE_TYPE VALUE="Data LUN"/>
						<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 2"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EGT20000C5236EYR"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 2"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 1"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EJE30000C523FA8T"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 1"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
					</LOGICAL_DRIVE>
					<LOGICAL_DRIVE>
						<LABEL VALUE="03"/>
						<STATUS VALUE="OK"/>
						<CAPACITY VALUE="231 GiB"/>
						<FAULT_TOLERANCE VALUE="RAID 0"/>
						<LOGICAL_DRIVE_TYPE VALUE="Data LUN"/>
						<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 2"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EGT20000C5236EYR"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 2"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 1"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EJE30000C523FA8T"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 1"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
					</LOGICAL_DRIVE>
					<LOGICAL_DRIVE>
						<LABEL VALUE="04"/>
						<STATUS VALUE="OK"/>
						<CAPACITY VALUE="231 GiB"/>
						<FAULT_TOLERANCE VALUE="RAID 0"/>
						<LOGICAL_DRIVE_TYPE VALUE="Data LUN"/>
						<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 2"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EGT20000C5236EYR"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 2"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
						<PHYSICAL_DRIVE>
							<LABEL VALUE="Port 1I Box 3 Bay 1"/>
							<STATUS VALUE="OK"/>
							<SERIAL_NUMBER VALUE="9XF3EJE30000C523FA8T"/>
							<MODEL VALUE="MM0500FBFVQ"/>
							<CAPACITY VALUE="465 GiB"/>
							<MARKETING_CAPACITY VALUE="341 GB"/>
							<LOCATION VALUE="Port 1I Box 3 Bay 1"/>
							<FW_VERSION VALUE="HPD8"/>
							<DRIVE_CONFIGURATION VALUE="Configured"/>
							<ENCRYPTION_STATUS VALUE="Not Encrypted"/>
							<MEDIA_TYPE VALUE="HDD"/>
						</PHYSICAL_DRIVE>
					</LOGICAL_DRIVE>
				</CONTROLLER>
				<DISCOVERY_STATUS>
					<STATUS VALUE="Discovery Complete"/>
				</DISCOVERY_STATUS>
			</STORAGE>
		</GET_EMBEDDED_HEALTH_DATA>
	`)

	// parse data
	health1, err := GetHealthData(data)
	if err != nil {
		t.Errorf("Failed to parse health1 data: %s", err.Error())
	}

	// marshal health1 data
	xml, err := xml.MarshalIndent(health1, "", "\t")
	if err != nil {
		t.Errorf("Failed to marshal health1 data: %s", err.Error())
	}

	// fix self-closing tags
	selfClosingXML := regexp.MustCompile("></([a-zA-Z0-9-_]+)>").ReplaceAll(xml, []byte("/>"))

	// parse data
	health2, err := GetHealthData(selfClosingXML)
	if err != nil {
		t.Errorf("Failed to parse health2 data: %s", err.Error())
	}

	if !reflect.DeepEqual(health1, health2) {
		t.Errorf("Unmarshaled Health data does not equal Re-Unmarshaled, must be equal")
	}

}
