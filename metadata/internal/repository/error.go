package repository

import "errors"

var (
	notFound = "not found"
)

var ErrNotFound = errors.New(notFound)
