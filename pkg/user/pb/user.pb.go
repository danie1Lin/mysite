// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string               `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type LoginRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Pswd                 string   `protobuf:"bytes,2,opt,name=pswd,proto3" json:"pswd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPswd() string {
	if m != nil {
		return m.Pswd
	}
	return ""
}

type LoginResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LoginResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SignUpRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Pswd                 string   `protobuf:"bytes,2,opt,name=pswd,proto3" json:"pswd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignUpRequest) GetPswd() string {
	if m != nil {
		return m.Pswd
	}
	return ""
}

type SingUpRespond struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingUpRespond) Reset()         { *m = SingUpRespond{} }
func (m *SingUpRespond) String() string { return proto.CompactTextString(m) }
func (*SingUpRespond) ProtoMessage()    {}
func (*SingUpRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *SingUpRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingUpRespond.Unmarshal(m, b)
}
func (m *SingUpRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingUpRespond.Marshal(b, m, deterministic)
}
func (m *SingUpRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingUpRespond.Merge(m, src)
}
func (m *SingUpRespond) XXX_Size() int {
	return xxx_messageInfo_SingUpRespond.Size(m)
}
func (m *SingUpRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_SingUpRespond.DiscardUnknown(m)
}

var xxx_messageInfo_SingUpRespond proto.InternalMessageInfo

func (m *SingUpRespond) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SingUpRespond) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*SignUpRequest)(nil), "pb.SignUpRequest")
	proto.RegisterType((*SingUpRespond)(nil), "pb.SingUpRespond")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xde, 0xb4, 0xb4, 0xd3, 0xb7, 0xa2, 0x8b, 0x87, 0x12, 0x04, 0x4b, 0x4e, 0x3d,
	0xc8, 0x16, 0xea, 0xa9, 0x17, 0xa5, 0x77, 0xf1, 0x90, 0x5a, 0x3c, 0x4a, 0xd2, 0x1d, 0xc3, 0xa2,
	0xc9, 0xae, 0xbb, 0x1b, 0xfd, 0x4a, 0x7e, 0x4c, 0xd9, 0x3f, 0x29, 0xa9, 0x08, 0x82, 0xb7, 0x9d,
	0xc9, 0x3c, 0xf3, 0xfc, 0x9e, 0x09, 0x40, 0xab, 0x51, 0x51, 0xa9, 0x84, 0x11, 0x24, 0x96, 0x65,
	0x7a, 0x59, 0x09, 0x51, 0xbd, 0xe2, 0xd2, 0x75, 0xca, 0xf6, 0x79, 0x69, 0x78, 0x8d, 0xda, 0x14,
	0xb5, 0xf4, 0x43, 0xd9, 0x67, 0x04, 0xc9, 0x4e, 0xa3, 0x22, 0x27, 0x10, 0x73, 0x36, 0x8b, 0xe6,
	0xd1, 0x62, 0x9c, 0xc7, 0x9c, 0x91, 0x14, 0x46, 0x76, 0xd7, 0x7d, 0x51, 0xe3, 0x2c, 0x76, 0xdd,
	0x43, 0x4d, 0xd6, 0x00, 0x7b, 0x85, 0x85, 0x41, 0xf6, 0x54, 0x98, 0xd9, 0xbf, 0x79, 0xb4, 0x98,
	0xac, 0x52, 0xea, 0xad, 0x68, 0x67, 0x45, 0x1f, 0x3a, 0xab, 0x7c, 0x1c, 0xa6, 0x37, 0xc6, 0x4a,
	0x5b, 0xc9, 0x3a, 0x69, 0xf2, 0xbb, 0x34, 0x4c, 0x6f, 0x4c, 0x76, 0x03, 0xff, 0xef, 0x44, 0xc5,
	0x9b, 0x1c, 0xdf, 0x5a, 0xd4, 0xe6, 0x88, 0x30, 0xfa, 0x46, 0x48, 0x20, 0x91, 0xfa, 0x83, 0x05,
	0x72, 0xf7, 0xce, 0x1e, 0x61, 0x1a, 0xf4, 0x5a, 0x8a, 0x46, 0x23, 0xb9, 0xf0, 0xd1, 0x9d, 0x78,
	0xb2, 0x1a, 0x51, 0x59, 0x52, 0x5b, 0xe7, 0xfe, 0x20, 0x04, 0x92, 0xbd, 0x60, 0x5d, 0x78, 0xf7,
	0x26, 0xe7, 0x30, 0x40, 0xa5, 0x84, 0x72, 0x99, 0xc7, 0xb9, 0x2f, 0xb2, 0x5b, 0x98, 0x6e, 0x79,
	0xd5, 0xec, 0xe4, 0x5f, 0xc9, 0xd6, 0x76, 0x41, 0x53, 0xd9, 0x05, 0x16, 0x8d, 0x1d, 0xbc, 0xa3,
	0x9f, 0xbc, 0xe3, 0x9e, 0xf7, 0xea, 0x05, 0x26, 0x96, 0x76, 0x8b, 0xea, 0x9d, 0xef, 0x91, 0x5c,
	0xc1, 0xc0, 0x65, 0x24, 0xa7, 0x36, 0x4d, 0xff, 0x5c, 0xe9, 0x59, 0xaf, 0x13, 0x0e, 0x40, 0x61,
	0xe8, 0xc1, 0x89, 0xfb, 0x78, 0x14, 0x22, 0x0d, 0xad, 0x1e, 0x56, 0x39, 0x74, 0x3f, 0xe8, 0xfa,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x10, 0xad, 0xd8, 0x51, 0x66, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SingUpRespond, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SingUpRespond, error) {
	out := new(SingUpRespond)
	err := c.cc.Invoke(ctx, "/pb.UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SignUp(context.Context, *SignUpRequest) (*SingUpRespond, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) SignUp(ctx context.Context, req *SignUpRequest) (*SingUpRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}