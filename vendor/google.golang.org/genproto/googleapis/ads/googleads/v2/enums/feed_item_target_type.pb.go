// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/feed_item_target_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Possible type of a feed item target.
type FeedItemTargetTypeEnum_FeedItemTargetType int32

const (
	// Not specified.
	FeedItemTargetTypeEnum_UNSPECIFIED FeedItemTargetTypeEnum_FeedItemTargetType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemTargetTypeEnum_UNKNOWN FeedItemTargetTypeEnum_FeedItemTargetType = 1
	// Feed item targets a campaign.
	FeedItemTargetTypeEnum_CAMPAIGN FeedItemTargetTypeEnum_FeedItemTargetType = 2
	// Feed item targets an ad group.
	FeedItemTargetTypeEnum_AD_GROUP FeedItemTargetTypeEnum_FeedItemTargetType = 3
	// Feed item targets a criterion.
	FeedItemTargetTypeEnum_CRITERION FeedItemTargetTypeEnum_FeedItemTargetType = 4
)

var FeedItemTargetTypeEnum_FeedItemTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CAMPAIGN",
	3: "AD_GROUP",
	4: "CRITERION",
}

var FeedItemTargetTypeEnum_FeedItemTargetType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CAMPAIGN":    2,
	"AD_GROUP":    3,
	"CRITERION":   4,
}

func (x FeedItemTargetTypeEnum_FeedItemTargetType) String() string {
	return proto.EnumName(FeedItemTargetTypeEnum_FeedItemTargetType_name, int32(x))
}

func (FeedItemTargetTypeEnum_FeedItemTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e4383fc8865747, []int{0, 0}
}

// Container for enum describing possible types of a feed item target.
type FeedItemTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemTargetTypeEnum) Reset()         { *m = FeedItemTargetTypeEnum{} }
func (m *FeedItemTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetTypeEnum) ProtoMessage()    {}
func (*FeedItemTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e4383fc8865747, []int{0}
}

func (m *FeedItemTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Unmarshal(m, b)
}
func (m *FeedItemTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetTypeEnum.Merge(m, src)
}
func (m *FeedItemTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Size(m)
}
func (m *FeedItemTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.FeedItemTargetTypeEnum_FeedItemTargetType", FeedItemTargetTypeEnum_FeedItemTargetType_name, FeedItemTargetTypeEnum_FeedItemTargetType_value)
	proto.RegisterType((*FeedItemTargetTypeEnum)(nil), "google.ads.googleads.v2.enums.FeedItemTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/feed_item_target_type.proto", fileDescriptor_51e4383fc8865747)
}

var fileDescriptor_51e4383fc8865747 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xfb, 0x30,
	0x1c, 0xfd, 0xaf, 0xfb, 0xe3, 0x47, 0xa6, 0x58, 0x7a, 0xa1, 0x20, 0xee, 0x62, 0x7b, 0x80, 0x14,
	0xea, 0x95, 0xf1, 0x2a, 0xdb, 0xba, 0x52, 0xc4, 0xae, 0xcc, 0x6d, 0x82, 0x54, 0x4a, 0x35, 0x3f,
	0x43, 0x61, 0x4d, 0xca, 0x92, 0x4d, 0xf6, 0x3a, 0x5e, 0xfa, 0x28, 0x3e, 0x8a, 0xf8, 0x10, 0xd2,
	0xd4, 0xed, 0x66, 0xe8, 0x4d, 0x38, 0xc9, 0xf9, 0xe0, 0xe4, 0xa0, 0x2b, 0x2e, 0x25, 0x9f, 0x83,
	0x9b, 0x31, 0xe5, 0xd6, 0xb0, 0x42, 0x2b, 0xcf, 0x05, 0xb1, 0x2c, 0x94, 0xfb, 0x02, 0xc0, 0xd2,
	0x5c, 0x43, 0x91, 0xea, 0x6c, 0xc1, 0x41, 0xa7, 0x7a, 0x5d, 0x02, 0x2e, 0x17, 0x52, 0x4b, 0xa7,
	0x5d, 0xeb, 0x71, 0xc6, 0x14, 0xde, 0x5a, 0xf1, 0xca, 0xc3, 0xc6, 0x7a, 0x7e, 0xb1, 0x49, 0x2e,
	0x73, 0x37, 0x13, 0x42, 0xea, 0x4c, 0xe7, 0x52, 0xa8, 0xda, 0xdc, 0x7d, 0x45, 0xa7, 0x43, 0x00,
	0x16, 0x6a, 0x28, 0x26, 0x26, 0x79, 0xb2, 0x2e, 0xc1, 0x17, 0xcb, 0xa2, 0xfb, 0x88, 0x9c, 0x5d,
	0xc6, 0x39, 0x41, 0xad, 0x69, 0x74, 0x17, 0xfb, 0xfd, 0x70, 0x18, 0xfa, 0x03, 0xfb, 0x9f, 0xd3,
	0x42, 0xfb, 0xd3, 0xe8, 0x26, 0x1a, 0xdd, 0x47, 0x76, 0xc3, 0x39, 0x42, 0x07, 0x7d, 0x7a, 0x1b,
	0xd3, 0x30, 0x88, 0x6c, 0xab, 0xba, 0xd1, 0x41, 0x1a, 0x8c, 0x47, 0xd3, 0xd8, 0x6e, 0x3a, 0xc7,
	0xe8, 0xb0, 0x3f, 0x0e, 0x27, 0xfe, 0x38, 0x1c, 0x45, 0xf6, 0xff, 0xde, 0x57, 0x03, 0x75, 0x9e,
	0x65, 0x81, 0xff, 0x2c, 0xdf, 0x3b, 0xdb, 0xad, 0x10, 0x57, 0xbd, 0xe3, 0xc6, 0x43, 0xef, 0xc7,
	0xc9, 0xe5, 0x3c, 0x13, 0x1c, 0xcb, 0x05, 0x77, 0x39, 0x08, 0xf3, 0xab, 0xcd, 0x82, 0x65, 0xae,
	0x7e, 0x19, 0xf4, 0xda, 0x9c, 0x6f, 0x56, 0x33, 0xa0, 0xf4, 0xdd, 0x6a, 0x07, 0x75, 0x14, 0x65,
	0x0a, 0xd7, 0xb0, 0x42, 0x33, 0x0f, 0x57, 0x43, 0xa8, 0x8f, 0x0d, 0x9f, 0x50, 0xa6, 0x92, 0x2d,
	0x9f, 0xcc, 0xbc, 0xc4, 0xf0, 0x9f, 0x56, 0xa7, 0x7e, 0x24, 0x84, 0x32, 0x45, 0xc8, 0x56, 0x41,
	0xc8, 0xcc, 0x23, 0xc4, 0x68, 0x9e, 0xf6, 0x4c, 0xb1, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0xfd, 0x5d, 0x51, 0xe8, 0x01, 0x00, 0x00,
}
