package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("new_poke", func() {
	Title("pokemon letters")
	Description("give a name and get back pokemon with same first letter")
    Server("new_poke", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
            URI("grpc://localhost:8080")
        })
    })
})

var _ = Service("new_poke", func() {
	Description("return pokemon with same first letter as name")

	Method("pokemon", func() {
		Payload(func() {
			Field(1, "name", String, "first name")
			
			Required("name")
		})

		Result(String)

		HTTP(func() {
			GET("/pokemon/{name}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

