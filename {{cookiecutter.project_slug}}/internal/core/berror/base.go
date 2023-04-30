package berror

import (
	"errors"
	"strconv"
)

type BusinessError struct {
	code   int
	Detail string
}

func (e *BusinessError) Code() int {
	return e.code
}

func (e *BusinessError) Error() string {
	return e.Detail + ": error code " + strconv.Itoa(e.code)
}

func (e *BusinessError) WrapWithError(err error) error {
	return errors.Join(e, err)
}
