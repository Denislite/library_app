package service

import (
	"errors"
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
	"time"
)

func GiveBook(id int, book, returnDate string) error {
	if book == "" {
		return errors.New("syntax error")
	}

	err := DateValidating(returnDate)
	if err != nil {
		return err
	}

	date, _ := time.Parse("2006-01-02", returnDate)
	now := -time.Now().Sub(date)
	days := int(now.Hours()/24) + 1

	user, _, err := GetUserByID(id)
	if err != nil {
		return err
	}

	var discount float64

	switch {
	case user.BookCount == 5:
		return errors.New("maximum book count")
	case user.BookCount > 3:
		discount = 0.85
	case user.BookCount > 1:
		discount = 0.9
	default:
		discount = 1
	}

	return env.Env.Model.GiveBook(id, days, book, returnDate, discount)
}

func TakeBook(id int, book, rating, fine string) error {
	err := DataValidating(FLOAT, fine)
	if err != nil {
		return err
	}

	err = DataValidating(RATING, rating)
	if err != nil {
		return err
	}

	if book == "" {
		return errors.New("syntax error")
	}

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

func UpdateDuty() error {
	return env.Env.Model.UpdateDuty()
}

func GetUsersWithDuty() ([]*models.User, error) {
	return env.Env.Model.GetUsersWithDuty()
}

func GetClosedOrders() ([]*models.Order, error) {
	return env.Env.Model.GetClosedOrders()
}

func GetActiveOrders() ([]*models.UserBooks, error) {
	return env.Env.Model.GetActiveOrders()
}
