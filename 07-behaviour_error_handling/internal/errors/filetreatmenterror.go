package errors

import (
	"github.com/pkg/errors"
)

type treatmentFileError struct {
	error
}

func TreatmentFileError(err error, format string, args ...interface{}) error {
	return &treatmentFileError{errors.Wrapf(err, format, args...)}
}

func NewTreatmentFileError(format string, args ...interface{}) error {
	return &treatmentFileError{errors.Errorf(format, args...)}
}
