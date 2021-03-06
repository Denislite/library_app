package main

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/service"
	"log"
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
		log.Println(err)
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
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "authors.page.tmpl", &env.TemplateData{
		Authors: data,
	})
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		surname := r.FormValue("surname")
		name := r.FormValue("name")
		middleName := r.FormValue("middle_name")
		imagePath := service.UploadImage("authors", r)

		err := service.ValidateAuthor(surname, name, middleName, imagePath)

		if err != nil {
			log.Println(err)
			service.Render(w, r, "addauthor.page.tmpl", &env.TemplateData{
				Error: err,
			})
		}

		http.Redirect(w, r, "/authors", 301)
	} else {

		service.Render(w, r, "addauthor.page.tmpl", &env.TemplateData{})
	}
}

func showAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	author, err := service.GetAuthorByID(id)
	if err != nil {
		log.Println(err)
		return
	}

	books, err := service.GetBooksByAuthorID(id)
	if err != nil {
		log.Println(err)
		return
	}

	service.Render(w, r, "showauthor.page.tmpl", &env.TemplateData{
		Author: author,
		Books:  books,
	})
}

//all functions for books
func showBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	book, authors, err := service.GetBookByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "showbook.page.tmpl", &env.TemplateData{
		Book:    book,
		Authors: authors,
	})
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "books.page.tmpl", &env.TemplateData{
		Books: data,
	})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	authors, err := service.GetAllAuthors()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	genres, err := service.GetAllGenres()
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(30)
		name := r.FormValue("name")
		altName := r.FormValue("alt_name")
		price := r.FormValue("price")
		count := r.FormValue("count")
		pricePerDay := r.FormValue("price_per_day")
		year := r.FormValue("year")
		imageLink := service.UploadImage("books", r)
		authorsData := r.PostForm["author"]
		genre := r.FormValue("genre")

		err := service.ValidateBook(name, altName, imageLink, genre, price, count, pricePerDay, year, authorsData)
		if err != nil {
			log.Println(err)
			service.Render(w, r, "addbook.page.tmpl", &env.TemplateData{
				Error:   err,
				Authors: authors,
				Genres:  genres,
			})
		}

		http.Redirect(w, r, "/books", 301)
	} else {
		service.Render(w, r, "addbook.page.tmpl", &env.TemplateData{
			Authors: authors,
			Genres:  genres,
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

	user, books, err := service.GetUserByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "showuser.page.tmpl", &env.TemplateData{
		User:  user,
		Books: books,
	})
}

func showUsers(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "users.page.tmpl", &env.TemplateData{
		Users: data,
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		surname := r.FormValue("surname")
		name := r.FormValue("name")
		middleName := r.FormValue("middle_name")
		passportData := r.FormValue("passport_data")
		birthdayDate := r.FormValue("birthday_date")
		email := r.FormValue("email")
		address := r.FormValue("address")

		err := service.ValidateUser(surname, name, middleName, passportData, birthdayDate, email, address)
		if err != nil {
			log.Println(err)
			service.Render(w, r, "adduser.page.tmpl", &env.TemplateData{
				Error: err,
			})
		}

		http.Redirect(w, r, "/users", 301)
	} else {
		service.Render(w, r, "adduser.page.tmpl", &env.TemplateData{})
	}
}

func giveBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	books, err := service.GetAvailableBooks(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	user, _, err := service.GetUserByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if r.Method == http.MethodPost {
		book := r.FormValue("book")
		returnDate := r.FormValue("return_date")

		err = service.GiveBook(id, book, returnDate)
		if err != nil {
			log.Println(err)
			service.Render(w, r, "givebook.page.tmpl", &env.TemplateData{
				Error: err,
				Books: books,
				User:  user,
			})
		}
		http.Redirect(w, r, "/users", 301)
	} else {
		service.Render(w, r, "givebook.page.tmpl", &env.TemplateData{
			Books: books,
			User:  user,
		})
	}
}

func takeBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	user, books, err := service.GetUserByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if r.Method == http.MethodPost {
		book := r.FormValue("book")
		rating := r.FormValue("rating")
		fine := r.FormValue("fine")

		err = service.TakeBook(id, book, rating, fine)
		if err != nil {
			log.Println(err)
			service.Render(w, r, "takebook.page.tmpl", &env.TemplateData{
				Error: err,
				Books: books,
				User:  user,
			})
		}
		http.Redirect(w, r, "/users", 301)
	} else {
		service.Render(w, r, "takebook.page.tmpl", &env.TemplateData{
			Books: books,
			User:  user,
		})
	}
}

func showOrders(w http.ResponseWriter, r *http.Request) {
	active, err := service.GetActiveOrders()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	orders, err := service.GetClosedOrders()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	service.Render(w, r, "orders.page.tmpl", &env.TemplateData{
		UserBooks: active,
		Orders:    orders,
	})
}
