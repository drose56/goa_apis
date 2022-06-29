package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("add", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
    Server("add", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
            URI("grpc://localhost:8080")
        })
    })
})

var _ = Service("add", func() {
	Description("The calc service performs operations on numbers.")

	Method("addnums", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/addnums/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
