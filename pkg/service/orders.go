package service

import (
	"github.com/Denislite/library_app/env"
	"net/http"
	"time"
)

func GiveBook(r *http.Request, id int) error {
	book := r.FormValue("book")
	returnDate := r.FormValue("return_date")

	dt1, _ := time.Parse("2006-01-02", returnDate)
	now := -time.Now().Sub(dt1)
	days := int(now.Hours() / 24)

	user, _, err := GetUserByID(id)
	if err != nil {
		return err
	}

	var discount float64

	switch {
	case user.BookCount > 3:
		discount = 0.85
	case user.BookCount > 1:
		discount = 0.9
	default:
		discount = 1
	}

	return env.Env.Model.GiveBook(id, days, book, returnDate, discount)
}

func TakeBook(r *http.Request, id int) error {
	book := r.FormValue("book")
	rating := r.FormValue("rating")
	fine := r.FormValue("fine")

	user, _, err := GetUserByID(id)
	if err != nil {
		return err
	}

	var discount float64

	switch {
	case user.BookCount > 3:
		discount = 0.9
	default:
		discount = 1
	}

	return env.Env.Model.TakeBook(id, book, rating, fine, discount)
}
