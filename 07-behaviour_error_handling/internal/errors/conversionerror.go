package errors

import (
	"github.com/pkg/errors"
)

type conversionError struct {
	error
}

func WraperConversionError(err error, format string, args ...interface{}) error {
	return &conversionError{errors.Wrapf(err, format, args...)}
}

func NewWraperConversionError(format string, args ...interface{}) error {
	return &conversionError{errors.Errorf(format, args...)}
}
