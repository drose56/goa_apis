// Code generated by goa v3.7.6, DO NOT EDIT.
//
// add gRPC client types
//
// Command:
// $ goa gen add/design

package client

import (
	add "add/gen/add"
	addpb "add/gen/grpc/add/pb"
)

// NewProtoAddnumsRequest builds the gRPC request type from the payload of the
// "addnums" endpoint of the "add" service.
func NewProtoAddnumsRequest(payload *add.AddnumsPayload) *addpb.AddnumsRequest {
	message := &addpb.AddnumsRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewAddnumsResult builds the result type of the "addnums" endpoint of the
// "add" service from the gRPC response type.
func NewAddnumsResult(message *addpb.AddnumsResponse) int {
	result := int(message.Field)
	return result
}
