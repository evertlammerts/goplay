// Code generated by protoc-gen-go.
// source: integrator.proto
// DO NOT EDIT!

/*
Package integrator is a generated protocol buffer package.

It is generated from these files:
	integrator.proto

It has these top-level messages:
	Query
	IntegralEstimate
	STPolynomial
	IntegralConfig
*/
package integrator

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
	Poly *STPolynomial   `protobuf:"bytes,1,opt,name=poly" json:"poly,omitempty"`
	Conf *IntegralConfig `protobuf:"bytes,2,opt,name=conf" json:"conf,omitempty"`
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

func (m *Query) GetConf() *IntegralConfig {
	if m != nil {
		return m.Conf
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

// Single term polynomials are defined by their constants. I.e., -5x^3-43 is
// actually -5x^3+0x^2+0x^1-43X^0, which this message expresses as {Constants:
// []int32{-5, 0, 0, -43}}
type STPolynomial struct {
	Constants []int32 `protobuf:"zigzag32,1,rep,packed,name=constants" json:"constants,omitempty"`
}

func (m *STPolynomial) Reset()                    { *m = STPolynomial{} }
func (m *STPolynomial) String() string            { return proto.CompactTextString(m) }
func (*STPolynomial) ProtoMessage()               {}
func (*STPolynomial) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *STPolynomial) GetConstants() []int32 {
	if m != nil {
		return m.Constants
	}
	return nil
}

// Definit integral estimates are defined by the amount of strips we create
// under the curve and the start and end of the interval we measure the area
// of.
type IntegralConfig struct {
	NumStrips int64   `protobuf:"varint,1,opt,name=numStrips" json:"numStrips,omitempty"`
	Start     float64 `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End       float64 `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
}

func (m *IntegralConfig) Reset()                    { *m = IntegralConfig{} }
func (m *IntegralConfig) String() string            { return proto.CompactTextString(m) }
func (*IntegralConfig) ProtoMessage()               {}
func (*IntegralConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntegralConfig) GetNumStrips() int64 {
	if m != nil {
		return m.NumStrips
	}
	return 0
}

func (m *IntegralConfig) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *IntegralConfig) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Query)(nil), "integrator.Query")
	proto.RegisterType((*IntegralEstimate)(nil), "integrator.IntegralEstimate")
	proto.RegisterType((*STPolynomial)(nil), "integrator.STPolynomial")
	proto.RegisterType((*IntegralConfig)(nil), "integrator.IntegralConfig")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Integrator service

type IntegratorClient interface {
	// Estimate the integral of the given single term polynomial at the given
	// interval.
	//
	// Accepts a message that defines the polynomial and the integral config.
	// Returns the estimated value of the area under the curve.
	SingleTermPolynomialIntegral(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegralEstimate, error)
}

type integratorClient struct {
	cc *grpc.ClientConn
}

func NewIntegratorClient(cc *grpc.ClientConn) IntegratorClient {
	return &integratorClient{cc}
}

func (c *integratorClient) SingleTermPolynomialIntegral(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegralEstimate, error) {
	out := new(IntegralEstimate)
	err := grpc.Invoke(ctx, "/integrator.Integrator/SingleTermPolynomialIntegral", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Integrator service

type IntegratorServer interface {
	// Estimate the integral of the given single term polynomial at the given
	// interval.
	//
	// Accepts a message that defines the polynomial and the integral config.
	// Returns the estimated value of the area under the curve.
	SingleTermPolynomialIntegral(context.Context, *Query) (*IntegralEstimate, error)
}

func RegisterIntegratorServer(s *grpc.Server, srv IntegratorServer) {
	s.RegisterService(&_Integrator_serviceDesc, srv)
}

func _Integrator_SingleTermPolynomialIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegratorServer).SingleTermPolynomialIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integrator.Integrator/SingleTermPolynomialIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegratorServer).SingleTermPolynomialIntegral(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Integrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "integrator.Integrator",
	HandlerType: (*IntegratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SingleTermPolynomialIntegral",
			Handler:    _Integrator_SingleTermPolynomialIntegral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrator.proto",
}

func init() { proto.RegisterFile("integrator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0xdb, 0x16, 0x1c, 0x45, 0xda, 0x20, 0x12, 0xca, 0x1e, 0x4a, 0x4e, 0x45, 0xca,
	0x1e, 0xea, 0x4f, 0x10, 0x0f, 0xbd, 0xd9, 0x6c, 0xf1, 0x2a, 0xb1, 0xa6, 0x4b, 0x20, 0x9b, 0x2c,
	0xc9, 0xec, 0x61, 0xff, 0xbd, 0x74, 0xea, 0xba, 0x2b, 0x78, 0x9b, 0x79, 0x7c, 0xc9, 0x9b, 0x79,
	0x03, 0x73, 0xeb, 0xd1, 0x54, 0x51, 0x63, 0x88, 0x45, 0x13, 0x03, 0x06, 0x0e, 0x83, 0x22, 0x0d,
	0x4c, 0xf7, 0xad, 0x89, 0x1d, 0xdf, 0xc0, 0xa4, 0x09, 0xae, 0x13, 0x6c, 0xc5, 0xd6, 0xb7, 0x5b,
	0x51, 0x8c, 0x5e, 0x95, 0x87, 0xb7, 0xe0, 0x3a, 0x1f, 0x6a, 0xab, 0x9d, 0x22, 0x8a, 0x17, 0x30,
	0x39, 0x06, 0x7f, 0x12, 0xd7, 0x44, 0x2f, 0xc7, 0xf4, 0xee, 0x52, 0xba, 0x97, 0xe0, 0x4f, 0xb6,
	0x52, 0xc4, 0xc9, 0x27, 0x98, 0xf7, 0xfa, 0x6b, 0x42, 0x5b, 0x6b, 0x34, 0xfc, 0x11, 0x66, 0xd1,
	0xa4, 0xd6, 0x21, 0x79, 0x32, 0xf5, 0xd3, 0xc9, 0x0d, 0xdc, 0x8d, 0x1d, 0x79, 0x0e, 0x37, 0xc7,
	0xe0, 0x13, 0x6a, 0x8f, 0x49, 0xb0, 0x55, 0xb6, 0x5e, 0xa8, 0x41, 0x90, 0xef, 0x70, 0xff, 0xd7,
	0xf1, 0xcc, 0xfb, 0xb6, 0x2e, 0x31, 0xda, 0x26, 0xd1, 0xd7, 0x99, 0x1a, 0x04, 0xfe, 0x00, 0xd3,
	0x84, 0x3a, 0x22, 0x8d, 0xce, 0xd4, 0xa5, 0xe1, 0x73, 0xc8, 0x8c, 0xff, 0x12, 0x19, 0x69, 0xe7,
	0x72, 0xfb, 0x01, 0xb0, 0xfb, 0x5d, 0x8a, 0xef, 0x21, 0x2f, 0xad, 0xaf, 0x9c, 0x39, 0x98, 0x58,
	0x0f, 0xb3, 0xf5, 0xce, 0x7c, 0x31, 0x4e, 0x80, 0x02, 0x5d, 0xe6, 0xff, 0x85, 0xd2, 0x2f, 0x2f,
	0xaf, 0x3e, 0x67, 0x74, 0x8c, 0xe7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x80, 0xe3, 0xf7,
	0xa0, 0x01, 0x00, 0x00,
}