// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: githubFetcher/proto/githubFetcher.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GithubFetcher_GetRepositoriesEvaluation_FullMethodName = "/proto.GithubFetcher/GetRepositoriesEvaluation"
)

// GithubFetcherClient is the client API for GithubFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubFetcherClient interface {
	GetRepositoriesEvaluation(ctx context.Context, in *GetRepositoriesEvaluationRequest, opts ...grpc.CallOption) (*GetRepositoriesEvaluationResponse, error)
}

type githubFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubFetcherClient(cc grpc.ClientConnInterface) GithubFetcherClient {
	return &githubFetcherClient{cc}
}

func (c *githubFetcherClient) GetRepositoriesEvaluation(ctx context.Context, in *GetRepositoriesEvaluationRequest, opts ...grpc.CallOption) (*GetRepositoriesEvaluationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRepositoriesEvaluationResponse)
	err := c.cc.Invoke(ctx, GithubFetcher_GetRepositoriesEvaluation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubFetcherServer is the server API for GithubFetcher service.
// All implementations must embed UnimplementedGithubFetcherServer
// for forward compatibility.
type GithubFetcherServer interface {
	GetRepositoriesEvaluation(context.Context, *GetRepositoriesEvaluationRequest) (*GetRepositoriesEvaluationResponse, error)
	mustEmbedUnimplementedGithubFetcherServer()
}

// UnimplementedGithubFetcherServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGithubFetcherServer struct{}

func (UnimplementedGithubFetcherServer) GetRepositoriesEvaluation(context.Context, *GetRepositoriesEvaluationRequest) (*GetRepositoriesEvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoriesEvaluation not implemented")
}
func (UnimplementedGithubFetcherServer) mustEmbedUnimplementedGithubFetcherServer() {}
func (UnimplementedGithubFetcherServer) testEmbeddedByValue()                       {}

// UnsafeGithubFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubFetcherServer will
// result in compilation errors.
type UnsafeGithubFetcherServer interface {
	mustEmbedUnimplementedGithubFetcherServer()
}

func RegisterGithubFetcherServer(s grpc.ServiceRegistrar, srv GithubFetcherServer) {
	// If the following call pancis, it indicates UnimplementedGithubFetcherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GithubFetcher_ServiceDesc, srv)
}

func _GithubFetcher_GetRepositoriesEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoriesEvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubFetcherServer).GetRepositoriesEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubFetcher_GetRepositoriesEvaluation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubFetcherServer).GetRepositoriesEvaluation(ctx, req.(*GetRepositoriesEvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubFetcher_ServiceDesc is the grpc.ServiceDesc for GithubFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GithubFetcher",
	HandlerType: (*GithubFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRepositoriesEvaluation",
			Handler:    _GithubFetcher_GetRepositoriesEvaluation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "githubFetcher/proto/githubFetcher.proto",
}
