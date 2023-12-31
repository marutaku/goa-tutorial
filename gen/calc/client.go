// Code generated by goa v3.14.1, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	MultiplyEndpoint    goa.Endpoint
	AdditionEndpoint    goa.Endpoint
	DivisionEndpoint    goa.Endpoint
	SubtractionEndpoint goa.Endpoint
	SquareRootEndpoint  goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(multiply, addition, division, subtraction, squareRoot goa.Endpoint) *Client {
	return &Client{
		MultiplyEndpoint:    multiply,
		AdditionEndpoint:    addition,
		DivisionEndpoint:    division,
		SubtractionEndpoint: subtraction,
		SquareRootEndpoint:  squareRoot,
	}
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res *MultiplyResult, err error) {
	var ires any
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MultiplyResult), nil
}

// Addition calls the "addition" endpoint of the "calc" service.
func (c *Client) Addition(ctx context.Context, p *AdditionPayload) (res *AdditionResult, err error) {
	var ires any
	ires, err = c.AdditionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AdditionResult), nil
}

// Division calls the "division" endpoint of the "calc" service.
func (c *Client) Division(ctx context.Context, p *DivisionPayload) (res *DivisionResult, err error) {
	var ires any
	ires, err = c.DivisionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DivisionResult), nil
}

// Subtraction calls the "subtraction" endpoint of the "calc" service.
func (c *Client) Subtraction(ctx context.Context, p *SubtractionPayload) (res *SubtractionResult, err error) {
	var ires any
	ires, err = c.SubtractionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SubtractionResult), nil
}

// SquareRoot calls the "SquareRoot" endpoint of the "calc" service.
func (c *Client) SquareRoot(ctx context.Context, p *SquareRootPayload) (res *SquareRootResult, err error) {
	var ires any
	ires, err = c.SquareRootEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SquareRootResult), nil
}
