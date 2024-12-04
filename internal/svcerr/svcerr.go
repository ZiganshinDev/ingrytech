package svcerr

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrNotFound      = errors.New("not found")
	ErrInternalError = errors.New("internal error")
)

func NewErr(baseErr error, msg string) error {
	return fmt.Errorf("%w: %s", baseErr, msg)
}
