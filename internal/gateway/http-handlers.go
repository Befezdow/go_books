package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import (
	"github.com/befezdow/go-books-rest-api/internal/core"
	"github.com/befezdow/go-books-rest-api/internal/shared"
)

func (s *Gateway) handleGetBooks() http.HandlerFunc {
	// here can be local variables and types
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get books invoked")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(core.GetBooks())
	}
}

func (s *Gateway) handleCreateBook() http.HandlerFunc {
	// here can be local variables and types
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create book invoked")
		w.Header().Set("Content-Type", "application/json")
		var book shared.Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book = core.AddBook(book)
		json.NewEncoder(w).Encode(book)
	}
}
