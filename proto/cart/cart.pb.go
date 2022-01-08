// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cart/cart.proto

/*
Package go_micro_service_cart is a generated protocol buffer package.

It is generated from these files:
	proto/cart/cart.proto

It has these top-level messages:
	CartInfo
	ResponseAdd
	Clean
	Response
	Item
	CartId
	CartFindAll
	CartAll
*/
package go_micro_service_cart

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

type CartInfo struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num" json:"num,omitempty"`
}

func (m *CartInfo) Reset()                    { *m = CartInfo{} }
func (m *CartInfo) String() string            { return proto.CompactTextString(m) }
func (*CartInfo) ProtoMessage()               {}
func (*CartInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CartInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CartInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CartInfo) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *CartInfo) GetSizeId() int64 {
	if m != nil {
		return m.SizeId
	}
	return 0
}

func (m *CartInfo) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ResponseAdd struct {
	CartId int64  `protobuf:"varint,1,opt,name=cart_id,json=cartId" json:"cart_id,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResponseAdd) Reset()                    { *m = ResponseAdd{} }
func (m *ResponseAdd) String() string            { return proto.CompactTextString(m) }
func (*ResponseAdd) ProtoMessage()               {}
func (*ResponseAdd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseAdd) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *ResponseAdd) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Clean struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Clean) Reset()                    { *m = Clean{} }
func (m *Clean) String() string            { return proto.CompactTextString(m) }
func (*Clean) ProtoMessage()               {}
func (*Clean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Clean) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Item struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ChangeNum int64 `protobuf:"varint,2,opt,name=change_num,json=changeNum" json:"change_num,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetChangeNum() int64 {
	if m != nil {
		return m.ChangeNum
	}
	return 0
}

type CartId struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CartId) Reset()                    { *m = CartId{} }
func (m *CartId) String() string            { return proto.CompactTextString(m) }
func (*CartId) ProtoMessage()               {}
func (*CartId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CartId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CartFindAll struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *CartFindAll) Reset()                    { *m = CartFindAll{} }
func (m *CartFindAll) String() string            { return proto.CompactTextString(m) }
func (*CartFindAll) ProtoMessage()               {}
func (*CartFindAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CartFindAll) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CartAll struct {
	CartInfo []*CartInfo `protobuf:"bytes,1,rep,name=cart_info,json=cartInfo" json:"cart_info,omitempty"`
}

func (m *CartAll) Reset()                    { *m = CartAll{} }
func (m *CartAll) String() string            { return proto.CompactTextString(m) }
func (*CartAll) ProtoMessage()               {}
func (*CartAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CartAll) GetCartInfo() []*CartInfo {
	if m != nil {
		return m.CartInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CartInfo)(nil), "go.micro.service.cart.CartInfo")
	proto.RegisterType((*ResponseAdd)(nil), "go.micro.service.cart.ResponseAdd")
	proto.RegisterType((*Clean)(nil), "go.micro.service.cart.Clean")
	proto.RegisterType((*Response)(nil), "go.micro.service.cart.Response")
	proto.RegisterType((*Item)(nil), "go.micro.service.cart.Item")
	proto.RegisterType((*CartId)(nil), "go.micro.service.cart.CartId")
	proto.RegisterType((*CartFindAll)(nil), "go.micro.service.cart.CartFindAll")
	proto.RegisterType((*CartAll)(nil), "go.micro.service.cart.CartAll")
}

func init() { proto.RegisterFile("proto/cart/cart.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x8b, 0xda, 0x40,
	0x14, 0x35, 0x26, 0x46, 0x73, 0x05, 0x29, 0x03, 0xd2, 0x60, 0xb5, 0xca, 0x3c, 0x14, 0x9f, 0x52,
	0xb0, 0x14, 0xfa, 0xd0, 0x97, 0x54, 0x51, 0x02, 0x45, 0x4a, 0xfe, 0x80, 0xa4, 0x99, 0xd1, 0x0d,
	0x24, 0x33, 0x32, 0x49, 0x16, 0x76, 0x61, 0xff, 0xe8, 0xfe, 0x9a, 0xe5, 0x4e, 0x74, 0x57, 0x76,
	0x37, 0xea, 0xc3, 0xbe, 0xc8, 0xfd, 0x38, 0xe7, 0x70, 0xee, 0x1c, 0x03, 0xfd, 0xbd, 0x92, 0x85,
	0xfc, 0x1e, 0x47, 0xaa, 0xd0, 0x3f, 0x9e, 0xee, 0x49, 0x7f, 0x27, 0xbd, 0x2c, 0x89, 0x95, 0xf4,
	0x72, 0xae, 0x6e, 0x93, 0x98, 0x7b, 0xb8, 0xa4, 0x0f, 0xd0, 0x99, 0x47, 0xaa, 0x08, 0xc4, 0x56,
	0x92, 0x1e, 0x34, 0x13, 0xe6, 0x1a, 0x13, 0x63, 0x6a, 0x86, 0xcd, 0x84, 0x91, 0xcf, 0xd0, 0x2e,
	0x73, 0xae, 0x36, 0x09, 0x73, 0x9b, 0x7a, 0x68, 0x63, 0x1b, 0x30, 0x32, 0x02, 0xd8, 0x2b, 0xc9,
	0xca, 0xb8, 0xc0, 0x9d, 0xa9, 0x77, 0xce, 0x61, 0x12, 0x68, 0x5e, 0x9e, 0xdc, 0x73, 0xdc, 0x59,
	0x15, 0x0f, 0xdb, 0x80, 0x91, 0x4f, 0x60, 0x8a, 0x32, 0x73, 0x5b, 0x7a, 0x88, 0x25, 0xfd, 0x05,
	0xdd, 0x90, 0xe7, 0x7b, 0x29, 0x72, 0xee, 0x33, 0xcd, 0x44, 0x57, 0x9b, 0x67, 0x1b, 0x36, 0xb6,
	0x15, 0x33, 0xcb, 0x77, 0xda, 0x86, 0x13, 0x62, 0x49, 0x27, 0xd0, 0x9a, 0xa7, 0x3c, 0x12, 0xa7,
	0x2e, 0x8d, 0x53, 0x97, 0x74, 0x08, 0x9d, 0xa3, 0xf6, 0x91, 0x6f, 0xbc, 0xf0, 0x7f, 0x82, 0x15,
	0x14, 0x3c, 0x7b, 0x73, 0xf4, 0x08, 0x20, 0xbe, 0x89, 0xc4, 0x8e, 0x6f, 0xd0, 0x6a, 0x75, 0xb7,
	0x53, 0x4d, 0xd6, 0x65, 0x46, 0x5d, 0xb0, 0xe7, 0x95, 0xa5, 0x57, 0x44, 0xfa, 0x0d, 0xba, 0xb8,
	0x59, 0x26, 0x82, 0xf9, 0x69, 0x5a, 0x6f, 0x6b, 0x05, 0x6d, 0xc4, 0x21, 0xe6, 0x37, 0x38, 0xd5,
	0xb9, 0x62, 0x2b, 0x5d, 0x63, 0x62, 0x4e, 0xbb, 0xb3, 0xb1, 0xf7, 0x6e, 0x4e, 0xde, 0x31, 0xa4,
	0xb0, 0x13, 0x1f, 0xaa, 0xd9, 0xa3, 0x09, 0x16, 0x8e, 0xc9, 0x3f, 0x68, 0xfb, 0x8c, 0xe9, 0xf2,
	0x12, 0x7d, 0x40, 0x6b, 0x00, 0x27, 0x29, 0xd0, 0x06, 0xf9, 0x0b, 0x8e, 0x7e, 0x5c, 0xad, 0x39,
	0xac, 0xd3, 0x44, 0xc4, 0x60, 0x7c, 0x41, 0x90, 0x36, 0xc8, 0x12, 0xac, 0x40, 0xc4, 0x8a, 0x7c,
	0xa9, 0x81, 0x62, 0x0e, 0x57, 0xea, 0x2c, 0xf8, 0x07, 0xe8, 0x84, 0xd0, 0x5b, 0xf0, 0x94, 0x17,
	0x1c, 0x09, 0x7f, 0xee, 0x82, 0x05, 0x19, 0x9d, 0x7b, 0x36, 0x76, 0x8d, 0xe6, 0x1a, 0xec, 0x15,
	0xd7, 0xa1, 0xd2, 0x33, 0x5a, 0x87, 0x3f, 0xc7, 0xe0, 0xeb, 0x19, 0x8c, 0x9f, 0xa6, 0xb4, 0xf1,
	0xdf, 0xd6, 0x5f, 0xed, 0x8f, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x26, 0x92, 0xd4, 0xce,
	0x03, 0x00, 0x00,
}
