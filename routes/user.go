package routes

import (
	userController "github.com/amreshpro/url-shortner-go/controller/usercontroller"
	"github.com/go-chi/chi/v5"
)

func UserRouter() error {
	userRouter := chi.NewRouter()
	userRouter.Get("/user", userController.getUser())
	userRouter.Get("/test", userController.testUser())

	return nil
}
