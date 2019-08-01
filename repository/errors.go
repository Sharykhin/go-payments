package repository

import "errors"

var (
	// NotFound error represents general error when a record could not be found
	NotFound = errors.New("record was not found")
)
