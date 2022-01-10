package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
	"net/http"
)

func ValidateBook(r *http.Request) error {
	r.ParseMultipartForm(30)
	name := r.FormValue("name")
	altName := r.FormValue("alt_name")
	genre := r.FormValue("genre")
	price := r.FormValue("price")
	count := r.FormValue("count")
	pricePerDay := r.FormValue("price_per_day")
	year := r.FormValue("year")
	imageLink := UploadImage("books", r)
	authors := r.PostForm["author"]

	env.Env.Model.InsertBook(name, altName, imageLink, genre, price, count, pricePerDay, year, authors)
	return nil
}

func GetAllBooks() ([]*models.Book, error) {
	return env.Env.Model.GetAllBooks()
}

func GetBookByID(id int) (*models.Book, []*models.Author, error) {
	return env.Env.Model.GetBookByID(id)
}

func GetTopBooks() ([]*models.Book, error) {
	return env.Env.Model.GetTopBooks()
}

func GetAvailableBooks() ([]*models.Book, error) {
	return env.Env.Model.GetAvailableBooks()
}
