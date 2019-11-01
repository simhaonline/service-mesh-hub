// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/cluster-route.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MeshBridge struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata   core.Metadata       `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	SourceMesh *ClusterResourceRef `protobuf:"bytes,3,opt,name=source_mesh,json=sourceMesh,proto3" json:"source_mesh,omitempty"`
	TargetMesh *ClusterResourceRef `protobuf:"bytes,4,opt,name=target_mesh,json=targetMesh,proto3" json:"target_mesh,omitempty"`
	// Namespace in which to find gloo
	GlooNamespace        string              `protobuf:"bytes,5,opt,name=gloo_namespace,json=glooNamespace,proto3" json:"gloo_namespace,omitempty"`
	Sources              []*core.ResourceRef `protobuf:"bytes,6,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshBridge) Reset()         { *m = MeshBridge{} }
func (m *MeshBridge) String() string { return proto.CompactTextString(m) }
func (*MeshBridge) ProtoMessage()    {}
func (*MeshBridge) Descriptor() ([]byte, []int) {
	return fileDescriptor_16671c1c9aa51037, []int{0}
}
func (m *MeshBridge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshBridge.Unmarshal(m, b)
}
func (m *MeshBridge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshBridge.Marshal(b, m, deterministic)
}
func (m *MeshBridge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshBridge.Merge(m, src)
}
func (m *MeshBridge) XXX_Size() int {
	return xxx_messageInfo_MeshBridge.Size(m)
}
func (m *MeshBridge) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshBridge.DiscardUnknown(m)
}

var xxx_messageInfo_MeshBridge proto.InternalMessageInfo

func (m *MeshBridge) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MeshBridge) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MeshBridge) GetSourceMesh() *ClusterResourceRef {
	if m != nil {
		return m.SourceMesh
	}
	return nil
}

func (m *MeshBridge) GetTargetMesh() *ClusterResourceRef {
	if m != nil {
		return m.TargetMesh
	}
	return nil
}

func (m *MeshBridge) GetGlooNamespace() string {
	if m != nil {
		return m.GlooNamespace
	}
	return ""
}

func (m *MeshBridge) GetSources() []*core.ResourceRef {
	if m != nil {
		return m.Sources
	}
	return nil
}

type ClusterResourceRef struct {
	Resource             *core.ResourceRef `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	ClusterName          string            `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceRef) Reset()         { *m = ClusterResourceRef{} }
func (m *ClusterResourceRef) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceRef) ProtoMessage()    {}
func (*ClusterResourceRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_16671c1c9aa51037, []int{1}
}
func (m *ClusterResourceRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceRef.Unmarshal(m, b)
}
func (m *ClusterResourceRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceRef.Marshal(b, m, deterministic)
}
func (m *ClusterResourceRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceRef.Merge(m, src)
}
func (m *ClusterResourceRef) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceRef.Size(m)
}
func (m *ClusterResourceRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceRef.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceRef proto.InternalMessageInfo

func (m *ClusterResourceRef) GetResource() *core.ResourceRef {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ClusterResourceRef) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func init() {
	proto.RegisterType((*MeshBridge)(nil), "zephyr.solo.io.MeshBridge")
	proto.RegisterType((*ClusterResourceRef)(nil), "zephyr.solo.io.ClusterResourceRef")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/v1/cluster-route.proto", fileDescriptor_16671c1c9aa51037)
}

var fileDescriptor_16671c1c9aa51037 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x8d, 0x09, 0xed, 0xa6, 0x44, 0x62, 0x55, 0x21, 0x37, 0x07, 0x1a, 0x2c, 0x21, 0x45,
	0x82, 0xac, 0x95, 0x56, 0x48, 0xa8, 0xdc, 0xd2, 0x03, 0xa7, 0x72, 0x30, 0x37, 0x2e, 0xd1, 0xda,
	0x19, 0xaf, 0x97, 0xc6, 0x19, 0x6b, 0x77, 0x8d, 0x04, 0xc7, 0xbe, 0x01, 0x6f, 0xc1, 0xa3, 0xf0,
	0x14, 0x3d, 0xf0, 0x06, 0xe5, 0x09, 0xd0, 0xfe, 0xb8, 0x22, 0x02, 0x41, 0x38, 0x79, 0x77, 0xe6,
	0xfb, 0xbe, 0x99, 0xf1, 0x37, 0x4b, 0x5e, 0x0b, 0x69, 0xaa, 0x36, 0x67, 0x05, 0xd6, 0xa9, 0xc6,
	0x35, 0xce, 0x24, 0xa6, 0x35, 0xe8, 0x6a, 0xd6, 0x28, 0xfc, 0x00, 0x85, 0xd1, 0x29, 0x6f, 0x64,
	0xfa, 0x71, 0x9e, 0x16, 0xeb, 0x56, 0x1b, 0x50, 0x33, 0x85, 0xad, 0x01, 0xd6, 0x28, 0x34, 0x48,
	0x47, 0x9f, 0xa1, 0xa9, 0x3e, 0x29, 0x66, 0x89, 0x4c, 0xe2, 0xf8, 0x48, 0xa0, 0x40, 0x97, 0x4a,
	0xed, 0xc9, 0xa3, 0xc6, 0xf3, 0x3f, 0x94, 0x70, 0xdf, 0x2b, 0x69, 0x3a, 0xf5, 0xee, 0x1e, 0x28,
	0x2f, 0x76, 0xa0, 0x28, 0x28, 0x03, 0x3a, 0xdd, 0xa5, 0x80, 0xe1, 0xa6, 0xd5, 0xff, 0xd1, 0x51,
	0x0d, 0x86, 0xaf, 0xb8, 0xe1, 0x9e, 0x92, 0x7c, 0xe9, 0x13, 0x72, 0x09, 0xba, 0x5a, 0x28, 0xb9,
	0x12, 0x40, 0xdf, 0x90, 0x81, 0x57, 0x8c, 0x7b, 0x93, 0xde, 0x74, 0x78, 0x7a, 0xc4, 0x0a, 0x54,
	0xd0, 0xfd, 0x08, 0xf6, 0xce, 0xe5, 0x16, 0xc7, 0xdf, 0x6e, 0x4e, 0xee, 0xfd, 0xb8, 0x39, 0x79,
	0x64, 0x40, 0x9b, 0x95, 0x2c, 0xcb, 0xf3, 0x44, 0x8a, 0x0d, 0x2a, 0x48, 0xb2, 0x40, 0xa7, 0xaf,
	0xc8, 0x7e, 0x57, 0x29, 0xde, 0x73, 0x52, 0x8f, 0xb7, 0xa5, 0x2e, 0x43, 0x76, 0x11, 0x59, 0xb1,
	0xec, 0x0e, 0x4d, 0x2f, 0xc8, 0x50, 0x63, 0xab, 0x0a, 0x58, 0x5a, 0xbf, 0xe2, 0xbe, 0x23, 0x27,
	0x6c, 0xdb, 0x12, 0x76, 0xe1, 0x6d, 0xcb, 0xc0, 0x63, 0x33, 0x28, 0x33, 0xe2, 0x8f, 0x76, 0x1a,
	0x2b, 0x62, 0xb8, 0x12, 0x60, 0xbc, 0x48, 0xb4, 0xbb, 0x88, 0xa7, 0x39, 0x91, 0x67, 0x64, 0x24,
	0xd6, 0x88, 0xcb, 0x0d, 0xaf, 0x41, 0x37, 0xbc, 0x80, 0xf8, 0xfe, 0xa4, 0x37, 0x3d, 0xc8, 0x1e,
	0xda, 0xe8, 0xdb, 0x2e, 0x48, 0xcf, 0xc8, 0x03, 0xcf, 0xd7, 0xf1, 0x60, 0xd2, 0x9f, 0x0e, 0x4f,
	0x8f, 0xb7, 0x27, 0xfd, 0x55, 0xbe, 0x43, 0x9e, 0x8f, 0xaf, 0x6f, 0xa3, 0x88, 0xec, 0xd5, 0xf9,
	0xf5, 0x6d, 0x34, 0xa2, 0x87, 0xb6, 0xcb, 0x65, 0xee, 0x3c, 0xd0, 0xc9, 0x86, 0xd0, 0xdf, 0x3b,
	0xa3, 0x2f, 0xc9, 0xbe, 0x0a, 0xd7, 0x60, 0xce, 0x5f, 0xea, 0xdc, 0x41, 0xe9, 0x53, 0x72, 0x18,
	0x56, 0xdc, 0xcd, 0xe1, 0xcc, 0x38, 0xc8, 0x86, 0x21, 0x66, 0xa7, 0x58, 0xcc, 0xbf, 0x7e, 0x7f,
	0xd2, 0x7b, 0xff, 0xfc, 0x9f, 0x2f, 0xa6, 0xb9, 0x12, 0x61, 0x8b, 0xf2, 0x81, 0xdb, 0x9e, 0xb3,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0xa7, 0x18, 0x6d, 0x67, 0x03, 0x00, 0x00,
}

func (this *MeshBridge) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshBridge)
	if !ok {
		that2, ok := that.(MeshBridge)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.SourceMesh.Equal(that1.SourceMesh) {
		return false
	}
	if !this.TargetMesh.Equal(that1.TargetMesh) {
		return false
	}
	if this.GlooNamespace != that1.GlooNamespace {
		return false
	}
	if len(this.Sources) != len(that1.Sources) {
		return false
	}
	for i := range this.Sources {
		if !this.Sources[i].Equal(that1.Sources[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClusterResourceRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterResourceRef)
	if !ok {
		that2, ok := that.(ClusterResourceRef)
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
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
