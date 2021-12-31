package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/author", app.showAuthor)
	mux.HandleFunc("/authors", app.showAuthors)
	mux.HandleFunc("/authors/create", app.createAuthor)
	mux.HandleFunc("/book", app.showBook)
	mux.HandleFunc("/books", app.showBooks)
	mux.HandleFunc("/books/create", app.createBook)
	mux.HandleFunc("/user", app.showUser)
	mux.HandleFunc("/users/", app.showUsers)
	mux.HandleFunc("/users/create", app.createUser)

	fileServer := http.FileServer(http.Dir("./template/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
