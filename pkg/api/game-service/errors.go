package apigameservice

import "github.com/pkg/errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrWrongParameters = errors.New("wrong parameters")
	ErrInternal        = errors.New("internal server error")
)
