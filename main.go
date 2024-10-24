package main

import (
	"fmt"
	"net/http"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "./static/index.html")
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "./static/BookList.html")
}

func main() {
	fmt.Println("HTMX For Beginners")
	SM := http.NewServeMux()
	Server := &http.Server{Addr: ":8080", Handler: SM}
	fileServer := http.FileServer(http.Dir("./static"))
	SM.Handle("/static/", http.StripPrefix("/static", fileServer))
	SM.HandleFunc("/", mainPageHandler)
	SM.HandleFunc("/books", booksHandler)
	err := Server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
