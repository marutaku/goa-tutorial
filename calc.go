package calcapi

import (
	calc "calc/gen/calc"
	"context"
	"log"
	"math"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res *calc.MultiplyResult, err error) {
	res = &calc.MultiplyResult{}
	s.logger.Print("calc.multiply")
	res.Result = p.A * p.B
	return res, nil
}

// Addition implements addition.
func (s *calcsrvc) Addition(ctx context.Context, p *calc.AdditionPayload) (res *calc.AdditionResult, err error) {
	res = &calc.AdditionResult{}
	s.logger.Print("calc.addition")
	res.Result = p.A + p.B
	return res, nil
}

// Division implements division.
func (s *calcsrvc) Division(ctx context.Context, p *calc.DivisionPayload) (res *calc.DivisionResult, err error) {
	res = &calc.DivisionResult{}
	res.Result = float64(p.A) / float64(p.B)
	s.logger.Print("calc.division")
	return res, nil
}

// Subtraction implements subtraction.
func (s *calcsrvc) Subtraction(ctx context.Context, p *calc.SubtractionPayload) (res *calc.SubtractionResult, err error) {
	res = &calc.SubtractionResult{}
	res.Result = p.A - p.B
	s.logger.Print("calc.subtraction")
	return res, nil
}

func (s *calcsrvc) SquareRoot(ctx context.Context, p *calc.SquareRootPayload) (res *calc.SquareRootResult, err error) {
	res = &calc.SquareRootResult{}
	res.Result = math.Sqrt(float64(p.A))
	s.logger.Print("calc.square_root")
	return res, nil
}
