package apigameservice

import "github.com/pkg/errors"

var (
	ErrWrongParameters = errors.New("wrong parameters")
	ErrInternal        = errors.New("internal server error")
)
