package models

import "errors"

var (
	ErrNotFound = errors.New("unable to find resource")

	ErrToCreate = errors.New("")

	ErrDuplicate = errors.New("duplicated record")
)
