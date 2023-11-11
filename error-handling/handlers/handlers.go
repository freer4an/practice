package handlers

import (
	"errors"
	"log"

	"github.com/freer4an/practice/error-handling/storage"
	"github.com/freer4an/practice/error-handling/usecase"
)

type Handler interface {
	HandleRequest()
	HandleRequestAgain()
}

type handler struct {
	usecase usecase.UseCaseService
}

func NewHandler(usecase usecase.UseCaseService) Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) HandleRequest() {
	err := h.usecase.DoSomething()
	if err != nil {
		// errors.Is() can define Some_err
		if errors.Is(err, storage.Some_err) {
			log.Println(err)
			return
		}
		log.Printf("couldn't handle %v", err)
	}
}

func (h *handler) HandleRequestAgain() {
	err := h.usecase.DoSomethingAgain()
	if err != nil {
		// errors.Is() can not define Some_err
		if errors.Is(err, storage.Some_err) {
			log.Println(err)
			return
		}
		log.Printf("couldn't handle %v", err)
	}
}
