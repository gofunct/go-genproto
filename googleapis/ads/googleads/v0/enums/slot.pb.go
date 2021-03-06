// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/slot.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enumerates possible positions of the Ad.
type SlotEnum_Slot int32

const (
	// Not specified.
	SlotEnum_UNSPECIFIED SlotEnum_Slot = 0
	// The value is unknown in this version.
	SlotEnum_UNKNOWN SlotEnum_Slot = 1
	// Google search: Side.
	SlotEnum_SEARCH_SIDE SlotEnum_Slot = 2
	// Google search: Top.
	SlotEnum_SEARCH_TOP SlotEnum_Slot = 3
	// Google search: Other.
	SlotEnum_SEARCH_OTHER SlotEnum_Slot = 4
	// Google Display Network.
	SlotEnum_CONTENT SlotEnum_Slot = 5
	// Search partners: Top.
	SlotEnum_SEARCH_PARTNER_TOP SlotEnum_Slot = 6
	// Search partners: Other.
	SlotEnum_SEARCH_PARTNER_OTHER SlotEnum_Slot = 7
	// Cross-network.
	SlotEnum_MIXED SlotEnum_Slot = 8
)

var SlotEnum_Slot_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH_SIDE",
	3: "SEARCH_TOP",
	4: "SEARCH_OTHER",
	5: "CONTENT",
	6: "SEARCH_PARTNER_TOP",
	7: "SEARCH_PARTNER_OTHER",
	8: "MIXED",
}
var SlotEnum_Slot_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"SEARCH_SIDE":          2,
	"SEARCH_TOP":           3,
	"SEARCH_OTHER":         4,
	"CONTENT":              5,
	"SEARCH_PARTNER_TOP":   6,
	"SEARCH_PARTNER_OTHER": 7,
	"MIXED":                8,
}

func (x SlotEnum_Slot) String() string {
	return proto.EnumName(SlotEnum_Slot_name, int32(x))
}
func (SlotEnum_Slot) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_slot_595ff7b86bb21b12, []int{0, 0}
}

// Container for enumeration of possible positions of the Ad.
type SlotEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlotEnum) Reset()         { *m = SlotEnum{} }
func (m *SlotEnum) String() string { return proto.CompactTextString(m) }
func (*SlotEnum) ProtoMessage()    {}
func (*SlotEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_slot_595ff7b86bb21b12, []int{0}
}
func (m *SlotEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotEnum.Unmarshal(m, b)
}
func (m *SlotEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotEnum.Marshal(b, m, deterministic)
}
func (dst *SlotEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotEnum.Merge(dst, src)
}
func (m *SlotEnum) XXX_Size() int {
	return xxx_messageInfo_SlotEnum.Size(m)
}
func (m *SlotEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SlotEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SlotEnum)(nil), "google.ads.googleads.v0.enums.SlotEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.SlotEnum_Slot", SlotEnum_Slot_name, SlotEnum_Slot_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/slot.proto", fileDescriptor_slot_595ff7b86bb21b12)
}

var fileDescriptor_slot_595ff7b86bb21b12 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xf4, 0x7f, 0x6f, 0x3f, 0x74, 0x18, 0x44, 0xdc, 0x74, 0x61, 0x57, 0xae, 0x26,
	0x01, 0x97, 0xae, 0xa6, 0xed, 0xd8, 0x06, 0x71, 0x12, 0x92, 0xb4, 0x8a, 0x04, 0xa4, 0x9a, 0x30,
	0x08, 0x49, 0xa6, 0x74, 0xd2, 0xbe, 0x8f, 0x2e, 0x5d, 0xf8, 0x00, 0x3e, 0x82, 0x4f, 0x25, 0x33,
	0x49, 0xbb, 0x10, 0x74, 0x13, 0xce, 0xbd, 0xbf, 0x73, 0x32, 0xdc, 0x03, 0x17, 0x42, 0x4a, 0x91,
	0xa5, 0xf6, 0x2a, 0x51, 0x76, 0x25, 0xb5, 0xda, 0x39, 0x76, 0x5a, 0x6c, 0x73, 0x65, 0xab, 0x4c,
	0x96, 0x64, 0xbd, 0x91, 0xa5, 0xc4, 0xc3, 0x0a, 0x93, 0x55, 0xa2, 0xc8, 0xc1, 0x49, 0x76, 0x0e,
	0x31, 0xce, 0xd1, 0x87, 0x05, 0xbd, 0x30, 0x93, 0x25, 0x2b, 0xb6, 0xf9, 0xe8, 0xd5, 0x82, 0x96,
	0x1e, 0xf0, 0x31, 0x0c, 0x16, 0x3c, 0xf4, 0xd9, 0xc4, 0xbd, 0x76, 0xd9, 0x14, 0xfd, 0xc3, 0x03,
	0xe8, 0x2e, 0xf8, 0x0d, 0xf7, 0xee, 0x38, 0xb2, 0x34, 0x0d, 0x19, 0x0d, 0x26, 0xf3, 0xc7, 0xd0,
	0x9d, 0x32, 0xd4, 0xc0, 0x47, 0x00, 0xf5, 0x22, 0xf2, 0x7c, 0xd4, 0xc4, 0x08, 0xfe, 0xd7, 0xb3,
	0x17, 0xcd, 0x59, 0x80, 0x5a, 0x3a, 0x3f, 0xf1, 0x78, 0xc4, 0x78, 0x84, 0xda, 0xf8, 0x14, 0x70,
	0x8d, 0x7d, 0x1a, 0x44, 0x9c, 0x05, 0x26, 0xd6, 0xc1, 0x67, 0x70, 0xf2, 0x63, 0x5f, 0xc5, 0xbb,
	0xb8, 0x0f, 0xed, 0x5b, 0xf7, 0x9e, 0x4d, 0x51, 0x6f, 0xfc, 0x69, 0xc1, 0xf9, 0xb3, 0xcc, 0xc9,
	0x9f, 0x67, 0x8d, 0xfb, 0xfa, 0x0c, 0x5f, 0x17, 0xe0, 0x5b, 0x0f, 0xe3, 0xda, 0x2b, 0x64, 0xb6,
	0x2a, 0x04, 0x91, 0x1b, 0x61, 0x8b, 0xb4, 0x30, 0xf5, 0xec, 0xcb, 0x5b, 0xbf, 0xa8, 0x5f, 0xba,
	0xbc, 0x32, 0xdf, 0xb7, 0x46, 0x73, 0x46, 0xe9, 0x7b, 0x63, 0x38, 0xab, 0x7e, 0x45, 0x13, 0x45,
	0x2a, 0xa9, 0xd5, 0xd2, 0x21, 0xba, 0x3f, 0xf5, 0xb5, 0xe7, 0x31, 0x4d, 0x54, 0x7c, 0xe0, 0xf1,
	0xd2, 0x89, 0x0d, 0x7f, 0xea, 0x98, 0x47, 0x2f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x68, 0x82,
	0x7a, 0x46, 0xbf, 0x01, 0x00, 0x00,
}
