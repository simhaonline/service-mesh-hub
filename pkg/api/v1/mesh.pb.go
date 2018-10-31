// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Any user-configurable settings for a service mesh
type MeshConfig struct {
	Identities           []*Identity `protobuf:"bytes,1,rep,name=identities" json:"identities,omitempty"`
	Policies             []*Policy   `protobuf:"bytes,2,rep,name=policies" json:"policies,omitempty"`
	Routes               []*Route    `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MeshConfig) Reset()         { *m = MeshConfig{} }
func (m *MeshConfig) String() string { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()    {}
func (*MeshConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_42e58bd64c07b4fa, []int{0}
}
func (m *MeshConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshConfig.Unmarshal(m, b)
}
func (m *MeshConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshConfig.Marshal(b, m, deterministic)
}
func (dst *MeshConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshConfig.Merge(dst, src)
}
func (m *MeshConfig) XXX_Size() int {
	return xxx_messageInfo_MeshConfig.Size(m)
}
func (m *MeshConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MeshConfig proto.InternalMessageInfo

func (m *MeshConfig) GetIdentities() []*Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

func (m *MeshConfig) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *MeshConfig) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Identity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_42e58bd64c07b4fa, []int{1}
}
func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (dst *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(dst, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Policy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_42e58bd64c07b4fa, []int{2}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

type Route struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_42e58bd64c07b4fa, []int{3}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MeshConfig)(nil), "supergloo.solo.io.MeshConfig")
	proto.RegisterType((*Identity)(nil), "supergloo.solo.io.Identity")
	proto.RegisterType((*Policy)(nil), "supergloo.solo.io.Policy")
	proto.RegisterType((*Route)(nil), "supergloo.solo.io.Route")
}
func (this *MeshConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshConfig)
	if !ok {
		that2, ok := that.(MeshConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Identities) != len(that1.Identities) {
		return false
	}
	for i := range this.Identities {
		if !this.Identities[i].Equal(that1.Identities[i]) {
			return false
		}
	}
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
			return false
		}
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Identity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Identity)
	if !ok {
		that2, ok := that.(Identity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("mesh.proto", fileDescriptor_mesh_42e58bd64c07b4fa) }

var fileDescriptor_mesh_42e58bd64c07b4fa = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0x04, 0x31,
	0x14, 0x85, 0x89, 0xab, 0xe3, 0x78, 0xad, 0x0c, 0x16, 0x51, 0x61, 0x59, 0xd2, 0xb8, 0xcd, 0x26,
	0xfe, 0x60, 0x65, 0xa7, 0x95, 0x85, 0x20, 0x29, 0xed, 0x76, 0xd7, 0x98, 0xb9, 0x38, 0x33, 0x37,
	0x4c, 0x32, 0xc2, 0xbe, 0x91, 0xb5, 0x8f, 0xe4, 0x93, 0x48, 0x32, 0xba, 0x08, 0x4e, 0x77, 0x20,
	0xdf, 0x77, 0x4e, 0xb8, 0x00, 0x8d, 0x0d, 0x95, 0xf2, 0x1d, 0x45, 0xe2, 0x47, 0xa1, 0xf7, 0xb6,
	0x73, 0x35, 0x91, 0x0a, 0x54, 0x93, 0x42, 0x3a, 0x3d, 0x76, 0xe4, 0x28, 0xbf, 0xea, 0x94, 0x06,
	0x50, 0x7e, 0x32, 0x80, 0x47, 0x1b, 0xaa, 0x7b, 0x6a, 0x5f, 0xd1, 0xf1, 0x5b, 0x00, 0x7c, 0xb1,
	0x6d, 0xc4, 0x88, 0x36, 0x08, 0x36, 0x9b, 0xcc, 0x0f, 0xaf, 0xce, 0xd4, 0xbf, 0x32, 0xf5, 0x30,
	0x40, 0x1b, 0xf3, 0x07, 0xe7, 0x37, 0x50, 0x7a, 0xaa, 0x71, 0x9d, 0xd4, 0x9d, 0xac, 0x9e, 0x8c,
	0xa8, 0x4f, 0x09, 0xd9, 0x98, 0x2d, 0xca, 0x2f, 0xa0, 0xe8, 0xa8, 0x8f, 0x36, 0x88, 0x49, 0x96,
	0xc4, 0x88, 0x64, 0x12, 0x60, 0x7e, 0x38, 0x39, 0x85, 0xf2, 0xf7, 0x03, 0x9c, 0xc3, 0x6e, 0xbb,
	0x6c, 0xac, 0x60, 0x33, 0x36, 0x3f, 0x30, 0x39, 0xcb, 0x12, 0x8a, 0x61, 0x45, 0xee, 0xc3, 0x5e,
	0x56, 0xef, 0x16, 0x1f, 0x5f, 0x53, 0xf6, 0x7c, 0xee, 0x30, 0x56, 0xfd, 0x4a, 0xad, 0xa9, 0xd1,
	0xa9, 0x7a, 0x81, 0xa4, 0xb7, 0x63, 0xda, 0xbf, 0x39, 0xbd, 0xf4, 0xa8, 0xdf, 0x2f, 0x57, 0x45,
	0xbe, 0xce, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xda, 0x59, 0x09, 0x54, 0x01, 0x00,
	0x00,
}