package newpokeapi

import (
	"context"
	"log"
	newpoke "new_poke/gen/new_poke"
	"fmt"
	"io/ioutil"
    "net/http"
    "os"
	"encoding/json"
	"strconv"
)

// new_poke service example implementation.
// The example methods log the requests and return zero values.
type newPokesrvc struct {
	logger *log.Logger
}

// NewNewPoke returns the new_poke service implementation.
func NewNewPoke(logger *log.Logger) newpoke.Service {
	return &newPokesrvc{logger}
}

type Response struct {
	Name string `json:"name"`
}


// Pokemon implements pokemon.
func (s *newPokesrvc) Pokemon(ctx context.Context, p *newpoke.PokemonPayload) (res string, err error) {
	var i int
	i = 1
	for true {
		x := strconv.Itoa(i)

		response, err := http.Get("https://pokeapi.co/api/v2/pokemon-form/" +  x)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)		

		if string(p.Name[0]) == string((responseObject.Name)[0]){
			return responseObject.Name, nil
		}
		//return responseObject.Name, nil
		i += 1
	}
	return "never gonna get here", nil
}
