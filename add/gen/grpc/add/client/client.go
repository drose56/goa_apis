// Code generated by goa v3.7.6, DO NOT EDIT.
//
// add gRPC client
//
// Command:
// $ goa gen add/design

package client

import (
	addpb "add/gen/grpc/add/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli addpb.AddClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the add service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: addpb.NewAddClient(cc),
		opts:    opts,
	}
}

// Addnums calls the "Addnums" function in addpb.AddClient interface.
func (c *Client) Addnums() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAddnumsFunc(c.grpccli, c.opts...),
			EncodeAddnumsRequest,
			DecodeAddnumsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
