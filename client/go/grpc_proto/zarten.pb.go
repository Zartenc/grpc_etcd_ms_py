// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zarten.proto

package zarten

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ZartenRequest struct {
	ZhihuName            string   `protobuf:"bytes,1,opt,name=zhihu_name,json=zhihuName,proto3" json:"zhihu_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZartenRequest) Reset()         { *m = ZartenRequest{} }
func (m *ZartenRequest) String() string { return proto.CompactTextString(m) }
func (*ZartenRequest) ProtoMessage()    {}
func (*ZartenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ef5befb305c52e, []int{0}
}
func (m *ZartenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ZartenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ZartenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ZartenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZartenRequest.Merge(m, src)
}
func (m *ZartenRequest) XXX_Size() int {
	return m.Size()
}
func (m *ZartenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ZartenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ZartenRequest proto.InternalMessageInfo

func (m *ZartenRequest) GetZhihuName() string {
	if m != nil {
		return m.ZhihuName
	}
	return ""
}

type ZartenResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Homepage             string   `protobuf:"bytes,2,opt,name=homepage,proto3" json:"homepage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZartenResponse) Reset()         { *m = ZartenResponse{} }
func (m *ZartenResponse) String() string { return proto.CompactTextString(m) }
func (*ZartenResponse) ProtoMessage()    {}
func (*ZartenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ef5befb305c52e, []int{1}
}
func (m *ZartenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ZartenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ZartenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ZartenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZartenResponse.Merge(m, src)
}
func (m *ZartenResponse) XXX_Size() int {
	return m.Size()
}
func (m *ZartenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ZartenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ZartenResponse proto.InternalMessageInfo

func (m *ZartenResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ZartenResponse) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func init() {
	proto.RegisterType((*ZartenRequest)(nil), "ZartenRequest")
	proto.RegisterType((*ZartenResponse)(nil), "ZartenResponse")
}

func init() { proto.RegisterFile("zarten.proto", fileDescriptor_12ef5befb305c52e) }

var fileDescriptor_12ef5befb305c52e = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xa9, 0x4a, 0x2c, 0x2a,
	0x49, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe3, 0xe2, 0x8d, 0x02, 0xf3, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb9, 0xb8, 0xaa, 0x32, 0x32, 0x33, 0x4a, 0xe3,
	0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0xc1, 0x22, 0x7e, 0x89,
	0xb9, 0xa9, 0x4a, 0x0e, 0x5c, 0x7c, 0x30, 0xf5, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42,
	0x5c, 0x2c, 0x48, 0x4a, 0xc1, 0x6c, 0x21, 0x29, 0x2e, 0x8e, 0x8c, 0xfc, 0xdc, 0xd4, 0x82, 0xc4,
	0xf4, 0x54, 0x09, 0x26, 0xb0, 0x38, 0x9c, 0x6f, 0x64, 0xc2, 0xc5, 0x06, 0x31, 0x41, 0x48, 0x8b,
	0x8b, 0xdd, 0x3d, 0xb5, 0xc4, 0x33, 0x2f, 0x2d, 0x5f, 0x88, 0x4f, 0x0f, 0xc5, 0x15, 0x52, 0xfc,
	0x7a, 0xa8, 0xb6, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x3d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x17, 0x08, 0x15, 0xd0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ZartenClient is the client API for Zarten service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZartenClient interface {
	GetInfo(ctx context.Context, in *ZartenRequest, opts ...grpc.CallOption) (*ZartenResponse, error)
}

type zartenClient struct {
	cc *grpc.ClientConn
}

func NewZartenClient(cc *grpc.ClientConn) ZartenClient {
	return &zartenClient{cc}
}

func (c *zartenClient) GetInfo(ctx context.Context, in *ZartenRequest, opts ...grpc.CallOption) (*ZartenResponse, error) {
	out := new(ZartenResponse)
	err := c.cc.Invoke(ctx, "/Zarten/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZartenServer is the server API for Zarten service.
type ZartenServer interface {
	GetInfo(context.Context, *ZartenRequest) (*ZartenResponse, error)
}

// UnimplementedZartenServer can be embedded to have forward compatible implementations.
type UnimplementedZartenServer struct {
}

func (*UnimplementedZartenServer) GetInfo(ctx context.Context, req *ZartenRequest) (*ZartenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterZartenServer(s *grpc.Server, srv ZartenServer) {
	s.RegisterService(&_Zarten_serviceDesc, srv)
}

func _Zarten_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZartenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZartenServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Zarten/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZartenServer).GetInfo(ctx, req.(*ZartenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Zarten_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Zarten",
	HandlerType: (*ZartenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Zarten_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zarten.proto",
}

func (m *ZartenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZartenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ZartenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ZhihuName) > 0 {
		i -= len(m.ZhihuName)
		copy(dAtA[i:], m.ZhihuName)
		i = encodeVarintZarten(dAtA, i, uint64(len(m.ZhihuName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ZartenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZartenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ZartenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Homepage) > 0 {
		i -= len(m.Homepage)
		copy(dAtA[i:], m.Homepage)
		i = encodeVarintZarten(dAtA, i, uint64(len(m.Homepage)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintZarten(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintZarten(dAtA []byte, offset int, v uint64) int {
	offset -= sovZarten(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ZartenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ZhihuName)
	if l > 0 {
		n += 1 + l + sovZarten(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ZartenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovZarten(uint64(l))
	}
	l = len(m.Homepage)
	if l > 0 {
		n += 1 + l + sovZarten(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovZarten(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZarten(x uint64) (n int) {
	return sovZarten(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ZartenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZarten
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ZartenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZartenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZhihuName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZarten
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZarten
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZarten
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZhihuName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZarten(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthZarten
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthZarten
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ZartenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZarten
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ZartenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZartenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZarten
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZarten
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZarten
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Homepage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZarten
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZarten
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZarten
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Homepage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZarten(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthZarten
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthZarten
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipZarten(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZarten
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZarten
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowZarten
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthZarten
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupZarten
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthZarten
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthZarten        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZarten          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupZarten = fmt.Errorf("proto: unexpected end of group")
)
