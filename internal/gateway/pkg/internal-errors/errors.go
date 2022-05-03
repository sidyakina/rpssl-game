package internalerrors

import "github.com/pkg/errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrWrongParameters = errors.New("wrong parameters")
)
