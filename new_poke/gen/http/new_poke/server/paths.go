// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the new_poke service.
//
// Command:
// $ goa gen new_poke/design

package server

import (
	"fmt"
)

// PokemonNewPokePath returns the URL path to the new_poke service pokemon HTTP endpoint.
func PokemonNewPokePath(name string) string {
	return fmt.Sprintf("/pokemon/%v", name)
}
