// Code generated by protoc-gen-go.
// source: testupdatefoo.proto
// DO NOT EDIT!

package testagg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestAggFooUpdated struct {
	AggregateId string `protobuf:"bytes,1,opt,name=aggregateId" json:"aggregateId,omitempty"`
	NewFoo      string `protobuf:"bytes,2,opt,name=newFoo" json:"newFoo,omitempty"`
}

func (m *TestAggFooUpdated) Reset()                    { *m = TestAggFooUpdated{} }
func (m *TestAggFooUpdated) String() string            { return proto.CompactTextString(m) }
func (*TestAggFooUpdated) ProtoMessage()               {}
func (*TestAggFooUpdated) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*TestAggFooUpdated)(nil), "eventstore.TestAggFooUpdated")
}

var fileDescriptor1 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x49, 0x2d, 0x2e,
	0x29, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x4d, 0xcb, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4a, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x2e, 0xc9, 0x2f, 0x4a, 0x55, 0xf2, 0xe5, 0x12, 0x0c,
	0x01, 0x2a, 0x71, 0x4c, 0x4f, 0x77, 0xcb, 0xcf, 0x0f, 0x05, 0x2b, 0x4c, 0x11, 0x52, 0xe0, 0xe2,
	0x4e, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x07, 0xf2, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x90, 0x85, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0x52, 0xcb, 0x81, 0x5a, 0x24, 0x98, 0xc0, 0x92,
	0x50, 0x5e, 0x12, 0x1b, 0xd8, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x1e, 0xfe,
	0x25, 0x78, 0x00, 0x00, 0x00,
}
