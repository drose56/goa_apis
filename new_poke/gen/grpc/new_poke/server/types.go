// Code generated by goa v3.7.6, DO NOT EDIT.
//
// new_poke gRPC server types
//
// Command:
// $ goa gen new_poke/design

package server

import (
	new_pokepb "new_poke/gen/grpc/new_poke/pb"
	newpoke "new_poke/gen/new_poke"
)

// NewPokemonPayload builds the payload of the "pokemon" endpoint of the
// "new_poke" service from the gRPC request type.
func NewPokemonPayload(message *new_pokepb.PokemonRequest) *newpoke.PokemonPayload {
	v := &newpoke.PokemonPayload{
		Name: message.Name,
	}
	return v
}

// NewProtoPokemonResponse builds the gRPC response type from the result of the
// "pokemon" endpoint of the "new_poke" service.
func NewProtoPokemonResponse(result string) *new_pokepb.PokemonResponse {
	message := &new_pokepb.PokemonResponse{}
	message.Field = result
	return message
}