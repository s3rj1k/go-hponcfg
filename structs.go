package hponcfg

import (
	"encoding/xml"
)

// MemoryDetailsSummaryOnCPU - memory summary details per CPU
type MemoryDetailsSummaryOnCPU struct {
	NumberOfSockets struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"NUMBER_OF_SOCKETS"`
	TotalMemorySize struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"TOTAL_MEMORY_SIZE"`
	OperatingFrequency struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"OPERATING_FREQUENCY"`
	OperatingVoltage struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"OPERATING_VOLTAGE"`
}

// MemoryDetailsOnCPU - memory details per CPU
type MemoryDetailsOnCPU struct {
	Socket struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SOCKET"`
	Status struct {
		Value     string `xml:"VALUE,attr"`
		Installed bool   `xml:"-"`
	} `xml:"STATUS"`
	HPSmartMemory struct {
		Value string `xml:"VALUE,attr"`
		Type  string `xml:"Type,attr"`
	} `xml:"HP_SMART_MEMORY"`
	Part struct {
		Number string `xml:"NUMBER,attr"`
	} `xml:"PART"`
	Type struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"TYPE"`
	Size struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SIZE"`
	Frequency struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FREQUENCY"`
	MinimumVoltage struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MINIMUM_VOLTAGE"`
	Ranks struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"RANKS"`
	Technology struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"TECHNOLOGY"`
}

// NIC - network interface information
type NIC struct {
	NetworkPort struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"NETWORK_PORT"`
	PortDescription struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"PORT_DESCRIPTION"`
	Location struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LOCATION"`
	MACAddress struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MAC_ADDRESS"`
	IPAddress struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"IP_ADDRESS"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	ILO struct {
		Value bool `xml:"-"`
	} `xml:"-"`
}

// Processor - CPU information
type Processor struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Name struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"NAME"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	Speed struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SPEED"`
	ExecutionTechnology struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"EXECUTION_TECHNOLOGY"`
	MemoryTechnology struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MEMORY_TECHNOLOGY"`
	InternalL1Cache struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"INTERNAL_L1_CACHE"`
	InternalL2Cache struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"INTERNAL_L2_CACHE"`
	InternalL3Cache struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"INTERNAL_L3_CACHE"`
}

// AdvancedMemoryProtection - memory protection details
type AdvancedMemoryProtection struct {
	AMPModeStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"AMP_MODE_STATUS"`
	ConfiguredAMPMode struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CONFIGURED_AMP_MODE"`
	AvailableAMPModes struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"AVAILABLE_AMP_MODES"`
}

// Firmware - firmware information
type Firmware struct {
	FirmwareName struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FIRMWARE_NAME"`
	FirmwareVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FIRMWARE_VERSION"`
	FirmwareFamily struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"FIRMWARE_FAMILY"`
}

// HealthAtAGlance - short health status overview
type HealthAtAGlance struct {
	BIOSHardware struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"BIOS_HARDWARE"`
	Fans []struct {
		Status     string `xml:"STATUS,attr,omitempty"`
		Redundancy string `xml:"REDUNDANCY,attr,omitempty"`
	} `xml:"FANS"`
	Fan struct {
		Status     string `xml:"-"`
		Redundancy string `xml:"-"`
	} `xml:"-"`
	Temperature struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"TEMPERATURE"`
	PowerSupplies []struct {
		Status     string `xml:"STATUS,attr,omitempty"`
		Redundancy string `xml:"REDUNDANCY,attr,omitempty"`
	} `xml:"POWER_SUPPLIES"`
	PowerSupply struct {
		Status     string `xml:"-"`
		Redundancy string `xml:"-"`
	} `xml:"-"`
	Battery struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"BATTERY"`
	Processor struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"PROCESSOR"`
	Memory struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"MEMORY"`
	Network struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"NETWORK"`
	Storage struct {
		Status string `xml:"STATUS,attr"`
	} `xml:"STORAGE"`
}

// PowerSupplySummary - power supply summary
type PowerSupplySummary struct {
	PresentPowerReading struct {
		Value     string `xml:"VALUE,attr"`
		Available bool   `xml:"-"`
	} `xml:"PRESENT_POWER_READING"`
	PowerManagementControllerFirmwareVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"POWER_MANAGEMENT_CONTROLLER_FIRMWARE_VERSION"`
	PowerSystemRedundancy struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"POWER_SYSTEM_REDUNDANCY"`
	HPPowerDiscoveryServicesRedundancyStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"HP_POWER_DISCOVERY_SERVICES_REDUNDANCY_STATUS"`
	HighEfficiencyMode struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"HIGH_EFFICIENCY_MODE"`
}

// Supply - power supply info
type Supply struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Present struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"PRESENT"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	PDS struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"PDS"`
	HotplugCapable struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"HOTPLUG_CAPABLE"`
	Model struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MODEL"`
	Spare struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SPARE"`
	SerialNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SERIAL_NUMBER"`
	Capacity struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CAPACITY"`
	FirmwareVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FIRMWARE_VERSION"`
}

// IPDU - intelligent power distribution unit info
type IPDU struct {
	Bay struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"BAY"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	PartNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"PART_NUMBER"`
	SerialNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SERIAL_NUMBER"`
	MACAddress struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MAC_ADDRESS"`
	IPDULink struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"IPDU_LINK"`
}

// SmartStorageBattery - smart storage battery info
type SmartStorageBattery struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Present struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"PRESENT"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	Model struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MODEL"`
	Spare struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SPARE"`
	SerialNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SERIAL_NUMBER"`
	Capacity struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CAPACITY"`
	FirmwareVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FIRMWARE_VERSION"`
}

// Temp - temperature sensor info
type Temp struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Location struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LOCATION"`
	Status struct {
		Value     string `xml:"VALUE,attr"`
		Ok        bool   `xml:"-"`
		Installed bool   `xml:"-"`
	} `xml:"STATUS"`
	CurrentReading struct {
		Value    string `xml:"VALUE,attr"`
		IntValue int64  `xml:"-"`
		Unit     string `xml:"UNIT,attr"`
	} `xml:"CURRENTREADING"`
	Caution struct {
		Value    string `xml:"VALUE,attr"`
		IntValue int64  `xml:"-"`
		Unit     string `xml:"UNIT,attr"`
	} `xml:"CAUTION"`
	Critical struct {
		Value    string `xml:"VALUE,attr"`
		IntValue int64  `xml:"-"`
		Unit     string `xml:"UNIT,attr"`
	} `xml:"CRITICAL"`
}

// Fan - fan sensor information
type Fan struct {
	Zone struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"ZONE"`
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Status struct {
		Value string `xml:"VALUE,attr"`
		Ok    bool   `xml:"-"`
	} `xml:"STATUS"`
	Speed struct {
		Value    string `xml:"VALUE,attr"`
		IntValue int64  `xml:"-"`
		Unit     string `xml:"UNIT,attr"`
	} `xml:"SPEED"`
}

// DiscoveryStatus - storage discovery status
type DiscoveryStatus struct {
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
}

// DriveEnclosure - storage drive enclosure status
type DriveEnclosure struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	DriveBay struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"DRIVE_BAY"`
}

// PhysicalDrive - storage physical drive info
type PhysicalDrive struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	SerialNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SERIAL_NUMBER"`
	Model struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MODEL"`
	Capacity struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CAPACITY"`
	MarketingCapacity struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MARKETING_CAPACITY"`
	Location struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LOCATION"`
	FWVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FW_VERSION"`
	DriveConfiguration struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"DRIVE_CONFIGURATION"`
	EncryptionStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"ENCRYPTION_STATUS"`
	MediaType struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MEDIA_TYPE"`
}

// LogicalDrive - storage logical drive status
type LogicalDrive struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	Capacity struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CAPACITY"`
	FaultTolerance struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FAULT_TOLERANCE"`
	LogicalDriveType struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LOGICAL_DRIVE_TYPE"`
	EncryptionStatus struct {
		Value string `xml:"VALUE,attr"`
	}
	PhysicalDrive []*PhysicalDrive `xml:"PHYSICAL_DRIVE,omitempty"`
}

// Controller - storage controller info
type Controller struct {
	Label struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"LABEL"`
	Status struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"STATUS"`
	ControllerStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CONTROLLER_STATUS"`
	SerialNumber struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"SERIAL_NUMBER"`
	Model struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"MODEL"`
	FWVersion struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"FW_VERSION"`
	CacheModuleStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CACHE_MODULE_STATUS"`
	CacheModuleSerialNum struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CACHE_MODULE_SERIAL_NUM"`
	CacheModuleMemory struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"CACHE_MODULE_MEMORY"`
	EncryptionStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"ENCRYPTION_STATUS"`
	EncryptionSelfTestStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"ENCRYPTION_SELF_TEST_STATUS"`
	EncryptionCSPStatus struct {
		Value string `xml:"VALUE,attr"`
	} `xml:"ENCRYPTION_CSP_STATUS"`
	DriveEnclosure []*DriveEnclosure `xml:"DRIVE_ENCLOSURE"`
	LogicalDrive   []*LogicalDrive   `xml:"LOGICAL_DRIVE,omitempty"`
}

// GetFWVersion - we use this to get missing info for BMC type and License type
type GetFWVersion struct {
	XMLName             xml.Name `xml:"GET_FW_VERSION"`
	ManagementProcessor string   `xml:"MANAGEMENT_PROCESSOR,attr"`
	LicenseType         string   `xml:"LICENSE_TYPE,attr"`
}

// GetEmbeddedHealthData - HP health data details
type GetEmbeddedHealthData struct {
	XMLName    xml.Name `xml:"GET_EMBEDDED_HEALTH_DATA"`
	Processors struct {
		Processor []*Processor `xml:"PROCESSOR,omitempty"`
	} `xml:"PROCESSORS,omitempty"`
	Memory struct {
		AdvancedMemoryProtection *AdvancedMemoryProtection `xml:"ADVANCED_MEMORY_PROTECTION,omitempty"`
		MemoryDetailsSummary     struct {
			CPU1   *MemoryDetailsSummaryOnCPU `xml:"CPU_1,omitempty"`
			IsCPU1 bool                       `xml:"IsCPU_1,omitempty"`
			CPU2   *MemoryDetailsSummaryOnCPU `xml:"CPU_2,omitempty"`
			IsCPU2 bool                       `xml:"IsCPU_2,omitempty"`
			CPU3   *MemoryDetailsSummaryOnCPU `xml:"CPU_3,omitempty"`
			IsCPU3 bool                       `xml:"IsCPU_3,omitempty"`
			CPU4   *MemoryDetailsSummaryOnCPU `xml:"CPU_4,omitempty"`
			IsCPU4 bool                       `xml:"IsCPU_4,omitempty"`
		} `xml:"MEMORY_DETAILS_SUMMARY,omitempty"`
		MemoryDetails struct {
			CPU1   []*MemoryDetailsOnCPU `xml:"CPU_1,omitempty"`
			IsCPU1 bool                  `xml:"IsCPU_1,omitempty"`
			CPU2   []*MemoryDetailsOnCPU `xml:"CPU_2,omitempty"`
			IsCPU2 bool                  `xml:"IsCPU_2,omitempty"`
			CPU3   []*MemoryDetailsOnCPU `xml:"CPU_3,omitempty"`
			IsCPU3 bool                  `xml:"IsCPU_3,omitempty"`
			CPU4   []*MemoryDetailsOnCPU `xml:"CPU_4,omitempty"`
			IsCPU4 bool                  `xml:"IsCPU_4,omitempty"`
		} `xml:"MEMORY_DETAILS,omitempty"`
	} `xml:"MEMORY,omitempty"`
	NICInformation struct {
		ILO    *NIC   `xml:"iLO,omitempty"`
		NICs   []*NIC `xml:"NIC,omitempty"`
		Merged []*NIC `xml:"-"`
	} `xml:"NIC_INFORMATION,omitempty"`
	FirmwareInformation struct {
		Index []*Firmware `xml:"INDEX,omitempty"`
	} `xml:"FIRMWARE_INFORMATION,omitempty"`
	HealthAtAGlance *HealthAtAGlance `xml:"HEALTH_AT_A_GLANCE,omitempty"`
	PowerSupplies   struct {
		PowerSupplySummary                *PowerSupplySummary `xml:"POWER_SUPPLY_SUMMARY,omitempty"`
		Supply                            []*Supply           `xml:"SUPPLY,omitempty"`
		PowerDiscoveryServicesIPDUSummary struct {
			IPDU *IPDU `xml:"IPDU,omitempty"`
		} `xml:"POWER_DISCOVERY_SERVICES_IPDU_SUMMARY,omitempty"`
		SmartStorageBattery *SmartStorageBattery `xml:"SMART_STORAGE_BATTERY,omitempty"`
	} `xml:"POWER_SUPPLIES,omitempty"`
	Temperature struct {
		Temp []*Temp `xml:"TEMP,omitempty"`
	} `xml:"TEMPERATURE,omitempty"`
	Fans struct {
		Fan []*Fan `xml:"FAN,omitempty"`
	} `xml:"FANS,omitempty"`
	Storage struct {
		Controller      *Controller      `xml:"CONTROLLER,omitempty"`
		DiscoveryStatus *DiscoveryStatus `xml:"DISCOVERY_STATUS,omitempty"`
	} `xml:"STORAGE,omitempty"`
}

// GetNetworkSettings - network settings details
type GetNetworkSettings struct {
	XMLName   xml.Name `xml:"GET_NETWORK_SETTINGS,omitempty"`
	EnableNIC struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ENABLE_NIC,omitempty"`
	SharedNetworkPort struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SHARED_NETWORK_PORT,omitempty"`
	VlanEnabled struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"VLAN_ENABLED,omitempty"`
	VlanID struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"VLAN_ID,omitempty"`
	SpeedAutoSelect struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SPEED_AUTOSELECT,omitempty"`
	NICSpeed struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"NIC_SPEED,omitempty"`
	FullDuplex struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"FULL_DUPLEX,omitempty"`
	DHCPEnable struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_ENABLE,omitempty"`
	DHCPGateway struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_GATEWAY,omitempty"`
	DHCPDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_DNS_SERVER,omitempty"`
	DHCPWINSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_WINS_SERVER,omitempty"`
	DHCPStaticRoute struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_STATIC_ROUTE,omitempty"`
	DHCPDomainName struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_DOMAIN_NAME,omitempty"`
	DHCPSNTPSettings struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCP_SNTP_SETTINGS,omitempty"`
	RegWINSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"REG_WINS_SERVER,omitempty"`
	RegDDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"REG_DDNS_SERVER,omitempty"`
	PingGateway struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"PING_GATEWAY,omitempty"`
	MACAddress struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"MAC_ADDRESS,omitempty"`
	IPAddress struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IP_ADDRESS,omitempty"`
	SubnetMask struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SUBNET_MASK,omitempty"`
	GatewayIPAddress struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"GATEWAY_IP_ADDRESS,omitempty"`
	DNSName struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DNS_NAME,omitempty"`
	DomainName struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DOMAIN_NAME,omitempty"`
	PrimDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"PRIM_DNS_SERVER,omitempty"`
	SecDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SEC_DNS_SERVER,omitempty"`
	TerDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"TER_DNS_SERVER,omitempty"`
	PrimWINSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"PRIM_WINS_SERVER,omitempty"`
	SecWINSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SEC_WINS_SERVER,omitempty"`
	SNTPServer1 struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SNTP_SERVER1,omitempty"`
	SNTPServer2 struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SNTP_SERVER2,omitempty"`
	Timezone struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"TIMEZONE,omitempty"`
	StaticRoute1 struct {
		Dest    string `xml:"DEST,attr,omitempty"`
		Mask    string `xml:"MASK,attr,omitempty"`
		Gateway string `xml:"GATEWAY,attr,omitempty"`
	} `xml:"STATIC_ROUTE_1,omitempty"`
	StaticRoute2 struct {
		Dest    string `xml:"DEST,attr,omitempty"`
		Mask    string `xml:"MASK,attr,omitempty"`
		Gateway string `xml:"GATEWAY,attr,omitempty"`
	} `xml:"STATIC_ROUTE_2,omitempty"`
	StaticRoute3 struct {
		Dest    string `xml:"DEST,attr,omitempty"`
		Mask    string `xml:"MASK,attr,omitempty"`
		Gateway string `xml:"GATEWAY,attr,omitempty"`
	} `xml:"STATIC_ROUTE_3,omitempty"`
	IPv6Address []struct {
		Value      string `xml:"VALUE,attr,omitempty"`
		PrefixLen  string `xml:"PREFIXLEN,attr,omitempty"`
		AddrSource string `xml:"ADDR_SOURCE,attr,omitempty"`
		AddrStatus string `xml:"ADDR_STATUS,attr,omitempty"`
	} `xml:"IPV6_ADDRESS,omitempty"`
	IPv6StaticRoute1 struct {
		IPv6Dest    string `xml:"IPV6_DEST,attr,omitempty"`
		PrefixLen   string `xml:"PREFIXLEN,attr,omitempty"`
		IPv6Gateway string `xml:"IPV6_GATEWAY,attr,omitempty"`
		AddrStatus  string `xml:"ADDR_STATUS,attr,omitempty"`
	} `xml:"IPV6_STATIC_ROUTE_1,omitempty"`
	IPv6StaticRoute2 struct {
		IPv6Dest    string `xml:"IPV6_DEST,attr,omitempty"`
		PrefixLen   string `xml:"PREFIXLEN,attr,omitempty"`
		IPv6Gateway string `xml:"IPV6_GATEWAY,attr,omitempty"`
		AddrStatus  string `xml:"ADDR_STATUS,attr,omitempty"`
	} `xml:"IPV6_STATIC_ROUTE_2,omitempty"`
	IPv6StaticRoute3 struct {
		IPv6Dest    string `xml:"IPV6_DEST,attr,omitempty"`
		PrefixLen   string `xml:"PREFIXLEN,attr,omitempty"`
		IPv6Gateway string `xml:"IPV6_GATEWAY,attr,omitempty"`
		AddrStatus  string `xml:"ADDR_STATUS,attr,omitempty"`
	} `xml:"IPV6_STATIC_ROUTE_3,omitempty"`
	IPv6PrimDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_PRIM_DNS_SERVER,omitempty"`
	IPv6SecDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_SEC_DNS_SERVER,omitempty"`
	IPv6TerDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_TER_DNS_SERVER,omitempty"`
	IPv6DefaultGateway struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_DEFAULT_GATEWAY,omitempty"`
	IPv6PreferredProtocol struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_PREFERRED_PROTOCOL,omitempty"`
	IPv6AddrAutocfg struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_ADDR_AUTOCFG,omitempty"`
	IPv6RegDDNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"IPV6_REG_DDNS_SERVER,omitempty"`
	DHCPv6StatlessEnable struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_STATELESS_ENABLE,omitempty"`
	DHCPv6StatefulEnable struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_STATEFUL_ENABLE,omitempty"`
	DHCPv6RapidCommit struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_RAPID_COMMIT,omitempty"`
	DHCPv6DomainName struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_DOMAIN_NAME,omitempty"`
	DHCPv6SNTPSettings struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_SNTP_SETTINGS,omitempty"`
	DHCPv6DNSServer struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"DHCPV6_DNS_SERVER,omitempty"`
	ILONICAutoSelect struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ILO_NIC_AUTO_SELECT,omitempty"`
	ILONICAutoSNPScan struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ILO_NIC_AUTO_SNP_SCAN,omitempty"`
	ILONICAutoDelay struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ILO_NIC_AUTO_DELAY,omitempty"`
	ILONICFailOver struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ILO_NIC_FAIL_OVER,omitempty"`
	ILONICFailOverDelay struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"ILO_NIC_FAIL_OVER_DELAY,omitempty"`
	SNPPort struct {
		Value string `xml:"VALUE,attr,omitempty"`
	} `xml:"SNP_PORT,omitempty"`
}
