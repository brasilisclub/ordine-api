package product

import "errors"

var (
	ErrorProductNotFound = errors.New("A product with this id was not found")
)
