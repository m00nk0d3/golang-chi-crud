package handlers

import "net/http"

type BookHandler struct {
}

func (b BookHandler) LisBooks(w http.ResponseWriter, r *http.Request)   {}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
