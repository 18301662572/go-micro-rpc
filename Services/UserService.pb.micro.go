// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: UserService.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	UserReg(ctx context.Context, in *UserModel, opts ...client.CallOption) (*RegReponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Services"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserReg(ctx context.Context, in *UserModel, opts ...client.CallOption) (*RegReponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserReg", in)
	out := new(RegReponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	UserReg(context.Context, *UserModel, *RegReponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		UserReg(ctx context.Context, in *UserModel, out *RegReponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) UserReg(ctx context.Context, in *UserModel, out *RegReponse) error {
	return h.UserServiceHandler.UserReg(ctx, in, out)
}
