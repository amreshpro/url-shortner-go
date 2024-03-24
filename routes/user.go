package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/amreshpro/url-shortner-go/controller"
)

func UserRouter() error {
	userRouter := chi.NewRouter()
	userRouter.Get("/user", controller.getUser())
	userRouter.Get("/test", controller.testUser())

	return nil
}
