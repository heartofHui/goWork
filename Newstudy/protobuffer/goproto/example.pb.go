// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Test
*/
package example

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

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64 `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x4c, 0xb9,
	0x58, 0x42, 0x52, 0x8b, 0x4b, 0x84, 0x78, 0xb9, 0x58, 0x73, 0x12, 0x93, 0x52, 0x73, 0x24, 0x18,
	0x15, 0x98, 0x34, 0x38, 0x85, 0x04, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x58, 0xad, 0x98, 0xcc, 0xcd, 0x85, 0x78, 0xb8, 0x58, 0x8a, 0x52, 0x0b, 0x8a, 0x25, 0x98,
	0x15, 0x98, 0x35, 0x98, 0xb5, 0x78, 0xb8, 0x98, 0xdd, 0xfc, 0xfd, 0x85, 0x58, 0xb9, 0x18, 0x23,
	0x04, 0x04, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xa8, 0xc1, 0x6b, 0x5d, 0x00, 0x00, 0x00,
}
