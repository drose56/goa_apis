package pokeapi

import (
	"context"
	"log"
	poke "poke/gen/poke"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
	"encoding/json"
)

// poke service example implementation.
// The example methods log the requests and return zero values.
type pokesrvc struct {
	logger *log.Logger
}

// NewPoke returns the poke service implementation.
func NewPoke(logger *log.Logger) poke.Service {
	return &pokesrvc{logger}
}

type Response struct {
    Pokemon []Pokemon `json:"results"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int `json:"entry_number"`
    Name string `json:"name"`
}



// Pokemon implements pokemon.
func (s *pokesrvc) Pokemon(ctx context.Context, p *poke.PokemonPayload) (res string, err error) {
    response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10000")

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
	
	//fmt.Println(responseObject.Name)
	//fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < 1000; i++ {
		//fmt.Print(responseObject.Pokemon[i].Name)
		if string(p.Name[0]) == string((responseObject.Pokemon[i].Name)[0]){
			return responseObject.Pokemon[i].Name, nil
	}
	}

	return responseObject.Pokemon[1].Name, nil
    //return string(responseData), nil
	//return string(p.Name[0]), nil
}
