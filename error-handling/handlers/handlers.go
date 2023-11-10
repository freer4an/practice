package handlers

import (
	"errors"
	"log"

	"github.com/freer4an/practice/error-handling/storage"
	"github.com/freer4an/practice/error-handling/usecase"
)

type Handler interface {
	HandleRequest()
}

type handler struct {
	usecase usecase.UseCase
}

func NewHandler(usecase usecase.UseCase) Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) HandleRequest() {
	err := h.usecase.DoSomethingService()
	if err != nil {
		if errors.Is(err, storage.Some_err) {
			log.Fatalf("couldn't handle %v", err)
			return
		}
		log.Println(err)
	}
}
