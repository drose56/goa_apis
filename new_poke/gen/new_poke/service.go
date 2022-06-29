// Code generated by goa v3.7.6, DO NOT EDIT.
//
// new_poke service
//
// Command:
// $ goa gen new_poke/design

package newpoke

import (
	"context"
)

// return pokemon with same first letter as name
type Service interface {
	// Pokemon implements pokemon.
	Pokemon(context.Context, *PokemonPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "new_poke"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"pokemon"}

// PokemonPayload is the payload type of the new_poke service pokemon method.
type PokemonPayload struct {
	// first name
	Name string
}