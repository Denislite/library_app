package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/service"
	"net/http"
	"strconv"
)

//home page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	render(w, r, "home.page.tmpl", &env.TemplateData{})
}

//all functions for authors
func showAuthors(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllAuthors()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	render(w, r, "author.page.tmpl", &env.TemplateData{
		Authors: data,
	})
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		surname := r.FormValue("surname")
		name := r.FormValue("name")
		middleName := r.FormValue("middle_name")
		imageLink := uploadImage("authors", r)

		err := service.ValidateAuthor(surname, name, middleName, imageLink, r)

		if err != nil {
			fmt.Println(err)
			return
		}

		http.Redirect(w, r, "/authors", 301)
	} else {

		render(w, r, "addauthor.page.tmpl", &env.TemplateData{})
	}
}

func showAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	data, err := service.GetAuthorByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return
		} else {
			return
		}
		return
	}
	fmt.Fprintf(w, "%v", data)
}

//all functions for books
func showBook(w http.ResponseWriter, r *http.Request) {

}

func showBooks(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

//all functions for users
func showUser(w http.ResponseWriter, r *http.Request) {

}

func showUsers(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

}
