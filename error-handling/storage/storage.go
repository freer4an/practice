package storage

import (
	"errors"
	"fmt"
)

var Some_err = errors.New("failed to do something")

type Storage interface {
	DoSomething() error
	DoSomethingAgain() error
}

type storage struct {
}

func (store *storage) DoSomething() error {
	return fmt.Errorf("error: %w", Some_err) // with %w we can wrap error
}

func (store *storage) DoSomethingAgain() error {
	return fmt.Errorf("error: %v", Some_err) // with %v format error changes so the error will be hard to handle
}

func New() Storage {
	return &storage{}
}
