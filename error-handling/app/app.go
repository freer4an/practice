package app

import (
	"github.com/freer4an/practice/error-handling/handlers"
	"github.com/freer4an/practice/error-handling/storage"
	"github.com/freer4an/practice/error-handling/usecase"
)

func Run() {
	storage := storage.New()
	usecase := usecase.New(storage)
	handler := handlers.NewHandler(usecase)
	handler.HandleRequest()      // -> DoSomething(): failed to do something
	handler.HandleRequestAgain() // -> couldn't handle error: failed to do something
	// same logic but different errors
}
