package infraerrors

import (
	"errors"
)

type InfrastructureError struct {
	detail string
}

func (e *InfrastructureError) Error() string {
	return e.detail
}

func (e *InfrastructureError) WrapWithError(err error) error {
	return errors.Join(e, err)
}
