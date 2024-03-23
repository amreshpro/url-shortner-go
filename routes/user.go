package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/amreshpro/go-url-shortner/controller/userController"
)

func UserRouter() error {
	userRouter := chi.NewRouter()
	userRouter.Get("/user", userController.getUser)
	userRouter.Get("/test", userController.testUser)

	return nil
}
