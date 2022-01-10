package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
	"net/http"
)

func ValidateUser(r *http.Request) error {
	surname := r.FormValue("surname")
	name := r.FormValue("name")
	middleName := r.FormValue("middle_name")
	passportData := r.FormValue("passport_data")
	birthdayDate := r.FormValue("birthday_date")
	email := r.FormValue("email")
	address := r.FormValue("address")
	return env.Env.Model.InsertUser(surname, name, middleName, passportData, birthdayDate, email, address)
}

func GetAllUsers() ([]*models.User, error) {
	return env.Env.Model.GetAllUsers()
}

func GetUserByID(id int) (*models.User, []*models.Book, error) {
	return env.Env.Model.GetUserByID(id)
}

func GiveBook(r *http.Request, id int) error {
	book := r.FormValue("book")
	returnDate := r.FormValue("return_date")
	return env.Env.Model.GiveBook(id, book, returnDate)
}

func TakeBook(r *http.Request, id int) error {
	book := r.FormValue("book")
	rating := r.FormValue("rating")
	return env.Env.Model.TakeBook(id, book, rating)
}
