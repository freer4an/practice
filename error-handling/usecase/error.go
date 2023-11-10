package usecase

import (
	"errors"
	"fmt"
)

type CustomServiceError struct {
	msg string
	err error
}

var ErrCase = errors.New("case error")

func (c *CustomServiceError) Error() string {
	return fmt.Sprintf("%s: %s", c.msg, c.err.Error())
}

func (c *CustomServiceError) Unwrap() error {
	return c.err
}
