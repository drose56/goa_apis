package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("data", func() {
        Title("clickhouse data given time")
        Description("Give two datetime strings and get all data in between from clickhouse")
        Server("data", func() {
                Host("localhost", func() { URI("http://localhost:8088") })
        })
})

// Service describes a service
var _ = Service("data", func() {
        Description("grabs data given datetimes")
        // Method describes a service method (endpoint)
        Method("pull_data", func() {
                // Payload describes the method payload
                // Here the payload is an object that consists of two fields
                Payload(func() {
                        // Attribute describes an object field
                        Attribute("a", Int, "Left operand")
                        Attribute("b", Int, "Right operand")
                        // Both attributes must be provided when invoking "multiply"
                        Required("a", "b")
                })
                // Result describes the method result
                // Here the result is a simple integer value
                Result(Int)
                // HTTP describes the HTTP transport mapping
                HTTP(func() {
                        // Requests to the service consist of HTTP GET requests
                        // The payload fields are encoded as path parameters
                        GET("/multiply/{a}/{b}")
                        // Responses use a "200 OK" HTTP status
                        // The result is encoded in the response body
                        Response(StatusOK)
                })
        })
})
