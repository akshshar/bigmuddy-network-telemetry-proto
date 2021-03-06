// Code generated by protoc-gen-go.
// source: ipv6_rib_edm_advert.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_advert.proto

It has these top-level messages:
	Ipv6RibEdmAdvert_KEYS
	Ipv6RibEdmAdvert
	Ipv6RibEdmAdvertItem
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert

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

// Route advertisement information
type Ipv6RibEdmAdvert_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ipv6RibEdmAdvert_KEYS) Reset()                    { *m = Ipv6RibEdmAdvert_KEYS{} }
func (m *Ipv6RibEdmAdvert_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvert_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmAdvert_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmAdvert_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv6RibEdmAdvert struct {
	// Next advertising proto
	Ipv6RibEdmAdvert []*Ipv6RibEdmAdvertItem `protobuf:"bytes,50,rep,name=ipv6_rib_edm_advert,json=ipv6RibEdmAdvert" json:"ipv6_rib_edm_advert,omitempty"`
}

func (m *Ipv6RibEdmAdvert) Reset()                    { *m = Ipv6RibEdmAdvert{} }
func (m *Ipv6RibEdmAdvert) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvert) ProtoMessage()               {}
func (*Ipv6RibEdmAdvert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmAdvert) GetIpv6RibEdmAdvert() []*Ipv6RibEdmAdvertItem {
	if m != nil {
		return m.Ipv6RibEdmAdvert
	}
	return nil
}

type Ipv6RibEdmAdvertItem struct {
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,1,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//   Client advertising the route
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,3,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// Extended communities
	ExtendedCommunities []byte `protobuf:"bytes,4,opt,name=extended_communities,json=extendedCommunities,proto3" json:"extended_communities,omitempty"`
	// OSPF area-id flags
	ProtocolOpaqueFlags uint32 `protobuf:"varint,5,opt,name=protocol_opaque_flags,json=protocolOpaqueFlags" json:"protocol_opaque_flags,omitempty"`
	// OSPF area-id
	ProtocolOpaque uint32 `protobuf:"varint,6,opt,name=protocol_opaque,json=protocolOpaque" json:"protocol_opaque,omitempty"`
	// Protocol code
	Code int32 `protobuf:"zigzag32,7,opt,name=code" json:"code,omitempty"`
	// Instance name
	InstanceName []byte `protobuf:"bytes,8,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (m *Ipv6RibEdmAdvertItem) Reset()                    { *m = Ipv6RibEdmAdvertItem{} }
func (m *Ipv6RibEdmAdvertItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvertItem) ProtoMessage()               {}
func (*Ipv6RibEdmAdvertItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6RibEdmAdvertItem) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetExtendedCommunities() []byte {
	if m != nil {
		return m.ExtendedCommunities
	}
	return nil
}

func (m *Ipv6RibEdmAdvertItem) GetProtocolOpaqueFlags() uint32 {
	if m != nil {
		return m.ProtocolOpaqueFlags
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetProtocolOpaque() uint32 {
	if m != nil {
		return m.ProtocolOpaque
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetInstanceName() []byte {
	if m != nil {
		return m.InstanceName
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6RibEdmAdvert_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert_KEYS")
	proto.RegisterType((*Ipv6RibEdmAdvert)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert")
	proto.RegisterType((*Ipv6RibEdmAdvertItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert_item")
}

func init() { proto.RegisterFile("ipv6_rib_edm_advert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x95, 0x6e, 0xb4, 0x9d, 0xd7, 0x8c, 0xe1, 0x82, 0x48, 0x35, 0x09, 0xaa, 0x71, 0x20,
	0x27, 0x4b, 0x14, 0x89, 0x3b, 0x9a, 0x8a, 0x34, 0x81, 0x98, 0x64, 0xb8, 0x70, 0xb2, 0x9c, 0xf8,
	0xef, 0x61, 0x29, 0x89, 0x83, 0xed, 0x46, 0xbd, 0xf1, 0x1c, 0x3c, 0x12, 0x57, 0xde, 0x80, 0x0b,
	0xcf, 0x81, 0x6c, 0x27, 0x20, 0xaa, 0x70, 0xde, 0xa5, 0x6e, 0xbe, 0xdf, 0xf7, 0x39, 0xff, 0xe8,
	0xfb, 0xa3, 0x95, 0x6a, 0xbb, 0x57, 0xcc, 0xa8, 0x82, 0x81, 0xa8, 0x19, 0x17, 0x1d, 0x18, 0x47,
	0x5a, 0xa3, 0x9d, 0xc6, 0x5f, 0x4b, 0x65, 0x4b, 0xcd, 0x94, 0xb6, 0x6c, 0x6f, 0x98, 0x6a, 0x83,
	0x2b, 0xd8, 0x75, 0x0b, 0x86, 0x0c, 0x41, 0xd2, 0x19, 0x69, 0xfd, 0x0f, 0xe1, 0xd2, 0x12, 0x2e,
	0x89, 0xf5, 0xa7, 0xe5, 0x92, 0xf4, 0x11, 0xa3, 0x77, 0x0e, 0x98, 0xe3, 0x45, 0x05, 0xac, 0xe1,
	0x35, 0xd8, 0xff, 0x01, 0x12, 0x5f, 0x6f, 0xfb, 0xf3, 0xf2, 0x47, 0x82, 0xb2, 0x91, 0xf1, 0xd8,
	0xdb, 0xed, 0xa7, 0x0f, 0x78, 0x85, 0xe6, 0x9d, 0x91, 0x21, 0x97, 0x25, 0xeb, 0x24, 0x3f, 0xa1,
	0xb3, 0xce, 0xc8, 0xf7, 0xbc, 0x06, 0xfc, 0x18, 0xcd, 0x78, 0x4f, 0x26, 0x81, 0x4c, 0x79, 0x04,
	0x2b, 0x34, 0xb7, 0x03, 0x39, 0x8a, 0x19, 0xdb, 0xa3, 0x1c, 0x9d, 0x1f, 0x8e, 0x93, 0x1d, 0x07,
	0xcb, 0x59, 0xd0, 0x3f, 0x7a, 0x39, 0x38, 0x33, 0x34, 0xe3, 0x42, 0x18, 0xb0, 0x36, 0xbb, 0x17,
	0xef, 0xe8, 0x1f, 0xf1, 0x33, 0x94, 0xb6, 0x06, 0xa4, 0xda, 0xb3, 0x0a, 0x9a, 0x5b, 0xf7, 0x39,
	0x9b, 0xae, 0x93, 0x3c, 0xa5, 0x8b, 0x28, 0xbe, 0x0b, 0xda, 0xe5, 0xcf, 0x04, 0x2d, 0x47, 0x3e,
	0x0a, 0x7f, 0x1f, 0xd7, 0xb3, 0xcd, 0xfa, 0x28, 0x3f, 0xdd, 0x7c, 0x4b, 0xc8, 0x1d, 0xb7, 0x41,
	0xc6, 0x9a, 0x50, 0x0e, 0x6a, 0x7a, 0xee, 0x09, 0x55, 0xc5, 0x56, 0xd4, 0xaf, 0x63, 0x71, 0xbf,
	0x26, 0xe3, 0xc5, 0x79, 0x3b, 0x7e, 0x8a, 0x4e, 0xc3, 0x7e, 0x95, 0xba, 0x62, 0x4a, 0x84, 0xee,
	0x52, 0x8a, 0x06, 0xe9, 0x5a, 0xe0, 0x0b, 0x74, 0x52, 0x56, 0x0a, 0x1a, 0xe7, 0xf1, 0x24, 0xe0,
	0x79, 0x14, 0xae, 0x05, 0xbe, 0x42, 0x4f, 0x9a, 0x5d, 0x5d, 0x80, 0x61, 0x5a, 0x32, 0xd8, 0x3b,
	0x68, 0x04, 0x08, 0x56, 0xea, 0xba, 0xde, 0x35, 0xca, 0x29, 0xb0, 0xa1, 0xd8, 0x94, 0x5e, 0x44,
	0xd7, 0x8d, 0xdc, 0xf6, 0x9e, 0xab, 0xbf, 0x16, 0xfc, 0x02, 0x3d, 0x1c, 0x8d, 0xfa, 0xc2, 0x17,
	0x74, 0x09, 0x23, 0x91, 0x0d, 0x7a, 0xf4, 0x67, 0x6a, 0xdd, 0xf2, 0x2f, 0x3b, 0x60, 0xb2, 0xe2,
	0xb7, 0x71, 0x07, 0x52, 0xba, 0x1c, 0xe0, 0x4d, 0x60, 0x6f, 0x3c, 0xc2, 0xcf, 0xd1, 0xfd, 0x83,
	0x4c, 0xbf, 0x11, 0x67, 0xff, 0xba, 0x31, 0x46, 0xc7, 0xa5, 0x16, 0x90, 0xcd, 0xd6, 0x49, 0xfe,
	0x80, 0x86, 0xff, 0x7e, 0x99, 0x54, 0x63, 0x1d, 0x6f, 0xca, 0x7e, 0x1b, 0xe7, 0x61, 0xb8, 0xc5,
	0x20, 0xfa, 0x5d, 0x2c, 0xa6, 0xe1, 0xa2, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x9b,
	0x25, 0x6b, 0xc6, 0x03, 0x00, 0x00,
}
