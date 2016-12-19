// Code generated by protoc-gen-go.
// source: mathserver.proto
// DO NOT EDIT!

/*
Package mathserver is a generated protocol buffer package.

It is generated from these files:
	mathserver.proto

It has these top-level messages:
	Query
	IntegralEstimate
	STPolynomial
	IntegralDefinition
*/
package mathserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Compound query message
type Query struct {
	Poly *STPolynomial       `protobuf:"bytes,1,opt,name=poly" json:"poly,omitempty"`
	Def  *IntegralDefinition `protobuf:"bytes,2,opt,name=def" json:"def,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Query) GetPoly() *STPolynomial {
	if m != nil {
		return m.Poly
	}
	return nil
}

func (m *Query) GetDef() *IntegralDefinition {
	if m != nil {
		return m.Def
	}
	return nil
}

// Estimated integral return message
type IntegralEstimate struct {
	Result float64 `protobuf:"fixed64,1,opt,name=result" json:"result,omitempty"`
}

func (m *IntegralEstimate) Reset()                    { *m = IntegralEstimate{} }
func (m *IntegralEstimate) String() string            { return proto.CompactTextString(m) }
func (*IntegralEstimate) ProtoMessage()               {}
func (*IntegralEstimate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IntegralEstimate) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

// Single term polynomials are defined by their degree and their constants.
// I.e., -5x^3-43 is {Degree: 3, Constants: []int32{-5, 0, 0, -43}}
type STPolynomial struct {
	Degree    uint32  `protobuf:"varint,1,opt,name=degree" json:"degree,omitempty"`
	Constants []int32 `protobuf:"zigzag32,2,rep,packed,name=constants" json:"constants,omitempty"`
}

func (m *STPolynomial) Reset()                    { *m = STPolynomial{} }
func (m *STPolynomial) String() string            { return proto.CompactTextString(m) }
func (*STPolynomial) ProtoMessage()               {}
func (*STPolynomial) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *STPolynomial) GetDegree() uint32 {
	if m != nil {
		return m.Degree
	}
	return 0
}

func (m *STPolynomial) GetConstants() []int32 {
	if m != nil {
		return m.Constants
	}
	return nil
}

// Definit integral estimates are defined by the amount of strips we create
// under the curve and the start and end of the interval we measure the area
// of.
type IntegralDefinition struct {
	NumStrips int64   `protobuf:"varint,1,opt,name=numStrips" json:"numStrips,omitempty"`
	Start     float64 `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End       float64 `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
}

func (m *IntegralDefinition) Reset()                    { *m = IntegralDefinition{} }
func (m *IntegralDefinition) String() string            { return proto.CompactTextString(m) }
func (*IntegralDefinition) ProtoMessage()               {}
func (*IntegralDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntegralDefinition) GetNumStrips() int64 {
	if m != nil {
		return m.NumStrips
	}
	return 0
}

func (m *IntegralDefinition) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *IntegralDefinition) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Query)(nil), "mathserver.Query")
	proto.RegisterType((*IntegralEstimate)(nil), "mathserver.IntegralEstimate")
	proto.RegisterType((*STPolynomial)(nil), "mathserver.STPolynomial")
	proto.RegisterType((*IntegralDefinition)(nil), "mathserver.IntegralDefinition")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MathServer service

type MathServerClient interface {
	// Estimate the integral of the given single term polynomial at the given
	// interval.
	//
	// Accepts a message that defines the polynomial and the integral definition.
	// Returns the estimated value of the area under the curve.
	SingleTermPolynomialIntegral(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegralEstimate, error)
}

type mathServerClient struct {
	cc *grpc.ClientConn
}

func NewMathServerClient(cc *grpc.ClientConn) MathServerClient {
	return &mathServerClient{cc}
}

func (c *mathServerClient) SingleTermPolynomialIntegral(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegralEstimate, error) {
	out := new(IntegralEstimate)
	err := grpc.Invoke(ctx, "/mathserver.MathServer/SingleTermPolynomialIntegral", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MathServer service

type MathServerServer interface {
	// Estimate the integral of the given single term polynomial at the given
	// interval.
	//
	// Accepts a message that defines the polynomial and the integral definition.
	// Returns the estimated value of the area under the curve.
	SingleTermPolynomialIntegral(context.Context, *Query) (*IntegralEstimate, error)
}

func RegisterMathServerServer(s *grpc.Server, srv MathServerServer) {
	s.RegisterService(&_MathServer_serviceDesc, srv)
}

func _MathServer_SingleTermPolynomialIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServerServer).SingleTermPolynomialIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mathserver.MathServer/SingleTermPolynomialIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServerServer).SingleTermPolynomialIntegral(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mathserver.MathServer",
	HandlerType: (*MathServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SingleTermPolynomialIntegral",
			Handler:    _MathServer_SingleTermPolynomialIntegral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mathserver.proto",
}

func init() { proto.RegisterFile("mathserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x63, 0x0b, 0x8e, 0x0a, 0xe9, 0x22, 0x12, 0x24, 0x48, 0xc9, 0xa9, 0x88, 0x14,
	0xa9, 0x5f, 0xa1, 0x1e, 0x3c, 0x08, 0x76, 0xd3, 0x93, 0x17, 0x59, 0xcd, 0x34, 0x5d, 0xd8, 0x3f,
	0x61, 0x77, 0x22, 0xe4, 0xdb, 0x4b, 0xd6, 0x86, 0x04, 0xec, 0x6d, 0xe7, 0xed, 0x6f, 0xf6, 0xcd,
	0xbc, 0x85, 0x44, 0x0b, 0x3a, 0x78, 0x74, 0x3f, 0xe8, 0x56, 0xb5, 0xb3, 0x64, 0x19, 0x0c, 0x4a,
	0x5e, 0xc1, 0x74, 0xdb, 0xa0, 0x6b, 0xd9, 0x23, 0x9c, 0xd7, 0x56, 0xb5, 0x69, 0xb4, 0x88, 0x96,
	0x97, 0xeb, 0x74, 0x35, 0xea, 0x2a, 0x76, 0xef, 0x56, 0xb5, 0xc6, 0x6a, 0x29, 0x14, 0x0f, 0x14,
	0x7b, 0x82, 0xb8, 0xc4, 0x7d, 0x3a, 0x09, 0xf0, 0xfd, 0x18, 0x7e, 0x35, 0x84, 0x95, 0x13, 0x6a,
	0x83, 0x7b, 0x69, 0x24, 0x49, 0x6b, 0x78, 0x87, 0xe6, 0x0f, 0x90, 0xf4, 0x57, 0x2f, 0x9e, 0xa4,
	0x16, 0x84, 0xec, 0x16, 0x66, 0x0e, 0x7d, 0xa3, 0x28, 0xb8, 0x46, 0xfc, 0x58, 0xe5, 0x1b, 0xb8,
	0x1a, 0x7b, 0x76, 0x5c, 0x89, 0x95, 0x43, 0x0c, 0xdc, 0x35, 0x3f, 0x56, 0x2c, 0x83, 0x8b, 0x6f,
	0x6b, 0x3c, 0x09, 0x43, 0x3e, 0x9d, 0x2c, 0xe2, 0xe5, 0x9c, 0x0f, 0x42, 0xfe, 0x01, 0xec, 0xff,
	0x30, 0x5d, 0x8f, 0x69, 0x74, 0x41, 0x4e, 0xd6, 0x3e, 0x3c, 0x17, 0xf3, 0x41, 0x60, 0x37, 0x30,
	0xf5, 0x24, 0x1c, 0x85, 0xcd, 0x22, 0xfe, 0x57, 0xb0, 0x04, 0x62, 0x34, 0x65, 0x1a, 0x07, 0xad,
	0x3b, 0xae, 0x3f, 0x01, 0xde, 0x04, 0x1d, 0x8a, 0xb0, 0x33, 0xdb, 0x42, 0x56, 0x48, 0x53, 0x29,
	0xdc, 0xa1, 0xd3, 0xc3, 0xdc, 0xbd, 0x3b, 0x9b, 0x8f, 0x03, 0x0a, 0x71, 0xdf, 0x65, 0xa7, 0x32,
	0xeb, 0x83, 0xc9, 0xcf, 0xbe, 0x66, 0xe1, 0xab, 0x9e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc4,
	0x0c, 0xd2, 0xb1, 0xbe, 0x01, 0x00, 0x00,
}
