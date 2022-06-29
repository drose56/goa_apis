// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goadesign_goagen_new_poke.proto

package new_pokepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NewPokeClient is the client API for NewPoke service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewPokeClient interface {
	// Pokemon implements pokemon.
	Pokemon(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error)
}

type newPokeClient struct {
	cc grpc.ClientConnInterface
}

func NewNewPokeClient(cc grpc.ClientConnInterface) NewPokeClient {
	return &newPokeClient{cc}
}

func (c *newPokeClient) Pokemon(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error) {
	out := new(PokemonResponse)
	err := c.cc.Invoke(ctx, "/new_poke.NewPoke/Pokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewPokeServer is the server API for NewPoke service.
// All implementations must embed UnimplementedNewPokeServer
// for forward compatibility
type NewPokeServer interface {
	// Pokemon implements pokemon.
	Pokemon(context.Context, *PokemonRequest) (*PokemonResponse, error)
	mustEmbedUnimplementedNewPokeServer()
}

// UnimplementedNewPokeServer must be embedded to have forward compatible implementations.
type UnimplementedNewPokeServer struct {
}

func (UnimplementedNewPokeServer) Pokemon(context.Context, *PokemonRequest) (*PokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pokemon not implemented")
}
func (UnimplementedNewPokeServer) mustEmbedUnimplementedNewPokeServer() {}

// UnsafeNewPokeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewPokeServer will
// result in compilation errors.
type UnsafeNewPokeServer interface {
	mustEmbedUnimplementedNewPokeServer()
}

func RegisterNewPokeServer(s grpc.ServiceRegistrar, srv NewPokeServer) {
	s.RegisterService(&NewPoke_ServiceDesc, srv)
}

func _NewPoke_Pokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewPokeServer).Pokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/new_poke.NewPoke/Pokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewPokeServer).Pokemon(ctx, req.(*PokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewPoke_ServiceDesc is the grpc.ServiceDesc for NewPoke service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewPoke_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "new_poke.NewPoke",
	HandlerType: (*NewPokeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pokemon",
			Handler:    _NewPoke_Pokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_new_poke.proto",
}
