// Code generated by goa v3.7.6, DO NOT EDIT.
//
// poke gRPC client types
//
// Command:
// $ goa gen poke/design

package client

import (
	pokepb "poke/gen/grpc/poke/pb"
	poke "poke/gen/poke"
)

// NewProtoPokemonRequest builds the gRPC request type from the payload of the
// "pokemon" endpoint of the "poke" service.
func NewProtoPokemonRequest(payload *poke.PokemonPayload) *pokepb.PokemonRequest {
	message := &pokepb.PokemonRequest{
		Name: payload.Name,
	}
	return message
}

// NewPokemonResult builds the result type of the "pokemon" endpoint of the
// "poke" service from the gRPC response type.
func NewPokemonResult(message *pokepb.PokemonResponse) string {
	result := message.Field
	return result
}
