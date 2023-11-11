package usecase

import (
	"errors"

	"github.com/freer4an/practice/error-handling/storage"
)

type UseCaseService interface {
	DoSomething() error
	DoSomethingAgain() error
}

type useCase struct {
	storage storage.Storage
}

func New(storage storage.Storage) UseCaseService {
	return &useCase{storage: storage}
}

func (service *useCase) DoSomething() error {
	if err := service.storage.DoSomething(); err != nil {
		if errors.Is(err, storage.Some_err) {
			return &CustomServiceError{
				msg: "DoSomething()",
				err: errors.Unwrap(err),
			}
		}
		return err
	}
	return nil
}

func (service *useCase) DoSomethingAgain() error {
	if err := service.storage.DoSomethingAgain(); err != nil {
		if errors.Is(err, storage.Some_err) {
			return &CustomServiceError{
				msg: "DoSomething()",
				err: errors.Unwrap(err),
			}
		}
		return err
	}
	return nil
}
