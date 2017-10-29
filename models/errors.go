package models

import "errors"

var (
	ErrNotFound = errors.New("unable to find resource")

	ErrToCreate = errors.New("")

	ErrDuplicated = errors.New("duplicated record")
)
