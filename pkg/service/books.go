package service

import (
	"errors"
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
)

func ValidateBook(name, altName, imageLink, genre, price, count, pricePerDay, year string, authors []string) error {
	if genre == "" {
		return errors.New("syntax error")
	}
	err := DataValidating(FLOAT, price)
	if err != nil {
		return err
	}
	err = DataValidating(NUMBER, count)
	if err != nil {
		return err
	}
	err = DataValidating(FLOAT, pricePerDay)
	if err != nil {
		return err
	}
	err = YearValidating(year)
	if err != nil {
		return err
	}
	if len(authors) == 0 {
		return errors.New("syntax error")
	}

	return env.Env.Model.InsertBook(name, altName, imageLink, genre, price, count, pricePerDay, year, authors)
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

func GetAvailableBooks(id int) ([]*models.Book, error) {
	return env.Env.Model.GetAvailableBooks(id)
}
