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

	data, err := service.GetTopBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "home.page.tmpl", &env.TemplateData{
		Books: data,
	})
}

//all functions for authors
func showAuthors(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllAuthors()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "authors.page.tmpl", &env.TemplateData{
		Authors: data,
	})
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := service.ValidateAuthor(r)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		http.Redirect(w, r, "/authors", 301)
	} else {

		service.Render(w, r, "addauthor.page.tmpl", &env.TemplateData{})
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

	service.Render(w, r, "showauthor.page.tmpl", &env.TemplateData{
		Author: data,
	})
}

//all functions for books
func showBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	data, err := service.GetBookByID(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "showbook.page.tmpl", &env.TemplateData{
		Book: data,
	})
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "books.page.tmpl", &env.TemplateData{
		Books: data,
	})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := service.ValidateBook(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		http.Redirect(w, r, "/books", 301)
	} else {
		data, err := service.GetAllAuthors()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		service.Render(w, r, "addbook.page.tmpl", &env.TemplateData{
			Authors: data,
		})
	}
}

//all functions for users
func showUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	data, err := service.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "showuser.page.tmpl", &env.TemplateData{
		User: data,
	})
}

func showUsers(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "users.page.tmpl", &env.TemplateData{
		Users: data,
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := service.ValidateUser(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		http.Redirect(w, r, "/users", 301)
	} else {
		service.Render(w, r, "adduser.page.tmpl", &env.TemplateData{})
	}
}

func giveBook(w http.ResponseWriter, r *http.Request) {

}

func takeBook(w http.ResponseWriter, r *http.Request) {

}
