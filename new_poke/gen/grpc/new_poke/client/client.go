// Code generated by goa v3.7.6, DO NOT EDIT.
//
// new_poke gRPC client
//
// Command:
// $ goa gen new_poke/design

package client

import (
	"context"
	new_pokepb "new_poke/gen/grpc/new_poke/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli new_pokepb.NewPokeClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the new_poke service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: new_pokepb.NewNewPokeClient(cc),
		opts:    opts,
	}
}

// Pokemon calls the "Pokemon" function in new_pokepb.NewPokeClient interface.
func (c *Client) Pokemon() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildPokemonFunc(c.grpccli, c.opts...),
			EncodePokemonRequest,
			DecodePokemonResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
