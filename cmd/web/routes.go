package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/author", showAuthor)
	mux.HandleFunc("/authors", showAuthors)
	mux.HandleFunc("/authors/create", createAuthor)
	mux.HandleFunc("/book", showBook)
	mux.HandleFunc("/books", showBooks)
	mux.HandleFunc("/books/create", createBook)
	mux.HandleFunc("/user", showUser)
	mux.HandleFunc("/users/", showUsers)
	mux.HandleFunc("/users/create", createUser)
	mux.HandleFunc("/user/give", giveBook)
	mux.HandleFunc("/user/take", takeBook)
	mux.HandleFunc("/orders", showOrders)

	fileServer := http.FileServer(http.Dir("./template/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
