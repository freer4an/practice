package storage

import (
	"errors"
	"fmt"
)

var Some_err = errors.New("failed to do something")

type Storage interface {
	DoSomething() error
}

type storage struct {
}

func (store *storage) DoSomething() error {
	return fmt.Errorf("error: %w", Some_err)
}

func New() Storage {
	return &storage{}
}
