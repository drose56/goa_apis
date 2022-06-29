// Code generated by goa v3.7.6, DO NOT EDIT.
//
// new_poke gRPC client CLI support package
//
// Command:
// $ goa gen new_poke/design

package client

import (
	"encoding/json"
	"fmt"
	new_pokepb "new_poke/gen/grpc/new_poke/pb"
	newpoke "new_poke/gen/new_poke"
)

// BuildPokemonPayload builds the payload for the new_poke pokemon endpoint
// from CLI flags.
func BuildPokemonPayload(newPokePokemonMessage string) (*newpoke.PokemonPayload, error) {
	var err error
	var message new_pokepb.PokemonRequest
	{
		if newPokePokemonMessage != "" {
			err = json.Unmarshal([]byte(newPokePokemonMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Quae quam exercitationem consectetur provident ut illo.\"\n   }'")
			}
		}
	}
	v := &newpoke.PokemonPayload{
		Name: message.Name,
	}

	return v, nil
}
