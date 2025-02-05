package order

import "errors"

var (
	ErrorOrderNotFound = errors.New("A order with provided id not founded")
)
