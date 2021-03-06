// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/app_placeholder_field.proto

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

// Possible values for App placeholder fields.
type AppPlaceholderFieldEnum_AppPlaceholderField int32

const (
	// Not specified.
	AppPlaceholderFieldEnum_UNSPECIFIED AppPlaceholderFieldEnum_AppPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	AppPlaceholderFieldEnum_UNKNOWN AppPlaceholderFieldEnum_AppPlaceholderField = 1
	// Data Type: INT64. The application store that the target application
	// belongs to. Valid values are: 1 = Apple iTunes Store; 2 = Google Play
	// Store.
	AppPlaceholderFieldEnum_STORE AppPlaceholderFieldEnum_AppPlaceholderField = 2
	// Data Type: STRING. The store-specific ID for the target application.
	AppPlaceholderFieldEnum_ID AppPlaceholderFieldEnum_AppPlaceholderField = 3
	// Data Type: STRING. The visible text displayed when the link is rendered
	// in an ad.
	AppPlaceholderFieldEnum_LINK_TEXT AppPlaceholderFieldEnum_AppPlaceholderField = 4
	// Data Type: STRING. The destination URL of the in-app link.
	AppPlaceholderFieldEnum_URL AppPlaceholderFieldEnum_AppPlaceholderField = 5
	// Data Type: URL_LIST. Final URLs for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_FINAL_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 6
	// Data Type: URL_LIST. Final Mobile URLs for the in-app link when using
	// Upgraded URLs.
	AppPlaceholderFieldEnum_FINAL_MOBILE_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 7
	// Data Type: URL. Tracking template for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_TRACKING_URL AppPlaceholderFieldEnum_AppPlaceholderField = 8
	// Data Type: STRING. Final URL suffix for the in-app link when using
	// parallel tracking.
	AppPlaceholderFieldEnum_FINAL_URL_SUFFIX AppPlaceholderFieldEnum_AppPlaceholderField = 9
)

var AppPlaceholderFieldEnum_AppPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STORE",
	3: "ID",
	4: "LINK_TEXT",
	5: "URL",
	6: "FINAL_URLS",
	7: "FINAL_MOBILE_URLS",
	8: "TRACKING_URL",
	9: "FINAL_URL_SUFFIX",
}
var AppPlaceholderFieldEnum_AppPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"STORE":             2,
	"ID":                3,
	"LINK_TEXT":         4,
	"URL":               5,
	"FINAL_URLS":        6,
	"FINAL_MOBILE_URLS": 7,
	"TRACKING_URL":      8,
	"FINAL_URL_SUFFIX":  9,
}

func (x AppPlaceholderFieldEnum_AppPlaceholderField) String() string {
	return proto.EnumName(AppPlaceholderFieldEnum_AppPlaceholderField_name, int32(x))
}
func (AppPlaceholderFieldEnum_AppPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_app_placeholder_field_f3f678b0f2e2b317, []int{0, 0}
}

// Values for App placeholder fields.
type AppPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPlaceholderFieldEnum) Reset()         { *m = AppPlaceholderFieldEnum{} }
func (m *AppPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*AppPlaceholderFieldEnum) ProtoMessage()    {}
func (*AppPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_placeholder_field_f3f678b0f2e2b317, []int{0}
}
func (m *AppPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *AppPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *AppPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPlaceholderFieldEnum.Merge(dst, src)
}
func (m *AppPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Size(m)
}
func (m *AppPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.AppPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AppPlaceholderFieldEnum_AppPlaceholderField", AppPlaceholderFieldEnum_AppPlaceholderField_name, AppPlaceholderFieldEnum_AppPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/app_placeholder_field.proto", fileDescriptor_app_placeholder_field_f3f678b0f2e2b317)
}

var fileDescriptor_app_placeholder_field_f3f678b0f2e2b317 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xea, 0x40,
	0x18, 0xc5, 0x6f, 0xcb, 0x05, 0x2e, 0x1f, 0xf7, 0x5e, 0xc7, 0x51, 0xa3, 0x1b, 0x16, 0xf2, 0x00,
	0xd3, 0x26, 0xae, 0x8c, 0xab, 0x16, 0x5a, 0x32, 0xa1, 0x96, 0xa6, 0x7f, 0x90, 0x98, 0x26, 0x4d,
	0xa5, 0xb5, 0x92, 0x94, 0xce, 0x84, 0x11, 0x1e, 0xc8, 0xa5, 0x89, 0x0f, 0x82, 0x4f, 0x65, 0xda,
	0x0a, 0x6e, 0xd0, 0xcd, 0xe4, 0xcc, 0x77, 0xe6, 0x37, 0x93, 0x73, 0x06, 0xae, 0x33, 0xc6, 0xb2,
	0x3c, 0x55, 0xe2, 0x44, 0x28, 0xb5, 0x2c, 0xd5, 0x46, 0x55, 0xd2, 0x62, 0xbd, 0x14, 0x4a, 0xcc,
	0x79, 0xc4, 0xf3, 0x78, 0x9e, 0x3e, 0xb1, 0x3c, 0x49, 0x57, 0xd1, 0xe3, 0x22, 0xcd, 0x13, 0xc2,
	0x57, 0xec, 0x99, 0xe1, 0x5e, 0x7d, 0x9e, 0xc4, 0x89, 0x20, 0x7b, 0x94, 0x6c, 0x54, 0x52, 0xa1,
	0xfd, 0xad, 0x04, 0xe7, 0x1a, 0xe7, 0xce, 0x17, 0x6d, 0x96, 0xb0, 0x51, 0xac, 0x97, 0xfd, 0x37,
	0x09, 0x4e, 0x0e, 0x78, 0xf8, 0x08, 0xba, 0x81, 0xed, 0x39, 0xc6, 0x80, 0x9a, 0xd4, 0x18, 0xa2,
	0x5f, 0xb8, 0x0b, 0xed, 0xc0, 0x1e, 0xdb, 0x93, 0x3b, 0x1b, 0x49, 0xb8, 0x03, 0x4d, 0xcf, 0x9f,
	0xb8, 0x06, 0x92, 0x71, 0x0b, 0x64, 0x3a, 0x44, 0x0d, 0xfc, 0x0f, 0x3a, 0x16, 0xb5, 0xc7, 0x91,
	0x6f, 0xcc, 0x7c, 0xf4, 0x1b, 0xb7, 0xa1, 0x11, 0xb8, 0x16, 0x6a, 0xe2, 0xff, 0x00, 0x26, 0xb5,
	0x35, 0x2b, 0x0a, 0x5c, 0xcb, 0x43, 0x2d, 0x7c, 0x06, 0xc7, 0xf5, 0xfe, 0x76, 0xa2, 0x53, 0xcb,
	0xa8, 0xc7, 0x6d, 0x8c, 0xe0, 0xaf, 0xef, 0x6a, 0x83, 0x31, 0xb5, 0x47, 0xe5, 0x08, 0xfd, 0xc1,
	0xa7, 0x80, 0xf6, 0x60, 0xe4, 0x05, 0xa6, 0x49, 0x67, 0xa8, 0xa3, 0x6f, 0x25, 0xb8, 0x9c, 0xb3,
	0x25, 0xf9, 0x31, 0xb1, 0x7e, 0x71, 0x20, 0x92, 0x53, 0x56, 0xe5, 0x48, 0xf7, 0xfa, 0x27, 0x9a,
	0xb1, 0x3c, 0x2e, 0x32, 0xc2, 0x56, 0x99, 0x92, 0xa5, 0x45, 0x55, 0xe4, 0xae, 0x77, 0xbe, 0x10,
	0xdf, 0x7c, 0xc3, 0x4d, 0xb5, 0xbe, 0xc8, 0x8d, 0x91, 0xa6, 0xbd, 0xca, 0xbd, 0x51, 0x7d, 0x95,
	0x96, 0x08, 0x52, 0xcb, 0x52, 0x4d, 0x55, 0x52, 0x56, 0x2b, 0xde, 0x77, 0x7e, 0xa8, 0x25, 0x22,
	0xdc, 0xfb, 0xe1, 0x54, 0x0d, 0x2b, 0xff, 0xa1, 0x55, 0x3d, 0x7a, 0xf5, 0x11, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x9f, 0x59, 0xbc, 0xfa, 0x01, 0x00, 0x00,
}
