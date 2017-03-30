// Code generated by protoc-gen-go.
// source: ospfv3_edm_interface.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_areas_area_interface_table_interface is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_interface.proto

It has these top-level messages:
	Ospfv3EdmInterface_KEYS
	Ospfv3EdmInterface
	Ospfv3EdmInterfaceNbr
	Ospfv3EdmInterfaceUp
	Ospfv3EdmInterfaceBfd
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_areas_area_interface_table_interface

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OSPFv3 interface information
type Ospfv3EdmInterface_KEYS struct {
	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName       string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AreaId        uint32 `protobuf:"varint,3,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ospfv3EdmInterface_KEYS) Reset()                    { *m = Ospfv3EdmInterface_KEYS{} }
func (m *Ospfv3EdmInterface_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterface_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmInterface_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmInterface_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmInterface_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmInterface_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Ospfv3EdmInterface_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmInterface struct {
	// Interface state
	InterfaceState string `protobuf:"bytes,50,opt,name=interface_state,json=interfaceState" json:"interface_state,omitempty"`
	// If true, line protocol is up
	IsInterfaceLineUp bool `protobuf:"varint,51,opt,name=is_interface_line_up,json=isInterfaceLineUp" json:"is_interface_line_up,omitempty"`
	// If true, interface IP security is required
	IsInterfaceIpSecurityRequired bool `protobuf:"varint,52,opt,name=is_interface_ip_security_required,json=isInterfaceIpSecurityRequired" json:"is_interface_ip_security_required,omitempty"`
	// If true, interface IP security is active
	IsInterfaceIpSecurityActive bool `protobuf:"varint,53,opt,name=is_interface_ip_security_active,json=isInterfaceIpSecurityActive" json:"is_interface_ip_security_active,omitempty"`
	// Interface IPv6 address
	InterfaceAddress string `protobuf:"bytes,54,opt,name=interface_address,json=interfaceAddress" json:"interface_address,omitempty"`
	// Interface number
	InterfaceNumber uint32 `protobuf:"varint,55,opt,name=interface_number,json=interfaceNumber" json:"interface_number,omitempty"`
	// Interface router ID
	InterfaceRouterId string `protobuf:"bytes,56,opt,name=interface_router_id,json=interfaceRouterId" json:"interface_router_id,omitempty"`
	// Network type
	NetworkType string `protobuf:"bytes,57,opt,name=network_type,json=networkType" json:"network_type,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,58,opt,name=interface_link_cost,json=interfaceLinkCost" json:"interface_link_cost,omitempty"`
	// If true, interface flood reduction is active
	IsInterfaceFloodReduction bool `protobuf:"varint,59,opt,name=is_interface_flood_reduction,json=isInterfaceFloodReduction" json:"is_interface_flood_reduction,omitempty"`
	// If true, configured as demand circuit
	IsDemandCircuitConfigured bool `protobuf:"varint,60,opt,name=is_demand_circuit_configured,json=isDemandCircuitConfigured" json:"is_demand_circuit_configured,omitempty"`
	// If true, interface running as demand circuit
	IsInterfaceDemandCircuit bool `protobuf:"varint,61,opt,name=is_interface_demand_circuit,json=isInterfaceDemandCircuit" json:"is_interface_demand_circuit,omitempty"`
	// Number of LSAs with demand circuit bit not set for the area in which the interface is running
	InterfaceDcBitlessLsAs uint32 `protobuf:"varint,62,opt,name=interface_dc_bitless_ls_as,json=interfaceDcBitlessLsAs" json:"interface_dc_bitless_ls_as,omitempty"`
	// Interface transmission delay (sec)
	TransmissionDelay uint32 `protobuf:"varint,63,opt,name=transmission_delay,json=transmissionDelay" json:"transmission_delay,omitempty"`
	// Interface state
	OspfInterfaceState string `protobuf:"bytes,64,opt,name=ospf_interface_state,json=ospfInterfaceState" json:"ospf_interface_state,omitempty"`
	// Interface priority
	InterfacePriority uint32 `protobuf:"varint,65,opt,name=interface_priority,json=interfacePriority" json:"interface_priority,omitempty"`
	// If true, designated router
	IsDesignatedRouter bool `protobuf:"varint,66,opt,name=is_designated_router,json=isDesignatedRouter" json:"is_designated_router,omitempty"`
	// Designated router ID
	DesignatedRouterId string `protobuf:"bytes,67,opt,name=designated_router_id,json=designatedRouterId" json:"designated_router_id,omitempty"`
	// Designated router interface address
	DesignatedRouterAddress string `protobuf:"bytes,68,opt,name=designated_router_address,json=designatedRouterAddress" json:"designated_router_address,omitempty"`
	// Backup designated router ID
	BackupDesignatedRouterId string `protobuf:"bytes,69,opt,name=backup_designated_router_id,json=backupDesignatedRouterId" json:"backup_designated_router_id,omitempty"`
	// Backup designated router interface address
	BackupDesignatedRouterAddress string `protobuf:"bytes,70,opt,name=backup_designated_router_address,json=backupDesignatedRouterAddress" json:"backup_designated_router_address,omitempty"`
	// The amount of time in seconds before flush timer for old network LSA expires
	NetworkLsaFlushTimer uint32 `protobuf:"varint,71,opt,name=network_lsa_flush_timer,json=networkLsaFlushTimer" json:"network_lsa_flush_timer,omitempty"`
	// Filter is configured for out going LSAs
	IsInterfaceLsaFiltered bool `protobuf:"varint,72,opt,name=is_interface_lsa_filtered,json=isInterfaceLsaFiltered" json:"is_interface_lsa_filtered,omitempty"`
	// Configured hello interval (s)
	HelloInterval uint32 `protobuf:"varint,73,opt,name=hello_interval,json=helloInterval" json:"hello_interval,omitempty"`
	// Configured dead interval (s)
	DeadInterval uint32 `protobuf:"varint,74,opt,name=dead_interval,json=deadInterval" json:"dead_interval,omitempty"`
	// Configured wait interval (s)
	WaitInterval uint32 `protobuf:"varint,75,opt,name=wait_interval,json=waitInterval" json:"wait_interval,omitempty"`
	// Configured retransmit interval (s)
	InterfaceRetransmissionInterval uint32 `protobuf:"varint,76,opt,name=interface_retransmission_interval,json=interfaceRetransmissionInterval" json:"interface_retransmission_interval,omitempty"`
	// Time until next Hello (s)
	NextHelloTime uint32 `protobuf:"varint,77,opt,name=next_hello_time,json=nextHelloTime" json:"next_hello_time,omitempty"`
	// Interface authentication spi
	InterfaceAuthenticationSpi uint32 `protobuf:"varint,78,opt,name=interface_authentication_spi,json=interfaceAuthenticationSpi" json:"interface_authentication_spi,omitempty"`
	// Interface authentication transmit
	InterfaceAuthenticationTransmit uint32 `protobuf:"varint,79,opt,name=interface_authentication_transmit,json=interfaceAuthenticationTransmit" json:"interface_authentication_transmit,omitempty"`
	// If true, interface encryption is enabled
	IsInterfaceEncryptionEnabled bool `protobuf:"varint,80,opt,name=is_interface_encryption_enabled,json=isInterfaceEncryptionEnabled" json:"is_interface_encryption_enabled,omitempty"`
	// If true prefix suppression is enabled
	IsPrefixSuppress bool `protobuf:"varint,81,opt,name=is_prefix_suppress,json=isPrefixSuppress" json:"is_prefix_suppress,omitempty"`
	// Interface encryption spi
	InterfaceEncryptionSpi uint32 `protobuf:"varint,82,opt,name=interface_encryption_spi,json=interfaceEncryptionSpi" json:"interface_encryption_spi,omitempty"`
	// Interface encryption transmitted
	InterfaceEncryptionTransmitted uint32 `protobuf:"varint,83,opt,name=interface_encryption_transmitted,json=interfaceEncryptionTransmitted" json:"interface_encryption_transmitted,omitempty"`
	// Interface encrypted authentication transmitted
	InterfaceEncryptedAuthenticationTransmitted uint32 `protobuf:"varint,84,opt,name=interface_encrypted_authentication_transmitted,json=interfaceEncryptedAuthenticationTransmitted" json:"interface_encrypted_authentication_transmitted,omitempty"`
	// Information for neighbors on the interface
	InterfaceNeighborList []*Ospfv3EdmInterfaceNbr `protobuf:"bytes,85,rep,name=interface_neighbor_list,json=interfaceNeighborList" json:"interface_neighbor_list,omitempty"`
	// Number of adjacent neighbors
	AdjacentNeighbor uint32 `protobuf:"varint,86,opt,name=adjacent_neighbor,json=adjacentNeighbor" json:"adjacent_neighbor,omitempty"`
	// Active interface details
	ActiveInterface *Ospfv3EdmInterfaceUp `protobuf:"bytes,87,opt,name=active_interface,json=activeInterface" json:"active_interface,omitempty"`
	// BFD information
	InterfaceBfd *Ospfv3EdmInterfaceBfd `protobuf:"bytes,88,opt,name=interface_bfd,json=interfaceBfd" json:"interface_bfd,omitempty"`
	// Interface reference count
	InterfaceReferences uint32 `protobuf:"varint,89,opt,name=interface_references,json=interfaceReferences" json:"interface_references,omitempty"`
	// If true, configured as LDP sync
	ConfiguredLdpSync bool `protobuf:"varint,90,opt,name=configured_ldp_sync,json=configuredLdpSync" json:"configured_ldp_sync,omitempty"`
	// If true, interface LDP sync is achieved
	InterfaceLdpSync bool `protobuf:"varint,91,opt,name=interface_ldp_sync,json=interfaceLdpSync" json:"interface_ldp_sync,omitempty"`
}

func (m *Ospfv3EdmInterface) Reset()                    { *m = Ospfv3EdmInterface{} }
func (m *Ospfv3EdmInterface) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterface) ProtoMessage()               {}
func (*Ospfv3EdmInterface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmInterface) GetInterfaceState() string {
	if m != nil {
		return m.InterfaceState
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetIsInterfaceLineUp() bool {
	if m != nil {
		return m.IsInterfaceLineUp
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceIpSecurityRequired() bool {
	if m != nil {
		return m.IsInterfaceIpSecurityRequired
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceIpSecurityActive() bool {
	if m != nil {
		return m.IsInterfaceIpSecurityActive
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfaceNumber() uint32 {
	if m != nil {
		return m.InterfaceNumber
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceRouterId() string {
	if m != nil {
		return m.InterfaceRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetNetworkType() string {
	if m != nil {
		return m.NetworkType
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceFloodReduction() bool {
	if m != nil {
		return m.IsInterfaceFloodReduction
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsDemandCircuitConfigured() bool {
	if m != nil {
		return m.IsDemandCircuitConfigured
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceDemandCircuit() bool {
	if m != nil {
		return m.IsInterfaceDemandCircuit
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceDcBitlessLsAs() uint32 {
	if m != nil {
		return m.InterfaceDcBitlessLsAs
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetTransmissionDelay() uint32 {
	if m != nil {
		return m.TransmissionDelay
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfacePriority() uint32 {
	if m != nil {
		return m.InterfacePriority
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsDesignatedRouter() bool {
	if m != nil {
		return m.IsDesignatedRouter
	}
	return false
}

func (m *Ospfv3EdmInterface) GetDesignatedRouterId() string {
	if m != nil {
		return m.DesignatedRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetDesignatedRouterAddress() string {
	if m != nil {
		return m.DesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetBackupDesignatedRouterId() string {
	if m != nil {
		return m.BackupDesignatedRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetBackupDesignatedRouterAddress() string {
	if m != nil {
		return m.BackupDesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetNetworkLsaFlushTimer() uint32 {
	if m != nil {
		return m.NetworkLsaFlushTimer
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceLsaFiltered() bool {
	if m != nil {
		return m.IsInterfaceLsaFiltered
	}
	return false
}

func (m *Ospfv3EdmInterface) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetDeadInterval() uint32 {
	if m != nil {
		return m.DeadInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetWaitInterval() uint32 {
	if m != nil {
		return m.WaitInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceRetransmissionInterval() uint32 {
	if m != nil {
		return m.InterfaceRetransmissionInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetNextHelloTime() uint32 {
	if m != nil {
		return m.NextHelloTime
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceAuthenticationSpi() uint32 {
	if m != nil {
		return m.InterfaceAuthenticationSpi
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceAuthenticationTransmit() uint32 {
	if m != nil {
		return m.InterfaceAuthenticationTransmit
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceEncryptionEnabled() bool {
	if m != nil {
		return m.IsInterfaceEncryptionEnabled
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsPrefixSuppress() bool {
	if m != nil {
		return m.IsPrefixSuppress
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptionSpi() uint32 {
	if m != nil {
		return m.InterfaceEncryptionSpi
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptionTransmitted() uint32 {
	if m != nil {
		return m.InterfaceEncryptionTransmitted
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptedAuthenticationTransmitted() uint32 {
	if m != nil {
		return m.InterfaceEncryptedAuthenticationTransmitted
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceNeighborList() []*Ospfv3EdmInterfaceNbr {
	if m != nil {
		return m.InterfaceNeighborList
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetAdjacentNeighbor() uint32 {
	if m != nil {
		return m.AdjacentNeighbor
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetActiveInterface() *Ospfv3EdmInterfaceUp {
	if m != nil {
		return m.ActiveInterface
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetInterfaceBfd() *Ospfv3EdmInterfaceBfd {
	if m != nil {
		return m.InterfaceBfd
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetInterfaceReferences() uint32 {
	if m != nil {
		return m.InterfaceReferences
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetConfiguredLdpSync() bool {
	if m != nil {
		return m.ConfiguredLdpSync
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceLdpSync() bool {
	if m != nil {
		return m.InterfaceLdpSync
	}
	return false
}

// OSPFv3 interface neighbor information
type Ospfv3EdmInterfaceNbr struct {
	// Neighbor router ID
	InterfaceNeighborId string `protobuf:"bytes,1,opt,name=interface_neighbor_id,json=interfaceNeighborId" json:"interface_neighbor_id,omitempty"`
	// Cost of link to neighbor
	InterfaceNeighborCost uint32 `protobuf:"varint,2,opt,name=interface_neighbor_cost,json=interfaceNeighborCost" json:"interface_neighbor_cost,omitempty"`
	// If true, designated router is found
	IsNeighborDr bool `protobuf:"varint,3,opt,name=is_neighbor_dr,json=isNeighborDr" json:"is_neighbor_dr,omitempty"`
	// If true, backup designated router is found
	IsNeighborBdr bool `protobuf:"varint,4,opt,name=is_neighbor_bdr,json=isNeighborBdr" json:"is_neighbor_bdr,omitempty"`
	// If true, hello is suppressed
	IsHelloSuppressed bool `protobuf:"varint,5,opt,name=is_hello_suppressed,json=isHelloSuppressed" json:"is_hello_suppressed,omitempty"`
}

func (m *Ospfv3EdmInterfaceNbr) Reset()                    { *m = Ospfv3EdmInterfaceNbr{} }
func (m *Ospfv3EdmInterfaceNbr) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceNbr) ProtoMessage()               {}
func (*Ospfv3EdmInterfaceNbr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3EdmInterfaceNbr) GetInterfaceNeighborId() string {
	if m != nil {
		return m.InterfaceNeighborId
	}
	return ""
}

func (m *Ospfv3EdmInterfaceNbr) GetInterfaceNeighborCost() uint32 {
	if m != nil {
		return m.InterfaceNeighborCost
	}
	return 0
}

func (m *Ospfv3EdmInterfaceNbr) GetIsNeighborDr() bool {
	if m != nil {
		return m.IsNeighborDr
	}
	return false
}

func (m *Ospfv3EdmInterfaceNbr) GetIsNeighborBdr() bool {
	if m != nil {
		return m.IsNeighborBdr
	}
	return false
}

func (m *Ospfv3EdmInterfaceNbr) GetIsHelloSuppressed() bool {
	if m != nil {
		return m.IsHelloSuppressed
	}
	return false
}

// OSPFv3 interface up-only information
type Ospfv3EdmInterfaceUp struct {
	// Wait time for DR/BDR selection (s)
	WaitTime uint32 `protobuf:"varint,1,opt,name=wait_time,json=waitTime" json:"wait_time,omitempty"`
	// Area scope LSAs flood index
	InterfaceAreaFloodIndex uint32 `protobuf:"varint,2,opt,name=interface_area_flood_index,json=interfaceAreaFloodIndex" json:"interface_area_flood_index,omitempty"`
	// AS scope LSAs flood index
	InterfaceAsFloodIndex uint32 `protobuf:"varint,3,opt,name=interface_as_flood_index,json=interfaceAsFloodIndex" json:"interface_as_flood_index,omitempty"`
	// Interface flood link index
	InterfaceLinkFloodIndex uint32 `protobuf:"varint,4,opt,name=interface_link_flood_index,json=interfaceLinkFloodIndex" json:"interface_link_flood_index,omitempty"`
	// Flood queue length
	FloodQueueLength uint32 `protobuf:"varint,5,opt,name=flood_queue_length,json=floodQueueLength" json:"flood_queue_length,omitempty"`
	// Next LSA to flood (Area scope)
	InterfaceAreaNextFlood uint32 `protobuf:"varint,6,opt,name=interface_area_next_flood,json=interfaceAreaNextFlood" json:"interface_area_next_flood,omitempty"`
	// Index of next LSA to flood (Area scope)
	InterfaceAreaNextFloodIndex uint32 `protobuf:"varint,7,opt,name=interface_area_next_flood_index,json=interfaceAreaNextFloodIndex" json:"interface_area_next_flood_index,omitempty"`
	// Next LSA to flood (AS scope)
	InterfaceAsNextFlood uint32 `protobuf:"varint,8,opt,name=interface_as_next_flood,json=interfaceAsNextFlood" json:"interface_as_next_flood,omitempty"`
	// Index of next LSA to flood (AS scope)
	InterfaceAsNextFloodIndex uint32 `protobuf:"varint,9,opt,name=interface_as_next_flood_index,json=interfaceAsNextFloodIndex" json:"interface_as_next_flood_index,omitempty"`
	// Interface link next flood information
	InterfaceLinkNextFlood uint32 `protobuf:"varint,10,opt,name=interface_link_next_flood,json=interfaceLinkNextFlood" json:"interface_link_next_flood,omitempty"`
	// Interface link next information index
	InterfaceLinkNextIndex uint32 `protobuf:"varint,11,opt,name=interface_link_next_index,json=interfaceLinkNextIndex" json:"interface_link_next_index,omitempty"`
	// Last flood scan length
	FloodScanLength uint32 `protobuf:"varint,12,opt,name=flood_scan_length,json=floodScanLength" json:"flood_scan_length,omitempty"`
	// Maximum flood length
	MaximumFloodLength uint32 `protobuf:"varint,13,opt,name=maximum_flood_length,json=maximumFloodLength" json:"maximum_flood_length,omitempty"`
	// Last flood scan time (ms)
	LastFloodTime uint32 `protobuf:"varint,14,opt,name=last_flood_time,json=lastFloodTime" json:"last_flood_time,omitempty"`
	// Maximum flood time (ms)
	MaximumFloodTime uint32 `protobuf:"varint,15,opt,name=maximum_flood_time,json=maximumFloodTime" json:"maximum_flood_time,omitempty"`
	// Time until next flood pacing timer (ms)
	InterfaceFloodPacingTimer uint32 `protobuf:"varint,16,opt,name=interface_flood_pacing_timer,json=interfaceFloodPacingTimer" json:"interface_flood_pacing_timer,omitempty"`
	// Total number of neighbors
	InterfaceNeighbors uint32 `protobuf:"varint,17,opt,name=interface_neighbors,json=interfaceNeighbors" json:"interface_neighbors,omitempty"`
	// Number of neighbors for which hellos are suppressed
	SuppressedHellos uint32 `protobuf:"varint,18,opt,name=suppressed_hellos,json=suppressedHellos" json:"suppressed_hellos,omitempty"`
}

func (m *Ospfv3EdmInterfaceUp) Reset()                    { *m = Ospfv3EdmInterfaceUp{} }
func (m *Ospfv3EdmInterfaceUp) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceUp) ProtoMessage()               {}
func (*Ospfv3EdmInterfaceUp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ospfv3EdmInterfaceUp) GetWaitTime() uint32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAreaFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAsFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceLinkFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetFloodQueueLength() uint32 {
	if m != nil {
		return m.FloodQueueLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaNextFlood() uint32 {
	if m != nil {
		return m.InterfaceAreaNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaNextFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAreaNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsNextFlood() uint32 {
	if m != nil {
		return m.InterfaceAsNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsNextFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAsNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkNextFlood() uint32 {
	if m != nil {
		return m.InterfaceLinkNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkNextIndex() uint32 {
	if m != nil {
		return m.InterfaceLinkNextIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetFloodScanLength() uint32 {
	if m != nil {
		return m.FloodScanLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetMaximumFloodLength() uint32 {
	if m != nil {
		return m.MaximumFloodLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetLastFloodTime() uint32 {
	if m != nil {
		return m.LastFloodTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetMaximumFloodTime() uint32 {
	if m != nil {
		return m.MaximumFloodTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceFloodPacingTimer() uint32 {
	if m != nil {
		return m.InterfaceFloodPacingTimer
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceNeighbors() uint32 {
	if m != nil {
		return m.InterfaceNeighbors
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetSuppressedHellos() uint32 {
	if m != nil {
		return m.SuppressedHellos
	}
	return 0
}

// Interface BFD information
type Ospfv3EdmInterfaceBfd struct {
	// BFD Enable Mode on the interface - Default/Strict
	BfdIntfEnableMode uint32 `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode" json:"bfd_intf_enable_mode,omitempty"`
	// BFD interval (ms)
	BfdInterval uint32 `protobuf:"varint,2,opt,name=bfd_interval,json=bfdInterval" json:"bfd_interval,omitempty"`
	// BFD detection multiplier
	BfdDetectionMultiplier uint32 `protobuf:"varint,3,opt,name=bfd_detection_multiplier,json=bfdDetectionMultiplier" json:"bfd_detection_multiplier,omitempty"`
}

func (m *Ospfv3EdmInterfaceBfd) Reset()                    { *m = Ospfv3EdmInterfaceBfd{} }
func (m *Ospfv3EdmInterfaceBfd) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceBfd) ProtoMessage()               {}
func (*Ospfv3EdmInterfaceBfd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ospfv3EdmInterfaceBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *Ospfv3EdmInterfaceBfd) GetBfdInterval() uint32 {
	if m != nil {
		return m.BfdInterval
	}
	return 0
}

func (m *Ospfv3EdmInterfaceBfd) GetBfdDetectionMultiplier() uint32 {
	if m != nil {
		return m.BfdDetectionMultiplier
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmInterface_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.interface_table.interface.ospfv3_edm_interface_KEYS")
	proto.RegisterType((*Ospfv3EdmInterface)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.interface_table.interface.ospfv3_edm_interface")
	proto.RegisterType((*Ospfv3EdmInterfaceNbr)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.interface_table.interface.ospfv3_edm_interface_nbr")
	proto.RegisterType((*Ospfv3EdmInterfaceUp)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.interface_table.interface.ospfv3_edm_interface_up")
	proto.RegisterType((*Ospfv3EdmInterfaceBfd)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.interface_table.interface.ospfv3_edm_interface_bfd")
}

func init() { proto.RegisterFile("ospfv3_edm_interface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x58, 0x5f, 0x57, 0x14, 0xc9,
	0x15, 0x3f, 0xa3, 0x46, 0xb0, 0x00, 0x81, 0x12, 0xa5, 0x10, 0x8d, 0x48, 0x12, 0x43, 0xa2, 0x8e,
	0x06, 0xe3, 0xff, 0x18, 0x05, 0x06, 0x64, 0x74, 0x24, 0x38, 0x83, 0x49, 0x4c, 0x1e, 0xea, 0xd4,
	0x74, 0x55, 0x43, 0x2d, 0x3d, 0xdd, 0x6d, 0x55, 0x35, 0xc2, 0xc3, 0xee, 0xd3, 0x7e, 0x89, 0x3d,
	0xfb, 0xb8, 0x2f, 0x7e, 0x82, 0xfd, 0x30, 0xfb, 0x69, 0xf6, 0xd4, 0xad, 0xea, 0xe9, 0xee, 0x61,
	0xe6, 0x75, 0x7d, 0xf1, 0x30, 0xf7, 0xfe, 0x7e, 0xf7, 0x5f, 0xdd, 0xba, 0x75, 0x5b, 0x74, 0x35,
	0xd1, 0x69, 0x78, 0xf4, 0x80, 0x0a, 0xde, 0xa3, 0x32, 0x36, 0x42, 0x85, 0x2c, 0x10, 0xf5, 0x54,
	0x25, 0x26, 0xc1, 0x3c, 0x90, 0x3a, 0x48, 0xa8, 0x4c, 0x34, 0x3d, 0x56, 0x54, 0xa6, 0x47, 0x8f,
	0xa8, 0x47, 0x27, 0xa9, 0x50, 0x75, 0xf7, 0xb7, 0xc5, 0x06, 0x42, 0x6b, 0xa1, 0xf3, 0xbf, 0xea,
	0x47, 0x2a, 0x84, 0x7f, 0xea, 0x4c, 0x09, 0xa6, 0xe1, 0xdf, 0x7a, 0xdf, 0x36, 0x35, 0xac, 0x1b,
	0x89, 0xe2, 0xf7, 0xf2, 0x0f, 0x35, 0xb4, 0x30, 0x2c, 0x08, 0xfa, 0x76, 0xf3, 0x63, 0x07, 0xdf,
	0x44, 0x93, 0xde, 0x2c, 0x8d, 0x59, 0x4f, 0x90, 0xda, 0x52, 0x6d, 0xe5, 0x42, 0x7b, 0xc2, 0xcb,
	0x76, 0x58, 0x4f, 0xe0, 0x05, 0x34, 0x7e, 0xa4, 0x42, 0xa7, 0x3e, 0x03, 0xea, 0xb1, 0x23, 0x15,
	0x82, 0x6a, 0x1e, 0x8d, 0x59, 0xef, 0x54, 0x72, 0x72, 0x76, 0xa9, 0xb6, 0x32, 0xd5, 0x3e, 0x6f,
	0x7f, 0x36, 0x39, 0xfe, 0x13, 0xba, 0x58, 0x38, 0x02, 0xe6, 0x39, 0x60, 0x4e, 0xf5, 0xa5, 0x96,
	0xbf, 0xfc, 0xe3, 0x65, 0x34, 0x37, 0x2c, 0x36, 0xfc, 0x67, 0x34, 0x5d, 0xf0, 0xb5, 0x61, 0x46,
	0x90, 0x55, 0x30, 0x50, 0x98, 0xed, 0x58, 0x29, 0xbe, 0x87, 0xe6, 0xa4, 0x2e, 0x25, 0x15, 0xc9,
	0x58, 0xd0, 0x2c, 0x25, 0x0f, 0x96, 0x6a, 0x2b, 0xe3, 0xed, 0x59, 0xa9, 0x9b, 0xb9, 0xaa, 0x25,
	0x63, 0xf1, 0x21, 0xc5, 0xdb, 0xe8, 0x66, 0x85, 0x20, 0x53, 0xaa, 0x45, 0x90, 0x29, 0x69, 0x4e,
	0xa8, 0x12, 0x9f, 0x32, 0xa9, 0x04, 0x27, 0x7f, 0x07, 0xf6, 0xf5, 0x12, 0xbb, 0x99, 0x76, 0x3c,
	0xaa, 0xed, 0x41, 0xb8, 0x81, 0x6e, 0x8c, 0xb4, 0xc4, 0x02, 0x23, 0x8f, 0x04, 0x79, 0x08, 0x76,
	0x16, 0x87, 0xda, 0x59, 0x03, 0x08, 0xbe, 0x8d, 0x66, 0x0b, 0x13, 0x8c, 0x73, 0x25, 0xb4, 0x26,
	0x8f, 0x20, 0xd7, 0x99, 0xbe, 0x62, 0xcd, 0xc9, 0xf1, 0x5f, 0xd0, 0x4c, 0xa9, 0xac, 0x59, 0xaf,
	0x2b, 0x14, 0x79, 0x0c, 0x85, 0x2f, 0xca, 0xb5, 0x03, 0x62, 0x5c, 0x47, 0x97, 0x0a, 0xa8, 0x4a,
	0x32, 0x23, 0x94, 0x3d, 0xa6, 0x27, 0x60, 0xb9, 0x70, 0xd9, 0x06, 0x4d, 0x93, 0xdb, 0x46, 0x88,
	0x85, 0xf9, 0x9c, 0xa8, 0x43, 0x6a, 0x4e, 0x52, 0x41, 0x9e, 0xba, 0x46, 0xf0, 0xb2, 0xbd, 0x93,
	0x54, 0x54, 0x4d, 0x46, 0x32, 0x3e, 0xa4, 0x41, 0xa2, 0x0d, 0x79, 0x06, 0x01, 0x14, 0x26, 0x5b,
	0x32, 0x3e, 0xdc, 0x48, 0xb4, 0xc1, 0x2f, 0xd1, 0xb5, 0x4a, 0x81, 0xc2, 0x28, 0x49, 0x38, 0x55,
	0x82, 0x67, 0x81, 0x91, 0x49, 0x4c, 0x9e, 0x43, 0x75, 0x16, 0x4a, 0xd5, 0xd9, 0xb2, 0x88, 0x76,
	0x0e, 0xf0, 0x06, 0xb8, 0xe8, 0xb1, 0x98, 0xd3, 0x40, 0xaa, 0x20, 0x93, 0x86, 0x06, 0x49, 0x1c,
	0xca, 0xfd, 0xcc, 0x1e, 0xd3, 0x3f, 0x72, 0x03, 0x0d, 0x80, 0x6c, 0x38, 0xc4, 0x46, 0x1f, 0x80,
	0x5f, 0xa0, 0xc5, 0x4a, 0x04, 0x55, 0x53, 0xe4, 0x05, 0xf0, 0x49, 0x29, 0x80, 0x8a, 0x21, 0xfc,
	0x0c, 0x5d, 0x2d, 0x71, 0x03, 0xda, 0x95, 0x26, 0xb2, 0x37, 0x25, 0xd2, 0x94, 0x69, 0xf2, 0x4f,
	0xc8, 0xfb, 0x4a, 0x1f, 0xd1, 0x08, 0xd6, 0x9d, 0xbe, 0xa5, 0xd7, 0x34, 0xbe, 0x8b, 0xb0, 0x51,
	0x2c, 0xd6, 0x3d, 0xa9, 0xb5, 0x4c, 0x62, 0xca, 0x45, 0xc4, 0x4e, 0xc8, 0x4b, 0x57, 0xab, 0xb2,
	0xa6, 0x61, 0x15, 0xf8, 0xbe, 0xbb, 0x08, 0x74, 0xb0, 0xeb, 0x5f, 0xc1, 0x31, 0x60, 0xab, 0x6b,
	0x56, 0x3b, 0xff, 0x2e, 0xc2, 0x05, 0x38, 0x55, 0x32, 0xb1, 0x4d, 0x45, 0xd6, 0x06, 0x0e, 0x63,
	0xd7, 0x2b, 0xac, 0x03, 0xa8, 0xa5, 0x96, 0xfb, 0x31, 0x33, 0x82, 0xfb, 0x9e, 0x20, 0xeb, 0x50,
	0x03, 0x6c, 0x6b, 0x98, 0xab, 0x5c, 0x4f, 0x58, 0xc6, 0x29, 0xb8, 0x6d, 0xa1, 0x0d, 0x17, 0x12,
	0x1f, 0xc0, 0x37, 0x39, 0x7e, 0x86, 0x16, 0x4e, 0x33, 0xf2, 0x9e, 0x6e, 0x00, 0x6d, 0x7e, 0x90,
	0x96, 0xb7, 0xf6, 0x0b, 0xb4, 0xd8, 0x65, 0xc1, 0x61, 0x96, 0xd2, 0xa1, 0x4e, 0x37, 0x81, 0x4d,
	0x1c, 0xa4, 0x71, 0xda, 0xf5, 0x6b, 0xb4, 0x34, 0x92, 0x9e, 0x47, 0xb0, 0x05, 0x36, 0xae, 0x0f,
	0xb7, 0x91, 0xc7, 0xf1, 0x10, 0xcd, 0xe7, 0xf7, 0x20, 0xd2, 0x8c, 0x86, 0x51, 0xa6, 0x0f, 0xa8,
	0x91, 0x3d, 0xa1, 0xc8, 0x6b, 0xa8, 0xed, 0x9c, 0x57, 0xb7, 0x34, 0xdb, 0xb2, 0xca, 0x3d, 0xab,
	0xc3, 0x4f, 0xd1, 0x42, 0x75, 0x0e, 0x59, 0xae, 0x8c, 0x8c, 0xb0, 0x7d, 0xba, 0x0d, 0x35, 0xbe,
	0x52, 0x1e, 0x46, 0x9a, 0x6d, 0x79, 0xad, 0x9d, 0x95, 0x07, 0x22, 0x8a, 0x12, 0xc7, 0x3e, 0x62,
	0x11, 0x69, 0x82, 0xa3, 0x29, 0x90, 0x36, 0xbd, 0x10, 0xff, 0x01, 0x4d, 0x71, 0xc1, 0x78, 0x81,
	0x7a, 0x03, 0xa8, 0x49, 0x2b, 0x2c, 0x83, 0x3e, 0x33, 0x69, 0x0a, 0xd0, 0x5b, 0x07, 0xb2, 0xc2,
	0x3e, 0xe8, 0x0d, 0xba, 0x59, 0x1a, 0x0d, 0xa2, 0xd2, 0xa6, 0x7d, 0x62, 0x0b, 0x88, 0x37, 0x8a,
	0x41, 0x51, 0xc1, 0xf5, 0x6d, 0xdd, 0x42, 0xd3, 0xb1, 0x38, 0x36, 0xd4, 0x65, 0x60, 0xeb, 0x44,
	0xde, 0xb9, 0xe8, 0xad, 0x78, 0xdb, 0x4a, 0x6d, 0x81, 0xf0, 0x2b, 0x74, 0xad, 0x34, 0xe6, 0x32,
	0x73, 0x20, 0x62, 0x23, 0x03, 0x66, 0xaf, 0x39, 0xd5, 0xa9, 0x24, 0x3b, 0x40, 0x2a, 0xae, 0xdb,
	0x5a, 0x05, 0xd2, 0x49, 0x65, 0x35, 0xea, 0x01, 0x0b, 0x3e, 0x34, 0x43, 0xfe, 0x35, 0x10, 0x75,
	0xd5, 0xcc, 0x9e, 0x87, 0xe1, 0xcd, 0x81, 0xd1, 0x2d, 0xe2, 0x40, 0x9d, 0xa4, 0x60, 0x4a, 0xc4,
	0xf6, 0xfd, 0xe4, 0x64, 0x17, 0xce, 0xec, 0x5a, 0xe9, 0xcc, 0x36, 0xfb, 0xa0, 0x4d, 0x87, 0xc1,
	0x77, 0x10, 0x96, 0x9a, 0xa6, 0x4a, 0x84, 0xf2, 0x98, 0xea, 0x2c, 0x4d, 0xa1, 0xcd, 0xde, 0x03,
	0x73, 0x46, 0xea, 0x5d, 0x50, 0x74, 0xbc, 0x1c, 0x3f, 0x41, 0x64, 0xa8, 0x47, 0x9b, 0x7e, 0x7b,
	0x60, 0x96, 0x14, 0xbe, 0x6c, 0xea, 0xdb, 0x68, 0x69, 0x28, 0x33, 0x4f, 0xdb, 0x08, 0x4e, 0x3a,
	0x60, 0xe1, 0xf7, 0x43, 0x2c, 0xec, 0x15, 0x28, 0x1c, 0xa0, 0xfa, 0x29, 0x4b, 0x82, 0x8f, 0x2a,
	0xa7, 0xb5, 0xbb, 0x07, 0x76, 0x6f, 0x0f, 0xda, 0x15, 0x7c, 0x78, 0x69, 0xad, 0x93, 0x9f, 0x6b,
	0x68, 0xbe, 0xf4, 0x4c, 0x09, 0xb9, 0x7f, 0xd0, 0x4d, 0x14, 0x8d, 0xa4, 0x36, 0xe4, 0xc3, 0xd2,
	0xd9, 0x95, 0x89, 0xd5, 0xef, 0xea, 0xbf, 0xc5, 0xea, 0x53, 0x1f, 0xba, 0xf6, 0xc4, 0x5d, 0xd5,
	0xbe, 0x5c, 0xbc, 0x96, 0x3e, 0xba, 0x96, 0xd4, 0xc6, 0xbe, 0xc5, 0x8c, 0x7f, 0xc3, 0x02, 0x11,
	0x9b, 0x7e, 0xd8, 0xe4, 0xdf, 0x50, 0x80, 0x99, 0x5c, 0x91, 0x13, 0xf0, 0x97, 0x1a, 0x9a, 0x71,
	0xcf, 0x7c, 0x61, 0x9c, 0xfc, 0x67, 0xa9, 0xb6, 0x32, 0xb1, 0xfa, 0xed, 0x57, 0x4c, 0x2f, 0x4b,
	0xdb, 0xd3, 0x2e, 0xac, 0x7e, 0xe3, 0xe2, 0x9f, 0x6a, 0xa8, 0x58, 0xbc, 0x68, 0x37, 0xe4, 0xe4,
	0xbf, 0x10, 0xe7, 0xd7, 0x3c, 0x86, 0x6e, 0xc8, 0xdb, 0x93, 0xfd, 0x9f, 0xeb, 0x21, 0xc7, 0x7f,
	0x43, 0x73, 0xe5, 0xb1, 0x14, 0x0a, 0x25, 0xe2, 0x40, 0x68, 0xf2, 0x11, 0x0e, 0xe0, 0x52, 0x69,
	0x12, 0xe5, 0x2a, 0xbb, 0x91, 0x14, 0xeb, 0x00, 0x8d, 0x78, 0x4a, 0xf5, 0x49, 0x1c, 0x90, 0xff,
	0xb9, 0xe5, 0xaf, 0x50, 0xb5, 0x78, 0xda, 0x39, 0x89, 0x03, 0xb8, 0xb0, 0xc5, 0x88, 0xce, 0xe1,
	0xff, 0xf7, 0x17, 0xb6, 0x3f, 0x9c, 0x1d, 0x7a, 0xf9, 0xfb, 0x33, 0x88, 0x8c, 0x6a, 0x21, 0xbc,
	0x8a, 0x2e, 0x0f, 0xe9, 0x71, 0xc9, 0xfd, 0x06, 0x7d, 0xe9, 0x54, 0x87, 0x35, 0x39, 0x7e, 0x34,
	0xf4, 0x5e, 0xc0, 0x12, 0x75, 0x06, 0x92, 0x3c, 0xdd, 0x97, 0xb0, 0x48, 0xfd, 0x11, 0x5d, 0x94,
	0xba, 0x20, 0x70, 0x05, 0xdb, 0xf6, 0x78, 0x7b, 0x52, 0xea, 0x1c, 0xd7, 0x50, 0x76, 0x14, 0x97,
	0x51, 0x5d, 0xae, 0x60, 0xe9, 0x1e, 0x6f, 0x4f, 0x15, 0xb0, 0x75, 0xee, 0x36, 0x43, 0xed, 0x07,
	0x76, 0x3e, 0xb4, 0x04, 0x27, 0xbf, 0xcb, 0x37, 0x66, 0x18, 0xda, 0x9d, 0xbe, 0x62, 0xf9, 0x97,
	0x31, 0x34, 0x3f, 0xa2, 0xd5, 0xf0, 0x22, 0xba, 0x00, 0xef, 0x0d, 0x0c, 0xfe, 0x1a, 0xe4, 0x30,
	0x6e, 0x05, 0x30, 0xf3, 0x9f, 0x97, 0xd7, 0x27, 0xf8, 0x4e, 0x70, 0x1b, 0xa0, 0x8c, 0xb9, 0x38,
	0xf6, 0x19, 0x17, 0x05, 0x59, 0x53, 0x82, 0xc1, 0xfe, 0xd7, 0xb4, 0x6a, 0xfc, 0xb8, 0x3c, 0x2d,
	0x99, 0xae, 0x50, 0xcf, 0x0e, 0x14, 0x6b, 0x4d, 0x97, 0x88, 0x15, 0xaf, 0xb0, 0xa5, 0x96, 0xa9,
	0xe7, 0x06, 0xbc, 0xda, 0x65, 0xb5, 0x44, 0xbe, 0x83, 0xb0, 0x43, 0x7f, 0xca, 0x44, 0x26, 0x68,
	0x24, 0xe2, 0x7d, 0x73, 0x00, 0xa5, 0x99, 0x6a, 0xcf, 0x80, 0xe6, 0xbd, 0x55, 0xb4, 0x40, 0x0e,
	0x8f, 0x7e, 0x35, 0x41, 0x78, 0x0b, 0x01, 0x47, 0xce, 0x0f, 0x8c, 0x74, 0x9b, 0xdf, 0x8e, 0x38,
	0x36, 0xe0, 0x0d, 0x3e, 0x1e, 0x46, 0x51, 0x7d, 0xa8, 0x63, 0x60, 0x60, 0x71, 0xb8, 0x01, 0x17,
	0xee, 0xc3, 0x72, 0x43, 0x31, 0x5d, 0x76, 0x3f, 0xee, 0x96, 0x95, 0x52, 0x8d, 0x0a, 0xe7, 0xaf,
	0xd0, 0xf5, 0x11, 0x34, 0xef, 0xfa, 0x02, 0x90, 0x17, 0x86, 0x91, 0x9d, 0xe3, 0x4a, 0xe6, 0x50,
	0xe4, 0x92, 0x6b, 0x34, 0x90, 0xb9, 0xad, 0x71, 0xe1, 0x7c, 0x04, 0xd5, 0x39, 0x9e, 0x18, 0x41,
	0x75, 0x5e, 0xff, 0x8a, 0x66, 0x5d, 0x94, 0x3a, 0x60, 0x71, 0x7e, 0x38, 0x93, 0xee, 0xfb, 0x07,
	0x14, 0x9d, 0x80, 0xc5, 0xfe, 0x6c, 0xee, 0xa3, 0xb9, 0x1e, 0x3b, 0x96, 0xbd, 0xac, 0xe7, 0x33,
	0xf3, 0xf0, 0x29, 0x80, 0x63, 0xaf, 0x83, 0x90, 0x3c, 0xe3, 0x16, 0x9a, 0x8e, 0x98, 0xce, 0x0b,
	0x01, 0x1d, 0x7d, 0xd1, 0xad, 0x32, 0x56, 0x0c, 0x48, 0x68, 0xeb, 0x3b, 0x08, 0x57, 0x2d, 0x03,
	0x74, 0xda, 0xf5, 0x48, 0xd9, 0x2e, 0xa0, 0x5f, 0x96, 0x17, 0x1f, 0x87, 0x4f, 0x59, 0x20, 0xe3,
	0x7d, 0xbf, 0x54, 0xce, 0x0c, 0x94, 0x1a, 0x98, 0xbb, 0x80, 0x70, 0x9b, 0xe5, 0xbd, 0xf2, 0x57,
	0x57, 0x7e, 0xbb, 0x35, 0x99, 0x75, 0x79, 0x9c, 0x1a, 0x18, 0xda, 0xbe, 0x62, 0xc5, 0xb5, 0x76,
	0xf7, 0x5c, 0x13, 0xec, 0xc2, 0x2b, 0x14, 0x70, 0xcb, 0xf5, 0xf2, 0x97, 0xda, 0x88, 0x19, 0xd7,
	0x0d, 0xb9, 0xfd, 0xb8, 0xee, 0x86, 0xb0, 0x71, 0x86, 0x7e, 0x2f, 0xa2, 0xbd, 0x84, 0xe7, 0x17,
	0x7d, 0xb6, 0x1b, 0xda, 0xc5, 0x33, 0x74, 0xdb, 0xd0, 0xbb, 0x84, 0x0b, 0xfb, 0x11, 0xe9, 0x09,
	0x6e, 0x89, 0x74, 0x77, 0x7c, 0xc2, 0x01, 0xdd, 0xc2, 0xf8, 0x04, 0x11, 0x0b, 0xe1, 0xc2, 0x08,
	0xf8, 0xc8, 0xa3, 0xbd, 0x2c, 0x32, 0x32, 0x8d, 0xa4, 0x50, 0xfe, 0x5e, 0x5f, 0xe9, 0x86, 0xbc,
	0x91, 0xab, 0xdf, 0xf5, 0xb5, 0xdd, 0xf3, 0xf0, 0xbf, 0x26, 0x0f, 0x7e, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xda, 0x9e, 0xc1, 0x5d, 0x53, 0x11, 0x00, 0x00,
}