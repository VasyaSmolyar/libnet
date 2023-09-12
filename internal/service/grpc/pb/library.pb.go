// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/service/grpc/pb/library.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BookRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRequest) Reset()         { *m = BookRequest{} }
func (m *BookRequest) String() string { return proto.CompactTextString(m) }
func (*BookRequest) ProtoMessage()    {}
func (*BookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{0}
}

func (m *BookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookRequest.Unmarshal(m, b)
}
func (m *BookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookRequest.Marshal(b, m, deterministic)
}
func (m *BookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRequest.Merge(m, src)
}
func (m *BookRequest) XXX_Size() int {
	return xxx_messageInfo_BookRequest.Size(m)
}
func (m *BookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookRequest proto.InternalMessageInfo

func (m *BookRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type AuthorRequest struct {
	LastName             string   `protobuf:"bytes,1,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorRequest) Reset()         { *m = AuthorRequest{} }
func (m *AuthorRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorRequest) ProtoMessage()    {}
func (*AuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{1}
}

func (m *AuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorRequest.Unmarshal(m, b)
}
func (m *AuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorRequest.Marshal(b, m, deterministic)
}
func (m *AuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorRequest.Merge(m, src)
}
func (m *AuthorRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorRequest.Size(m)
}
func (m *AuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorRequest proto.InternalMessageInfo

func (m *AuthorRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type Book struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{2}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{3}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type BookResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookResponse) Reset()         { *m = BookResponse{} }
func (m *BookResponse) String() string { return proto.CompactTextString(m) }
func (*BookResponse) ProtoMessage()    {}
func (*BookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{4}
}

func (m *BookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookResponse.Unmarshal(m, b)
}
func (m *BookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookResponse.Marshal(b, m, deterministic)
}
func (m *BookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookResponse.Merge(m, src)
}
func (m *BookResponse) XXX_Size() int {
	return xxx_messageInfo_BookResponse.Size(m)
}
func (m *BookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookResponse proto.InternalMessageInfo

func (m *BookResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type AuthorResponse struct {
	Authors              []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AuthorResponse) Reset()         { *m = AuthorResponse{} }
func (m *AuthorResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorResponse) ProtoMessage()    {}
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b3c66e1403fcf3, []int{5}
}

func (m *AuthorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorResponse.Unmarshal(m, b)
}
func (m *AuthorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorResponse.Marshal(b, m, deterministic)
}
func (m *AuthorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorResponse.Merge(m, src)
}
func (m *AuthorResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorResponse.Size(m)
}
func (m *AuthorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorResponse proto.InternalMessageInfo

func (m *AuthorResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

func init() {
	proto.RegisterType((*BookRequest)(nil), "library.BookRequest")
	proto.RegisterType((*AuthorRequest)(nil), "library.AuthorRequest")
	proto.RegisterType((*Book)(nil), "library.Book")
	proto.RegisterType((*Author)(nil), "library.Author")
	proto.RegisterType((*BookResponse)(nil), "library.BookResponse")
	proto.RegisterType((*AuthorResponse)(nil), "library.AuthorResponse")
}

func init() {
	proto.RegisterFile("internal/service/grpc/pb/library.proto", fileDescriptor_c1b3c66e1403fcf3)
}

var fileDescriptor_c1b3c66e1403fcf3 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0x75, 0xab, 0x7b, 0x75, 0x13, 0xc2, 0xd4, 0xd2, 0xd3, 0xc8, 0x40, 0x26, 0xca,
	0x0a, 0x9b, 0x37, 0x45, 0x70, 0x07, 0x4f, 0xe2, 0x61, 0x47, 0x6f, 0xe9, 0x1a, 0x34, 0xac, 0x36,
	0x35, 0xc9, 0x04, 0xff, 0x02, 0xff, 0x6d, 0x69, 0x5f, 0x32, 0x67, 0xcb, 0x8e, 0x2f, 0xdf, 0xf7,
	0xfb, 0xde, 0x7b, 0x79, 0x70, 0x29, 0x0a, 0xc3, 0x55, 0xc1, 0xf2, 0x44, 0x73, 0xf5, 0x25, 0x56,
	0x3c, 0x79, 0x53, 0xe5, 0x2a, 0x29, 0xd3, 0x24, 0x17, 0xa9, 0x62, 0xea, 0x7b, 0x5a, 0x2a, 0x69,
	0x24, 0x09, 0x6c, 0x49, 0xc7, 0x10, 0x2e, 0xa4, 0x5c, 0x2f, 0xf9, 0xe7, 0x86, 0x6b, 0x43, 0x86,
	0xd0, 0x31, 0xc2, 0xe4, 0x3c, 0xf2, 0x46, 0xde, 0xa4, 0xb7, 0xc4, 0x82, 0x5e, 0x43, 0xff, 0x71,
	0x63, 0xde, 0xa5, 0x72, 0xb6, 0x18, 0x8e, 0x73, 0xa6, 0xcd, 0x0b, 0xfb, 0x70, 0xce, 0x6d, 0x4d,
	0x6f, 0xe0, 0xa8, 0x4a, 0x24, 0x03, 0xf0, 0x45, 0x66, 0x55, 0x5f, 0x64, 0x7f, 0xd1, 0xfe, 0x6e,
	0xf4, 0x2d, 0x74, 0x31, 0xba, 0xe5, 0xdf, 0xed, 0xe1, 0x37, 0x7a, 0xcc, 0xe1, 0x04, 0xa7, 0xd6,
	0xa5, 0x2c, 0x34, 0x27, 0x63, 0xe8, 0xa4, 0x52, 0xae, 0x75, 0xe4, 0x8d, 0x0e, 0x27, 0xe1, 0xac,
	0x3f, 0x75, 0xdb, 0xd6, 0x2e, 0xd4, 0xe8, 0x1d, 0x0c, 0xdc, 0x16, 0x16, 0xbb, 0x82, 0x80, 0xd5,
	0x2f, 0x0e, 0x3c, 0xdd, 0x82, 0xd6, 0xe9, 0xf4, 0xd9, 0x8f, 0x07, 0xc1, 0x33, 0x6a, 0xe4, 0x1e,
	0x7a, 0x4f, 0xa2, 0xc8, 0xaa, 0x6c, 0x4d, 0xce, 0x9b, 0x08, 0x7e, 0x51, 0x7c, 0xf6, 0x7f, 0x06,
	0xdb, 0x92, 0x1e, 0x90, 0x07, 0x08, 0x2b, 0x1a, 0xdd, 0x9a, 0x0c, 0x1b, 0x3e, 0xa4, 0x2f, 0x5a,
	0xa9, 0x8e, 0x5f, 0xc4, 0xaf, 0xd1, 0xbe, 0x23, 0xa7, 0xdd, 0xfa, 0xba, 0xf3, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x99, 0x33, 0x70, 0x07, 0x02, 0x00, 0x00,
}
