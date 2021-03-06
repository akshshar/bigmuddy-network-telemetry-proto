// Code generated by protoc-gen-go.
// source: ipv6_if_summary.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_summary is a generated protocol buffer package.

It is generated from these files:
	ipv6_if_summary.proto

It has these top-level messages:
	Ipv6IfSummary_KEYS
	Ipv6IfSummary
	IfSummaryT
*/
package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_summary

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

// Summary info of IPv6 interfaces
type Ipv6IfSummary_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *Ipv6IfSummary_KEYS) Reset()                    { *m = Ipv6IfSummary_KEYS{} }
func (m *Ipv6IfSummary_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6IfSummary_KEYS) ProtoMessage()               {}
func (*Ipv6IfSummary_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6IfSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6IfSummary struct {
	// Number of interfaces (up,up)
	IfUpUp *IfSummaryT `protobuf:"bytes,50,opt,name=if_up_up,json=ifUpUp" json:"if_up_up,omitempty"`
	// Number of interfaces (up,down)
	IfUpDown *IfSummaryT `protobuf:"bytes,51,opt,name=if_up_down,json=ifUpDown" json:"if_up_down,omitempty"`
	// Number of interfaces (down,down)
	IfDownDown *IfSummaryT `protobuf:"bytes,52,opt,name=if_down_down,json=ifDownDown" json:"if_down_down,omitempty"`
	// Number of interfaces (shutdown,down)
	IfShutdownDown *IfSummaryT `protobuf:"bytes,53,opt,name=if_shutdown_down,json=ifShutdownDown" json:"if_shutdown_down,omitempty"`
	// Number of interfaces (up,down) with basecaps up
	IfUpDownBasecapsUp uint32 `protobuf:"varint,54,opt,name=if_up_down_basecaps_up,json=ifUpDownBasecapsUp" json:"if_up_down_basecaps_up,omitempty"`
}

func (m *Ipv6IfSummary) Reset()                    { *m = Ipv6IfSummary{} }
func (m *Ipv6IfSummary) String() string            { return proto.CompactTextString(m) }
func (*Ipv6IfSummary) ProtoMessage()               {}
func (*Ipv6IfSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6IfSummary) GetIfUpUp() *IfSummaryT {
	if m != nil {
		return m.IfUpUp
	}
	return nil
}

func (m *Ipv6IfSummary) GetIfUpDown() *IfSummaryT {
	if m != nil {
		return m.IfUpDown
	}
	return nil
}

func (m *Ipv6IfSummary) GetIfDownDown() *IfSummaryT {
	if m != nil {
		return m.IfDownDown
	}
	return nil
}

func (m *Ipv6IfSummary) GetIfShutdownDown() *IfSummaryT {
	if m != nil {
		return m.IfShutdownDown
	}
	return nil
}

func (m *Ipv6IfSummary) GetIfUpDownBasecapsUp() uint32 {
	if m != nil {
		return m.IfUpDownBasecapsUp
	}
	return 0
}

// Count of assigned/unnumbered interfaces
type IfSummaryT struct {
	// Number of interfaces with explicit addresses
	IpAssigned uint32 `protobuf:"varint,1,opt,name=ip_assigned,json=ipAssigned" json:"ip_assigned,omitempty"`
	// Number of unnumbered interfaces with explicit addresses
	IpUnnumbered uint32 `protobuf:"varint,2,opt,name=ip_unnumbered,json=ipUnnumbered" json:"ip_unnumbered,omitempty"`
	// Number of unassigned interfaces without explicit address
	IpUnassigned uint32 `protobuf:"varint,3,opt,name=ip_unassigned,json=ipUnassigned" json:"ip_unassigned,omitempty"`
}

func (m *IfSummaryT) Reset()                    { *m = IfSummaryT{} }
func (m *IfSummaryT) String() string            { return proto.CompactTextString(m) }
func (*IfSummaryT) ProtoMessage()               {}
func (*IfSummaryT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IfSummaryT) GetIpAssigned() uint32 {
	if m != nil {
		return m.IpAssigned
	}
	return 0
}

func (m *IfSummaryT) GetIpUnnumbered() uint32 {
	if m != nil {
		return m.IpUnnumbered
	}
	return 0
}

func (m *IfSummaryT) GetIpUnassigned() uint32 {
	if m != nil {
		return m.IpUnassigned
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6IfSummary_KEYS)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.summary.ipv6_if_summary_KEYS")
	proto.RegisterType((*Ipv6IfSummary)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.summary.ipv6_if_summary")
	proto.RegisterType((*IfSummaryT)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.summary.if_summary_t")
}

func init() { proto.RegisterFile("ipv6_if_summary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x4a, 0x69, 0xa7, 0xad, 0xca, 0xa2, 0x12, 0xf0, 0x60, 0xa9, 0x97, 0x9e, 0x72,
	0x68, 0xb5, 0x77, 0x45, 0x41, 0x10, 0x3c, 0xa4, 0x44, 0xf0, 0x34, 0x6c, 0x93, 0x8d, 0x0e, 0x92,
	0xdd, 0x21, 0x9b, 0xd8, 0xf6, 0xe6, 0xef, 0xf1, 0x57, 0x4a, 0xb6, 0x69, 0x2a, 0x3d, 0xdb, 0x4b,
	0x20, 0x6f, 0xde, 0xcc, 0x37, 0xbc, 0xdd, 0x85, 0x73, 0xe2, 0xaf, 0x29, 0x52, 0x8a, 0xb6, 0xcc,
	0x32, 0x99, 0xaf, 0x02, 0xce, 0x4d, 0x61, 0xc4, 0x53, 0x4c, 0x36, 0x36, 0x48, 0xc6, 0xe2, 0x32,
	0x47, 0xe7, 0xc9, 0x24, 0x1a, 0x56, 0x79, 0xe0, 0x7e, 0xb4, 0x2a, 0x16, 0x26, 0xff, 0x0c, 0xb4,
	0x49, 0x94, 0x75, 0xdf, 0x80, 0x74, 0xa1, 0xf2, 0x54, 0xc6, 0x0a, 0x13, 0x59, 0xc8, 0xa0, 0x9e,
	0x37, 0x9c, 0xc0, 0xd9, 0x0e, 0x02, 0x9f, 0x1f, 0xdf, 0x66, 0xe2, 0x12, 0x3a, 0x55, 0x1b, 0x6a,
	0x99, 0x29, 0xdf, 0x1b, 0x78, 0xa3, 0x4e, 0xd8, 0xae, 0x84, 0x17, 0x99, 0xa9, 0xe1, 0xcf, 0x11,
	0x9c, 0xec, 0x74, 0x09, 0x86, 0x36, 0xa5, 0x58, 0x32, 0x96, 0xec, 0x8f, 0x07, 0xde, 0xa8, 0x3b,
	0x7e, 0x0d, 0xfe, 0x6b, 0xcb, 0xe0, 0xcf, 0x76, 0x45, 0xd8, 0xa2, 0x34, 0xe2, 0x88, 0x45, 0x01,
	0xb0, 0x26, 0x26, 0x66, 0xa1, 0xfd, 0xc9, 0x5e, 0x99, 0xed, 0x8a, 0xf9, 0x60, 0x16, 0x5a, 0x2c,
	0xa1, 0x47, 0xa9, 0x43, 0xae, 0xb9, 0x37, 0x7b, 0xe5, 0x02, 0xa5, 0x15, 0xd5, 0x91, 0xbf, 0x3d,
	0x38, 0xad, 0x8a, 0x1f, 0x65, 0xb1, 0xc5, 0xdf, 0xee, 0x15, 0x7f, 0x4c, 0xe9, 0xac, 0xc6, 0xb9,
	0x15, 0xc6, 0x70, 0xb1, 0x8d, 0x1c, 0xe7, 0xd2, 0xaa, 0x58, 0xb2, 0xad, 0x8e, 0x7c, 0x3a, 0xf0,
	0x46, 0xfd, 0x50, 0x6c, 0x62, 0xba, 0xaf, 0x4b, 0x11, 0x0f, 0x57, 0x2e, 0xb0, 0x66, 0xa6, 0xb8,
	0x82, 0x2e, 0x31, 0x4a, 0x6b, 0xe9, 0x5d, 0xab, 0xc4, 0xdd, 0xad, 0x7e, 0x08, 0xc4, 0x77, 0xb5,
	0x22, 0xae, 0xa1, 0x4f, 0x8c, 0xa5, 0xd6, 0x65, 0x36, 0x57, 0xb9, 0x4a, 0xfc, 0x03, 0x67, 0xe9,
	0x11, 0x47, 0x8d, 0xd6, 0x98, 0x9a, 0x39, 0x87, 0x5b, 0xd3, 0x46, 0x9b, 0xb7, 0xdc, 0x6b, 0x99,
	0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xfe, 0x0e, 0x90, 0x46, 0x03, 0x00, 0x00,
}
