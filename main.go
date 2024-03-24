package main

import (
	"fmt"
	userRouter "github.com/amreshpro/url-shortner-go/routes/userroutes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	// load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	//main router
	appRouter := chi.NewRouter()

	//logger middleware
	appRouter.Use(middleware.Logger)

	appRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	appRouter.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	appRouter.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	// user router
	appRouter.Mount("/api/v1/", userRouter.userRouter)

	PORT := os.Getenv("PORT")
	fmt.Println(PORT)
	http.ListenAndServe(":"+PORT, appRouter)
}
