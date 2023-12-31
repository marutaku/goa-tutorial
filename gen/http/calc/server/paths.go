// Code generated by goa v3.14.1, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen calc/design

package server

import (
	"fmt"
)

// MultiplyCalcPath returns the URL path to the calc service multiply HTTP endpoint.
func MultiplyCalcPath(a int, b int) string {
	return fmt.Sprintf("/multiply/%v/%v", a, b)
}

// AdditionCalcPath returns the URL path to the calc service addition HTTP endpoint.
func AdditionCalcPath(a int, b int) string {
	return fmt.Sprintf("/addition/%v/%v", a, b)
}

// DivisionCalcPath returns the URL path to the calc service division HTTP endpoint.
func DivisionCalcPath(a int, b int) string {
	return fmt.Sprintf("/division/%v/%v", a, b)
}

// SubtractionCalcPath returns the URL path to the calc service subtraction HTTP endpoint.
func SubtractionCalcPath(a int, b int) string {
	return fmt.Sprintf("/subtraction/%v/%v", a, b)
}

// SquareRootCalcPath returns the URL path to the calc service SquareRoot HTTP endpoint.
func SquareRootCalcPath() string {
	return "/square_root"
}