// Code generated by protoc-gen-go.
// source: hapi/chart/metadata.proto
// DO NOT EDIT!

package chart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// Maintainer:
//
// 		A descriptor of the Chart maintainer(s).
//
type Maintainer struct {
	// Name is a user name or organization name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Maintainer) Reset()                    { *m = Maintainer{} }
func (m *Maintainer) String() string            { return proto.CompactTextString(m) }
func (*Maintainer) ProtoMessage()               {}
func (*Maintainer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

//
// Metadata:
//
// 		Metadata for a Chart file. This models the structure
// 		of a Chart.yaml file.
//
// 		Spec: https://github.com/kubernetes/helm/blob/master/docs/design/chart_format.md#the-chart-file
//
type Metadata struct {
	// The name of the chart
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The URL to a relecant project page, git repo, or contact person
	Home string `protobuf:"bytes,2,opt,name=home" json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `protobuf:"bytes,3,rep,name=sources" json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `protobuf:"bytes,6,rep,name=keywords" json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `protobuf:"bytes,7,rep,name=maintainers" json:"maintainers,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Metadata) GetMaintainers() []*Maintainer {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func init() {
	proto.RegisterType((*Maintainer)(nil), "hapi.chart.Maintainer")
	proto.RegisterType((*Metadata)(nil), "hapi.chart.Metadata")
}

var fileDescriptor2 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc4, 0x30,
	0x0c, 0xc5, 0x55, 0xee, 0x7a, 0x3d, 0xdc, 0xcd, 0x42, 0x28, 0x30, 0x55, 0x37, 0x31, 0xe5, 0x24,
	0x90, 0x10, 0x33, 0xfb, 0x2d, 0x37, 0xb2, 0x99, 0xd6, 0x52, 0x23, 0x48, 0x53, 0x25, 0x01, 0xc4,
	0x97, 0xe5, 0xb3, 0x90, 0xba, 0xf4, 0xcf, 0xc0, 0x60, 0xc9, 0xef, 0xfd, 0xfc, 0x2c, 0xd9, 0x70,
	0xd3, 0x52, 0x6f, 0x8e, 0x75, 0x4b, 0x3e, 0x1e, 0x2d, 0x47, 0x6a, 0x28, 0x92, 0xee, 0xbd, 0x8b,
	0x0e, 0x61, 0x40, 0x5a, 0xd0, 0xe1, 0x11, 0xe0, 0x44, 0xa6, 0x8b, 0xa9, 0xd8, 0x23, 0xc2, 0xb6,
	0x23, 0xcb, 0x2a, 0xab, 0xb2, 0xbb, 0xcb, 0xb3, 0xf4, 0x78, 0x05, 0x39, 0x5b, 0x32, 0xef, 0xea,
	0x42, 0xcc, 0x51, 0x1c, 0x7e, 0x32, 0xd8, 0x9f, 0xfe, 0xd6, 0xfe, 0x1b, 0x4b, 0x5e, 0xeb, 0x92,
	0x37, 0xa6, 0xa4, 0x47, 0x05, 0x45, 0x70, 0x1f, 0xbe, 0xe6, 0xa0, 0x36, 0xd5, 0x26, 0xd9, 0x93,
	0x1c, 0xc8, 0x27, 0xfb, 0x60, 0x5c, 0xa7, 0xb6, 0x12, 0x98, 0x24, 0x56, 0x50, 0x36, 0x1c, 0x6a,
	0x6f, 0xfa, 0x38, 0xd0, 0x5c, 0xe8, 0xda, 0xc2, 0x5b, 0xd8, 0xbf, 0xf1, 0xf7, 0x97, 0xf3, 0x4d,
	0x50, 0x3b, 0x59, 0x3b, 0x6b, 0x7c, 0x82, 0xd2, 0xce, 0xe7, 0x05, 0x55, 0x24, 0x5c, 0xde, 0x5f,
	0xeb, 0xe5, 0x01, 0x7a, 0xb9, 0xfe, 0xbc, 0x1e, 0x7d, 0x2e, 0x5e, 0x72, 0x19, 0x78, 0xdd, 0xc9,
	0xd3, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xaf, 0x44, 0xa7, 0x51, 0x01, 0x00, 0x00,
}