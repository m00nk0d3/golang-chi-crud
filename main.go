package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/m00nk0d3/golang-chi-crud/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/books", BookRoutes())
	http.ListenAndServe(port, r)
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bookHandler := handlers.BookHandler{}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}
