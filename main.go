package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	port = 5432
)

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add Book page")
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List Book page")
}

func downloadBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Download Books page")
}

func viewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "View Book page")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database.....")
	http.HandleFunc("/books", listBooks)
	http.HandleFunc("/book:book-id", viewBook)
	http.HandleFunc("/download", downloadBook)
	fmt.Println("server is running on Port 8080...")
	http.ListenAndServe(":8080", nil)
}
