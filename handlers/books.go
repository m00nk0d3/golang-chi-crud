package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/m00nk0d3/golang-chi-crud/models"
)

type BookHandler struct {
}

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"name"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.ListBooks())
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := models.GetBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.StoreBook(book)

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook := models.UpdateBook(id, book)
	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := models.DeleteBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
