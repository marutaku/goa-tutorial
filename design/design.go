package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calulator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")
	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(func() {
			Attribute("result", Int)
			Required("result")
		})

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

	})
	Method("addition", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(func() {
			Attribute("result", Int)
			Required("result")
		})

		HTTP(func() {
			GET("/addition/{a}/{b}")
		})
	})
	Method("division", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(func() {
			Attribute("result", Float64)
			Required("result")
		})

		HTTP(func() {
			GET("/division/{a}/{b}")
		})
	})
	Method("subtraction", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(func() {
			Attribute("result", Int)
			Required("result")
		})

		HTTP(func() {
			GET("/subtraction/{a}/{b}")
		})
	})

	Method("SquareRoot", func() {
		Payload(func() {
			Field(1, "a", Int, "number")
			Required("a")
		})

		Result(func() {
			Attribute("result", Float64)
			Required("result")
		})

		HTTP(func() {
			POST("/square_root")
			Body(func() {
				Attribute("a")
			})
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
