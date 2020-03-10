// Code generated by protoc-gen-go. DO NOT EDIT.
// source: starwars/starwars.proto

package starwars

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_HUMAN Type = 0
	Type_DROID Type = 1
)

var Type_name = map[int32]string{
	0: "HUMAN",
	1: "DROID",
}

var Type_value = map[string]int32{
	"HUMAN": 0,
	"DROID": 1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{0}
}

type Episode int32

const (
	Episode__       Episode = 0
	Episode_NEWHOPE Episode = 1
	Episode_EMPIRE  Episode = 2
	Episode_JEDI    Episode = 3
)

var Episode_name = map[int32]string{
	0: "_",
	1: "NEWHOPE",
	2: "EMPIRE",
	3: "JEDI",
}

var Episode_value = map[string]int32{
	"_":       0,
	"NEWHOPE": 1,
	"EMPIRE":  2,
	"JEDI":    3,
}

func (x Episode) String() string {
	return proto.EnumName(Episode_name, int32(x))
}

func (Episode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{1}
}

type Character struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends              []*Character `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
	AppearsIn            []Episode    `protobuf:"varint,4,rep,packed,name=appears_in,json=appearsIn,proto3,enum=startwars.Episode" json:"appears_in,omitempty"`
	HomePlanet           string       `protobuf:"bytes,5,opt,name=home_planet,json=homePlanet,proto3" json:"home_planet,omitempty"`
	PrimaryFunction      string       `protobuf:"bytes,6,opt,name=primary_function,json=primaryFunction,proto3" json:"primary_function,omitempty"`
	Type                 Type         `protobuf:"varint,7,opt,name=type,proto3,enum=startwars.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Character) Reset()         { *m = Character{} }
func (m *Character) String() string { return proto.CompactTextString(m) }
func (*Character) ProtoMessage()    {}
func (*Character) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{0}
}

func (m *Character) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Character.Unmarshal(m, b)
}
func (m *Character) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Character.Marshal(b, m, deterministic)
}
func (m *Character) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Character.Merge(m, src)
}
func (m *Character) XXX_Size() int {
	return xxx_messageInfo_Character.Size(m)
}
func (m *Character) XXX_DiscardUnknown() {
	xxx_messageInfo_Character.DiscardUnknown(m)
}

var xxx_messageInfo_Character proto.InternalMessageInfo

func (m *Character) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Character) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Character) GetFriends() []*Character {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *Character) GetAppearsIn() []Episode {
	if m != nil {
		return m.AppearsIn
	}
	return nil
}

func (m *Character) GetHomePlanet() string {
	if m != nil {
		return m.HomePlanet
	}
	return ""
}

func (m *Character) GetPrimaryFunction() string {
	if m != nil {
		return m.PrimaryFunction
	}
	return ""
}

func (m *Character) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_HUMAN
}

type GetHeroRequest struct {
	// If omitted, returns the hero of the whope saga. If provided, returns the hero of that particular episode.
	Episode              Episode  `protobuf:"varint,1,opt,name=episode,proto3,enum=startwars.Episode" json:"episode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHeroRequest) Reset()         { *m = GetHeroRequest{} }
func (m *GetHeroRequest) String() string { return proto.CompactTextString(m) }
func (*GetHeroRequest) ProtoMessage()    {}
func (*GetHeroRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{1}
}

func (m *GetHeroRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeroRequest.Unmarshal(m, b)
}
func (m *GetHeroRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeroRequest.Marshal(b, m, deterministic)
}
func (m *GetHeroRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeroRequest.Merge(m, src)
}
func (m *GetHeroRequest) XXX_Size() int {
	return xxx_messageInfo_GetHeroRequest.Size(m)
}
func (m *GetHeroRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeroRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeroRequest proto.InternalMessageInfo

func (m *GetHeroRequest) GetEpisode() Episode {
	if m != nil {
		return m.Episode
	}
	return Episode__
}

type GetHumanRequest struct {
	// id of the human
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHumanRequest) Reset()         { *m = GetHumanRequest{} }
func (m *GetHumanRequest) String() string { return proto.CompactTextString(m) }
func (*GetHumanRequest) ProtoMessage()    {}
func (*GetHumanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{2}
}

func (m *GetHumanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHumanRequest.Unmarshal(m, b)
}
func (m *GetHumanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHumanRequest.Marshal(b, m, deterministic)
}
func (m *GetHumanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHumanRequest.Merge(m, src)
}
func (m *GetHumanRequest) XXX_Size() int {
	return xxx_messageInfo_GetHumanRequest.Size(m)
}
func (m *GetHumanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHumanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHumanRequest proto.InternalMessageInfo

func (m *GetHumanRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetDroidRequest struct {
	// id of the droid
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDroidRequest) Reset()         { *m = GetDroidRequest{} }
func (m *GetDroidRequest) String() string { return proto.CompactTextString(m) }
func (*GetDroidRequest) ProtoMessage()    {}
func (*GetDroidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{3}
}

func (m *GetDroidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDroidRequest.Unmarshal(m, b)
}
func (m *GetDroidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDroidRequest.Marshal(b, m, deterministic)
}
func (m *GetDroidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDroidRequest.Merge(m, src)
}
func (m *GetDroidRequest) XXX_Size() int {
	return xxx_messageInfo_GetDroidRequest.Size(m)
}
func (m *GetDroidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDroidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDroidRequest proto.InternalMessageInfo

func (m *GetDroidRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListEmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEmptyRequest) Reset()         { *m = ListEmptyRequest{} }
func (m *ListEmptyRequest) String() string { return proto.CompactTextString(m) }
func (*ListEmptyRequest) ProtoMessage()    {}
func (*ListEmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{4}
}

func (m *ListEmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEmptyRequest.Unmarshal(m, b)
}
func (m *ListEmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEmptyRequest.Marshal(b, m, deterministic)
}
func (m *ListEmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEmptyRequest.Merge(m, src)
}
func (m *ListEmptyRequest) XXX_Size() int {
	return xxx_messageInfo_ListEmptyRequest.Size(m)
}
func (m *ListEmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEmptyRequest proto.InternalMessageInfo

type ListHumansResponse struct {
	Humans               []*Character `protobuf:"bytes,1,rep,name=humans,proto3" json:"humans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListHumansResponse) Reset()         { *m = ListHumansResponse{} }
func (m *ListHumansResponse) String() string { return proto.CompactTextString(m) }
func (*ListHumansResponse) ProtoMessage()    {}
func (*ListHumansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{5}
}

func (m *ListHumansResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHumansResponse.Unmarshal(m, b)
}
func (m *ListHumansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHumansResponse.Marshal(b, m, deterministic)
}
func (m *ListHumansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHumansResponse.Merge(m, src)
}
func (m *ListHumansResponse) XXX_Size() int {
	return xxx_messageInfo_ListHumansResponse.Size(m)
}
func (m *ListHumansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHumansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHumansResponse proto.InternalMessageInfo

func (m *ListHumansResponse) GetHumans() []*Character {
	if m != nil {
		return m.Humans
	}
	return nil
}

type ListDroidsResponse struct {
	Droids               []*Character `protobuf:"bytes,1,rep,name=droids,proto3" json:"droids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListDroidsResponse) Reset()         { *m = ListDroidsResponse{} }
func (m *ListDroidsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDroidsResponse) ProtoMessage()    {}
func (*ListDroidsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da60d3f39fcaf9, []int{6}
}

func (m *ListDroidsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDroidsResponse.Unmarshal(m, b)
}
func (m *ListDroidsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDroidsResponse.Marshal(b, m, deterministic)
}
func (m *ListDroidsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDroidsResponse.Merge(m, src)
}
func (m *ListDroidsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDroidsResponse.Size(m)
}
func (m *ListDroidsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDroidsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDroidsResponse proto.InternalMessageInfo

func (m *ListDroidsResponse) GetDroids() []*Character {
	if m != nil {
		return m.Droids
	}
	return nil
}

func init() {
	proto.RegisterEnum("startwars.Type", Type_name, Type_value)
	proto.RegisterEnum("startwars.Episode", Episode_name, Episode_value)
	proto.RegisterType((*Character)(nil), "startwars.Character")
	proto.RegisterType((*GetHeroRequest)(nil), "startwars.GetHeroRequest")
	proto.RegisterType((*GetHumanRequest)(nil), "startwars.GetHumanRequest")
	proto.RegisterType((*GetDroidRequest)(nil), "startwars.GetDroidRequest")
	proto.RegisterType((*ListEmptyRequest)(nil), "startwars.ListEmptyRequest")
	proto.RegisterType((*ListHumansResponse)(nil), "startwars.ListHumansResponse")
	proto.RegisterType((*ListDroidsResponse)(nil), "startwars.ListDroidsResponse")
}

func init() { proto.RegisterFile("starwars/starwars.proto", fileDescriptor_d6da60d3f39fcaf9) }

var fileDescriptor_d6da60d3f39fcaf9 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xda, 0x4a,
	0x14, 0x8d, 0xf9, 0x32, 0x5c, 0x24, 0xb0, 0xee, 0x7b, 0x2f, 0xf1, 0xe3, 0xbd, 0xaa, 0x88, 0x6e,
	0x48, 0x94, 0x40, 0x42, 0x94, 0x4d, 0x17, 0x95, 0x9a, 0xe2, 0x14, 0xa2, 0xe6, 0x43, 0x4e, 0xab,
	0x4a, 0xdd, 0xd0, 0x89, 0x99, 0x60, 0x4b, 0xd8, 0x33, 0x19, 0x0f, 0x4d, 0x59, 0x77, 0xc7, 0xcf,
	0xe8, 0x7f, 0xe0, 0x67, 0xf4, 0x3f, 0x55, 0x1e, 0xc6, 0x09, 0xa0, 0x34, 0x55, 0x57, 0xbe, 0xbe,
	0x73, 0xee, 0x39, 0x57, 0xe7, 0x8c, 0x06, 0xb6, 0x62, 0x49, 0xc4, 0x1d, 0x11, 0x71, 0x3b, 0x2d,
	0x5a, 0x5c, 0x30, 0xc9, 0xb0, 0x94, 0xfc, 0xcb, 0xa4, 0x51, 0xfb, 0x67, 0x24, 0x08, 0xf7, 0x6f,
	0xc7, 0x6d, 0xfd, 0x5d, 0x20, 0x1a, 0xdf, 0x32, 0x50, 0x7a, 0xe3, 0x13, 0x41, 0x3c, 0x49, 0x05,
	0x56, 0x20, 0x13, 0x0c, 0x6d, 0xa3, 0x6e, 0x34, 0xb3, 0x6e, 0x26, 0x18, 0x22, 0x42, 0x2e, 0x22,
	0x21, 0xb5, 0x33, 0x75, 0xa3, 0x59, 0x72, 0x55, 0x8d, 0x2d, 0x30, 0x6f, 0x44, 0x40, 0xa3, 0x61,
	0x6c, 0x67, 0xeb, 0xd9, 0x66, 0xb9, 0xf3, 0x77, 0xeb, 0x5e, 0xa5, 0x75, 0x4f, 0xe5, 0xa6, 0x20,
	0x3c, 0x00, 0x20, 0x9c, 0x53, 0x22, 0xe2, 0x41, 0x10, 0xd9, 0xb9, 0x7a, 0xb6, 0x59, 0xe9, 0xe0,
	0xd2, 0x88, 0xc3, 0x83, 0x98, 0x0d, 0xa9, 0x5b, 0xd2, 0xa8, 0x7e, 0x84, 0xcf, 0xa1, 0xec, 0xb3,
	0x90, 0x0e, 0xf8, 0x98, 0x44, 0x54, 0xda, 0x79, 0xa5, 0x0e, 0x49, 0xeb, 0x52, 0x75, 0x70, 0x1b,
	0x2c, 0x2e, 0x82, 0x90, 0x88, 0xe9, 0xe0, 0x66, 0x12, 0x79, 0x32, 0x60, 0x91, 0x5d, 0x50, 0xa8,
	0xaa, 0xee, 0x9f, 0xe8, 0x36, 0xbe, 0x80, 0x9c, 0x9c, 0x72, 0x6a, 0x9b, 0x75, 0xa3, 0x59, 0xe9,
	0x54, 0x97, 0x84, 0xdf, 0x4f, 0x39, 0x75, 0xd5, 0x61, 0xe3, 0x15, 0x54, 0xde, 0x52, 0xd9, 0xa3,
	0x82, 0xb9, 0xf4, 0x76, 0x42, 0x63, 0x89, 0xbb, 0x60, 0xd2, 0xc5, 0x62, 0xca, 0x8e, 0xc7, 0x57,
	0x4e, 0x21, 0x8d, 0x6d, 0xa8, 0x26, 0xf3, 0x93, 0x90, 0x44, 0x29, 0xc1, 0xe6, 0x83, 0x95, 0xc7,
	0x85, 0xef, 0x73, 0x3b, 0x53, 0x34, 0x12, 0x4b, 0x35, 0xb4, 0x2b, 0x58, 0x30, 0xfc, 0x1d, 0x14,
	0xc1, 0x7a, 0x17, 0xc4, 0xd2, 0x09, 0xb9, 0x9c, 0x6a, 0x6c, 0xe3, 0x18, 0x30, 0xe9, 0x29, 0xa9,
	0xd8, 0xa5, 0x31, 0x67, 0x51, 0x4c, 0x71, 0x17, 0x0a, 0xbe, 0xea, 0xd8, 0xc6, 0x13, 0x91, 0x68,
	0x4c, 0xca, 0xa1, 0x76, 0x58, 0xe1, 0x18, 0xaa, 0xce, 0xd3, 0x1c, 0x0b, 0xcc, 0xce, 0xff, 0x90,
	0x4b, 0xfc, 0xc3, 0x12, 0xe4, 0x7b, 0x1f, 0xce, 0x5e, 0x9f, 0x5b, 0x1b, 0x49, 0xd9, 0x75, 0x2f,
	0xfa, 0x5d, 0xcb, 0xd8, 0x39, 0x04, 0x53, 0x7b, 0x84, 0x79, 0x30, 0x06, 0xd6, 0x06, 0x96, 0xc1,
	0x3c, 0x77, 0x3e, 0xf6, 0x2e, 0x2e, 0x1d, 0xcb, 0x40, 0x80, 0x82, 0x73, 0x76, 0xd9, 0x77, 0x1d,
	0x2b, 0x83, 0x45, 0xc8, 0x9d, 0x3a, 0xdd, 0xbe, 0x95, 0xed, 0xfc, 0xc8, 0x82, 0x75, 0x95, 0x4a,
	0x5e, 0x51, 0xf1, 0x25, 0xf0, 0x28, 0x9e, 0x80, 0xa9, 0x93, 0xc1, 0x7f, 0x97, 0x16, 0x5a, 0x4d,
	0xab, 0xf6, 0xe8, 0xae, 0x0d, 0x98, 0xcd, 0xed, 0x02, 0xe4, 0xfc, 0x64, 0xb8, 0x0f, 0xc5, 0x34,
	0x21, 0xac, 0xad, 0x11, 0x2d, 0xc5, 0xf6, 0x0b, 0xa6, 0xf2, 0x6c, 0x6e, 0x9b, 0x90, 0x57, 0xfe,
	0x69, 0x2a, 0xe5, 0xde, 0x3a, 0xd5, 0x72, 0xac, 0x4f, 0x53, 0x29, 0x1b, 0xf1, 0x33, 0xc0, 0x43,
	0x9a, 0xf8, 0xdf, 0xd2, 0xc0, 0x7a, 0xf0, 0xb5, 0x67, 0x6b, 0x87, 0xab, 0x37, 0xa0, 0xb1, 0x39,
	0x9b, 0xdb, 0x98, 0xde, 0x82, 0x5a, 0x11, 0x75, 0x95, 0x2a, 0x2c, 0xb2, 0xfe, 0x33, 0x85, 0xd5,
	0xfb, 0x91, 0x2a, 0x2c, 0xf2, 0x4f, 0x14, 0x74, 0xb5, 0x35, 0x9b, 0xdb, 0x7f, 0x41, 0x75, 0xcc,
	0x3c, 0x32, 0xf6, 0x59, 0x2c, 0x5f, 0x1e, 0xed, 0xef, 0x1f, 0x1d, 0x58, 0xc6, 0xf1, 0xe9, 0xa7,
	0xde, 0x28, 0x90, 0xfe, 0xe4, 0xba, 0xe5, 0xb1, 0xb0, 0x3d, 0x8d, 0x27, 0xa3, 0x20, 0x64, 0x92,
	0xb5, 0x47, 0x82, 0x7b, 0x7b, 0xfa, 0x15, 0xda, 0x1b, 0x11, 0x49, 0xef, 0xc8, 0xb4, 0x4d, 0xbf,
	0x92, 0x90, 0x8f, 0x69, 0xfb, 0xe1, 0x25, 0xe3, 0xd4, 0xbb, 0xff, 0xbb, 0x2e, 0xa8, 0xd7, 0xea,
	0xf0, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x3f, 0x59, 0xc7, 0xea, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StartwarsServiceClient is the client API for StartwarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StartwarsServiceClient interface {
	GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error)
	GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error)
	GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error)
	ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error)
	ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error)
}

type startwarsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStartwarsServiceClient(cc *grpc.ClientConn) StartwarsServiceClient {
	return &startwarsServiceClient{cc}
}

func (c *startwarsServiceClient) GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetDroid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error) {
	out := new(ListHumansResponse)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/ListHumans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error) {
	out := new(ListDroidsResponse)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/ListDroids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartwarsServiceServer is the server API for StartwarsService service.
type StartwarsServiceServer interface {
	GetHero(context.Context, *GetHeroRequest) (*Character, error)
	GetHuman(context.Context, *GetHumanRequest) (*Character, error)
	GetDroid(context.Context, *GetDroidRequest) (*Character, error)
	ListHumans(context.Context, *ListEmptyRequest) (*ListHumansResponse, error)
	ListDroids(context.Context, *ListEmptyRequest) (*ListDroidsResponse, error)
}

// UnimplementedStartwarsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStartwarsServiceServer struct {
}

func (*UnimplementedStartwarsServiceServer) GetHero(ctx context.Context, req *GetHeroRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHero not implemented")
}
func (*UnimplementedStartwarsServiceServer) GetHuman(ctx context.Context, req *GetHumanRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHuman not implemented")
}
func (*UnimplementedStartwarsServiceServer) GetDroid(ctx context.Context, req *GetDroidRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDroid not implemented")
}
func (*UnimplementedStartwarsServiceServer) ListHumans(ctx context.Context, req *ListEmptyRequest) (*ListHumansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHumans not implemented")
}
func (*UnimplementedStartwarsServiceServer) ListDroids(ctx context.Context, req *ListEmptyRequest) (*ListDroidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDroids not implemented")
}

func RegisterStartwarsServiceServer(s *grpc.Server, srv StartwarsServiceServer) {
	s.RegisterService(&_StartwarsService_serviceDesc, srv)
}

func _StartwarsService_GetHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHero(ctx, req.(*GetHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHumanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHuman(ctx, req.(*GetHumanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetDroid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDroidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetDroid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetDroid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetDroid(ctx, req.(*GetDroidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListHumans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListHumans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/ListHumans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListHumans(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListDroids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListDroids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/ListDroids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListDroids(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StartwarsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "startwars.StartwarsService",
	HandlerType: (*StartwarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHero",
			Handler:    _StartwarsService_GetHero_Handler,
		},
		{
			MethodName: "GetHuman",
			Handler:    _StartwarsService_GetHuman_Handler,
		},
		{
			MethodName: "GetDroid",
			Handler:    _StartwarsService_GetDroid_Handler,
		},
		{
			MethodName: "ListHumans",
			Handler:    _StartwarsService_ListHumans_Handler,
		},
		{
			MethodName: "ListDroids",
			Handler:    _StartwarsService_ListDroids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starwars/starwars.proto",
}