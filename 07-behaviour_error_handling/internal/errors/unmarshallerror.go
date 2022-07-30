package errors

import (
	"github.com/pkg/errors"
)

type unmarshallError struct {
	error
}

func UnmarshallError(err error, format string, args ...interface{}) error {
	return &unmarshallError{errors.Wrapf(err, format, args...)}
}

func NewUnmarshallError(format string, args ...interface{}) error {
	return &unmarshallError{errors.Errorf(format, args...)}
}
