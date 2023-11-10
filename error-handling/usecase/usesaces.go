package usecase

import (
	"errors"

	"github.com/freer4an/practice/error-handling/storage"
)

type UseCase interface {
	DoSomethingService() error
}

type useCase struct {
	storage storage.Storage
}

func New(storage storage.Storage) UseCase {
	return &useCase{storage: storage}
}

func (service *useCase) DoSomethingService() error {
	if err := service.storage.DoSomething(); err != nil {
		if errors.Is(err, storage.Some_err) {
			return &CustomServiceError{
				msg: "DoSomething()",
				err: err,
			}
		}
		return err
	}
	return nil
}
