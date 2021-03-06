// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplier.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	supplier.proto

It has these top-level messages:
	AddSupplierReq
	ReadByNamaSupplierReq
	ReadByNamaSupplierResp
	ReadByKeteranganReq
	ReadByKeteranganResp
	ReadSupplierResp
	UpdateSupplierReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddSupplierReq struct {
	NamaSupplier string `protobuf:"bytes,1,opt,name=namaSupplier" json:"namaSupplier,omitempty"`
	Alamat       string `protobuf:"bytes,2,opt,name=alamat" json:"alamat,omitempty"`
	Telepon      string `protobuf:"bytes,3,opt,name=telepon" json:"telepon,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Status       int32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	CreatedBy    string `protobuf:"bytes,6,opt,name=createdBy" json:"createdBy,omitempty"`
	Keterangan   string `protobuf:"bytes,7,opt,name=keterangan" json:"keterangan,omitempty"`
}

func (m *AddSupplierReq) Reset()                    { *m = AddSupplierReq{} }
func (m *AddSupplierReq) String() string            { return proto.CompactTextString(m) }
func (*AddSupplierReq) ProtoMessage()               {}
func (*AddSupplierReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddSupplierReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *AddSupplierReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *AddSupplierReq) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *AddSupplierReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddSupplierReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddSupplierReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddSupplierReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadByNamaSupplierReq struct {
	NamaSupplier string `protobuf:"bytes,1,opt,name=namaSupplier" json:"namaSupplier,omitempty"`
}

func (m *ReadByNamaSupplierReq) Reset()                    { *m = ReadByNamaSupplierReq{} }
func (m *ReadByNamaSupplierReq) String() string            { return proto.CompactTextString(m) }
func (*ReadByNamaSupplierReq) ProtoMessage()               {}
func (*ReadByNamaSupplierReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadByNamaSupplierReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

type ReadByNamaSupplierResp struct {
	IDSupplier   int32  `protobuf:"varint,1,opt,name=IDSupplier" json:"IDSupplier,omitempty"`
	NamaSupplier string `protobuf:"bytes,2,opt,name=namaSupplier" json:"namaSupplier,omitempty"`
	Alamat       string `protobuf:"bytes,3,opt,name=alamat" json:"alamat,omitempty"`
	Telepon      string `protobuf:"bytes,4,opt,name=telepon" json:"telepon,omitempty"`
	Email        string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Status       int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
	CreatedBy    string `protobuf:"bytes,7,opt,name=createdBy" json:"createdBy,omitempty"`
	Keterangan   string `protobuf:"bytes,8,opt,name=keterangan" json:"keterangan,omitempty"`
}

func (m *ReadByNamaSupplierResp) Reset()                    { *m = ReadByNamaSupplierResp{} }
func (m *ReadByNamaSupplierResp) String() string            { return proto.CompactTextString(m) }
func (*ReadByNamaSupplierResp) ProtoMessage()               {}
func (*ReadByNamaSupplierResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadByNamaSupplierResp) GetIDSupplier() int32 {
	if m != nil {
		return m.IDSupplier
	}
	return 0
}

func (m *ReadByNamaSupplierResp) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *ReadByNamaSupplierResp) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *ReadByNamaSupplierResp) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *ReadByNamaSupplierResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReadByNamaSupplierResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadByNamaSupplierResp) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadByNamaSupplierResp) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadByKeteranganReq struct {
	IDSupplier   int32  `protobuf:"varint,1,opt,name=IDSupplier" json:"IDSupplier,omitempty"`
	NamaSupplier string `protobuf:"bytes,2,opt,name=namaSupplier" json:"namaSupplier,omitempty"`
	Alamat       string `protobuf:"bytes,3,opt,name=alamat" json:"alamat,omitempty"`
	Telepon      string `protobuf:"bytes,4,opt,name=telepon" json:"telepon,omitempty"`
	Email        string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Status       int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
	CreatedBy    string `protobuf:"bytes,7,opt,name=createdBy" json:"createdBy,omitempty"`
	Keterangan   string `protobuf:"bytes,8,opt,name=keterangan" json:"keterangan,omitempty"`
}

func (m *ReadByKeteranganReq) Reset()                    { *m = ReadByKeteranganReq{} }
func (m *ReadByKeteranganReq) String() string            { return proto.CompactTextString(m) }
func (*ReadByKeteranganReq) ProtoMessage()               {}
func (*ReadByKeteranganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadByKeteranganReq) GetIDSupplier() int32 {
	if m != nil {
		return m.IDSupplier
	}
	return 0
}

func (m *ReadByKeteranganReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *ReadByKeteranganReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *ReadByKeteranganReq) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *ReadByKeteranganReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReadByKeteranganReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadByKeteranganReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadByKeteranganReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadByKeteranganResp struct {
	AllKeterangan []*ReadByNamaSupplierResp `protobuf:"bytes,1,rep,name=allKeterangan" json:"allKeterangan,omitempty"`
}

func (m *ReadByKeteranganResp) Reset()                    { *m = ReadByKeteranganResp{} }
func (m *ReadByKeteranganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadByKeteranganResp) ProtoMessage()               {}
func (*ReadByKeteranganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadByKeteranganResp) GetAllKeterangan() []*ReadByNamaSupplierResp {
	if m != nil {
		return m.AllKeterangan
	}
	return nil
}

type ReadSupplierResp struct {
	AllSupplier []*ReadByNamaSupplierResp `protobuf:"bytes,1,rep,name=allSupplier" json:"allSupplier,omitempty"`
}

func (m *ReadSupplierResp) Reset()                    { *m = ReadSupplierResp{} }
func (m *ReadSupplierResp) String() string            { return proto.CompactTextString(m) }
func (*ReadSupplierResp) ProtoMessage()               {}
func (*ReadSupplierResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadSupplierResp) GetAllSupplier() []*ReadByNamaSupplierResp {
	if m != nil {
		return m.AllSupplier
	}
	return nil
}

type UpdateSupplierReq struct {
	IDSupplier   int32  `protobuf:"varint,1,opt,name=IDSupplier" json:"IDSupplier,omitempty"`
	NamaSupplier string `protobuf:"bytes,2,opt,name=namaSupplier" json:"namaSupplier,omitempty"`
	Alamat       string `protobuf:"bytes,3,opt,name=alamat" json:"alamat,omitempty"`
	Telepon      string `protobuf:"bytes,4,opt,name=telepon" json:"telepon,omitempty"`
	Email        string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Status       int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
	UpdateBy     string `protobuf:"bytes,7,opt,name=updateBy" json:"updateBy,omitempty"`
	Keterangan   string `protobuf:"bytes,8,opt,name=keterangan" json:"keterangan,omitempty"`
}

func (m *UpdateSupplierReq) Reset()                    { *m = UpdateSupplierReq{} }
func (m *UpdateSupplierReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateSupplierReq) ProtoMessage()               {}
func (*UpdateSupplierReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateSupplierReq) GetIDSupplier() int32 {
	if m != nil {
		return m.IDSupplier
	}
	return 0
}

func (m *UpdateSupplierReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *UpdateSupplierReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *UpdateSupplierReq) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *UpdateSupplierReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateSupplierReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateSupplierReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *UpdateSupplierReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func init() {
	proto.RegisterType((*AddSupplierReq)(nil), "grpc.AddSupplierReq")
	proto.RegisterType((*ReadByNamaSupplierReq)(nil), "grpc.ReadByNamaSupplierReq")
	proto.RegisterType((*ReadByNamaSupplierResp)(nil), "grpc.ReadByNamaSupplierResp")
	proto.RegisterType((*ReadByKeteranganReq)(nil), "grpc.ReadByKeteranganReq")
	proto.RegisterType((*ReadByKeteranganResp)(nil), "grpc.ReadByKeteranganResp")
	proto.RegisterType((*ReadSupplierResp)(nil), "grpc.ReadSupplierResp")
	proto.RegisterType((*UpdateSupplierReq)(nil), "grpc.UpdateSupplierReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for SupplierService service

type SupplierServiceClient interface {
	AddSupplier(ctx context.Context, in *AddSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadByNamaSupplier(ctx context.Context, in *ReadByNamaSupplierReq, opts ...grpc1.CallOption) (*ReadByNamaSupplierResp, error)
	ReadSupplier(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadSupplierResp, error)
	UpdateSupplier(ctx context.Context, in *UpdateSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadByKeterangan(ctx context.Context, in *ReadByKeteranganReq, opts ...grpc1.CallOption) (*ReadByKeteranganResp, error)
}

type supplierServiceClient struct {
	cc *grpc1.ClientConn
}

func NewSupplierServiceClient(cc *grpc1.ClientConn) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) AddSupplier(ctx context.Context, in *AddSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/AddSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) ReadByNamaSupplier(ctx context.Context, in *ReadByNamaSupplierReq, opts ...grpc1.CallOption) (*ReadByNamaSupplierResp, error) {
	out := new(ReadByNamaSupplierResp)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/ReadByNamaSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) ReadSupplier(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadSupplierResp, error) {
	out := new(ReadSupplierResp)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/ReadSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) UpdateSupplier(ctx context.Context, in *UpdateSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/UpdateSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) ReadByKeterangan(ctx context.Context, in *ReadByKeteranganReq, opts ...grpc1.CallOption) (*ReadByKeteranganResp, error) {
	out := new(ReadByKeteranganResp)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/ReadByKeterangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SupplierService service

type SupplierServiceServer interface {
	AddSupplier(context.Context, *AddSupplierReq) (*google_protobuf.Empty, error)
	ReadByNamaSupplier(context.Context, *ReadByNamaSupplierReq) (*ReadByNamaSupplierResp, error)
	ReadSupplier(context.Context, *google_protobuf.Empty) (*ReadSupplierResp, error)
	UpdateSupplier(context.Context, *UpdateSupplierReq) (*google_protobuf.Empty, error)
	ReadByKeterangan(context.Context, *ReadByKeteranganReq) (*ReadByKeteranganResp, error)
}

func RegisterSupplierServiceServer(s *grpc1.Server, srv SupplierServiceServer) {
	s.RegisterService(&_SupplierService_serviceDesc, srv)
}

func _SupplierService_AddSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSupplierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).AddSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/AddSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).AddSupplier(ctx, req.(*AddSupplierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_ReadByNamaSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadByNamaSupplierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).ReadByNamaSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/ReadByNamaSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).ReadByNamaSupplier(ctx, req.(*ReadByNamaSupplierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_ReadSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).ReadSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/ReadSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).ReadSupplier(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_UpdateSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupplierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).UpdateSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/UpdateSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).UpdateSupplier(ctx, req.(*UpdateSupplierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_ReadByKeterangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadByKeteranganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).ReadByKeterangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/ReadByKeterangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).ReadByKeterangan(ctx, req.(*ReadByKeteranganReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupplierService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddSupplier",
			Handler:    _SupplierService_AddSupplier_Handler,
		},
		{
			MethodName: "ReadByNamaSupplier",
			Handler:    _SupplierService_ReadByNamaSupplier_Handler,
		},
		{
			MethodName: "ReadSupplier",
			Handler:    _SupplierService_ReadSupplier_Handler,
		},
		{
			MethodName: "UpdateSupplier",
			Handler:    _SupplierService_UpdateSupplier_Handler,
		},
		{
			MethodName: "ReadByKeterangan",
			Handler:    _SupplierService_ReadByKeterangan_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "supplier.proto",
}

func init() { proto.RegisterFile("supplier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x26, 0x71, 0xd2, 0x4e, 0x4a, 0x80, 0x21, 0x84, 0xc5, 0xad, 0x50, 0xb4, 0xa7, 0x9c,
	0x5c, 0xa9, 0x1c, 0x11, 0x08, 0x02, 0x1c, 0x50, 0x25, 0x24, 0x5c, 0x71, 0xe1, 0xb6, 0x8d, 0x97,
	0x28, 0x62, 0x63, 0x6f, 0xed, 0x35, 0x52, 0xfe, 0x26, 0x7f, 0x85, 0x43, 0x25, 0x4e, 0x68, 0xd7,
	0x71, 0xba, 0xc6, 0xb1, 0xcb, 0x15, 0x8e, 0xf3, 0xe6, 0x63, 0x67, 0xde, 0x9b, 0x1d, 0x18, 0x65,
	0xb9, 0x52, 0x72, 0x25, 0xd2, 0x40, 0xa5, 0x89, 0x4e, 0xb0, 0xb7, 0x4c, 0xd5, 0xc2, 0x3f, 0x59,
	0x26, 0xc9, 0x52, 0x8a, 0x33, 0x8b, 0x5d, 0xe5, 0x5f, 0xcf, 0xc4, 0x5a, 0xe9, 0x4d, 0x11, 0xc2,
	0x7e, 0x10, 0x18, 0xbd, 0x89, 0xa2, 0xcb, 0x6d, 0x62, 0x28, 0xae, 0x91, 0xc1, 0x71, 0xcc, 0xd7,
	0xbc, 0x84, 0x28, 0x99, 0x92, 0xd9, 0x51, 0x58, 0xc1, 0x70, 0x02, 0x7d, 0x2e, 0xf9, 0x9a, 0x6b,
	0xda, 0xb1, 0xde, 0xad, 0x85, 0x14, 0x06, 0x5a, 0x48, 0xa1, 0x92, 0x98, 0x76, 0xad, 0xa3, 0x34,
	0x71, 0x0c, 0x9e, 0x58, 0xf3, 0x95, 0xa4, 0x3d, 0x8b, 0x17, 0x86, 0xa9, 0x93, 0x69, 0xae, 0xf3,
	0x8c, 0x7a, 0x53, 0x32, 0xf3, 0xc2, 0xad, 0x85, 0xa7, 0x70, 0xb4, 0x48, 0x05, 0xd7, 0x22, 0x9a,
	0x6f, 0x68, 0xdf, 0x66, 0xdc, 0x02, 0xf8, 0x0c, 0xe0, 0x9b, 0xd0, 0x22, 0xe5, 0xf1, 0x92, 0xc7,
	0x74, 0x60, 0xdd, 0x0e, 0xc2, 0x5e, 0xc0, 0xe3, 0x50, 0xf0, 0x68, 0xbe, 0xf9, 0xe8, 0xf4, 0xfc,
	0x97, 0xa3, 0xb1, 0x5f, 0x04, 0x26, 0xfb, 0xb2, 0x33, 0x65, 0xde, 0xfd, 0xf0, 0xae, 0x92, 0xec,
	0x85, 0x0e, 0x52, 0x2b, 0xdf, 0x69, 0x65, 0xae, 0xdb, 0xc4, 0x5c, 0xaf, 0x81, 0x39, 0x6f, 0x3f,
	0x73, 0xfd, 0x66, 0xe6, 0x06, 0xed, 0xcc, 0x1d, 0xd6, 0x98, 0xbb, 0x21, 0xf0, 0xa8, 0x18, 0xfe,
	0x62, 0x07, 0x1a, 0xe2, 0xfe, 0xff, 0xc9, 0xbf, 0xc0, 0xb8, 0x3e, 0x78, 0xa6, 0x70, 0x0e, 0xf7,
	0xb8, 0x94, 0xb7, 0x20, 0x25, 0xd3, 0xee, 0x6c, 0x78, 0x7e, 0x1a, 0x98, 0xbf, 0x15, 0xec, 0x5f,
	0x94, 0xb0, 0x9a, 0xc2, 0x42, 0x78, 0x60, 0x02, 0x2b, 0xbb, 0xf4, 0x0a, 0x86, 0x5c, 0x4a, 0x87,
	0xd2, 0xbb, 0xab, 0xba, 0x09, 0xec, 0x27, 0x81, 0x87, 0x9f, 0x55, 0xc4, 0xb5, 0x70, 0x17, 0xfc,
	0x5f, 0xd0, 0xc9, 0x87, 0xc3, 0xdc, 0x36, 0xbe, 0x93, 0x69, 0x67, 0xdf, 0xa5, 0xd2, 0xf9, 0x4d,
	0x07, 0xee, 0x97, 0x8d, 0x5e, 0x8a, 0xf4, 0xfb, 0x6a, 0x21, 0xf0, 0x25, 0x0c, 0x9d, 0x0b, 0x86,
	0xe3, 0x82, 0xc3, 0xea, 0x51, 0xf3, 0x27, 0x41, 0x71, 0x05, 0x83, 0xf2, 0x0a, 0x06, 0xef, 0xcd,
	0x15, 0x64, 0x07, 0xf8, 0x09, 0xb0, 0xce, 0x37, 0x9e, 0x34, 0x2b, 0x71, 0xed, 0xb7, 0xca, 0xc4,
	0x0e, 0xf0, 0x35, 0x1c, 0xbb, 0x7a, 0x63, 0xc3, 0xe3, 0xa6, 0xa9, 0xb2, 0xce, 0x1f, 0x15, 0xde,
	0xc2, 0xa8, 0x2a, 0x2e, 0x3e, 0x29, 0x62, 0x6b, 0x92, 0xb7, 0x4c, 0x76, 0x51, 0xac, 0x9d, 0xbb,
	0xd2, 0xf8, 0xd4, 0x6d, 0xbd, 0xf2, 0xc7, 0x7d, 0xbf, 0xc9, 0x65, 0x3a, 0xba, 0xea, 0xdb, 0xf2,
	0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x79, 0x79, 0xef, 0x03, 0x64, 0x06, 0x00, 0x00,
}
