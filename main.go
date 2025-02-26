package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	port = 5432
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books     []book
	bookMutex sync.Mutex
)

func generateRandomId() string {
	return uuid.New().String()
}
func listBooksHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook book
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read from request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &newBook)
	if err != nil || newBook.Author == "" {
		http.Error(w, "No Author added", http.StatusBadRequest)
		return
	} else if err != nil || newBook.Title == "" {
		http.Error(w, "No Book titile added", http.StatusBadRequest)
		return
	}

	newBook.ID = generateRandomId()

	bookMutex.Lock()
	books = append(books, newBook)
	bookMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)

}

func viewBookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/add", addBookHandler)
	http.HandleFunc("/books", listBooksHandler)
	http.HandleFunc("/books/", viewBookHandler)
	fmt.Println("server is running on Port 8080...")
	http.ListenAndServe(":8080", nil)
}
