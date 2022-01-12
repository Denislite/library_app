package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
)

func ValidateUser(surname, name, middleName, passportData, birthdayDate, email, address string) error {
	err := DataValidating(NAME, name)
	if err != nil {
		return err
	}

	err = BirthdayDateValidating(birthdayDate)
	if err != nil {
		return err
	}

	err = DataValidating(NAME, surname)
	if err != nil {
		return err
	}

	err = DataValidating(NAME, middleName)
	if err != nil {
		return err
	}

	if passportData != "" {
		err = DataValidating(PASS, passportData)
		if err != nil {
			return err
		}
	}

	err = DataValidating(EMAIL, email)
	if err != nil {
		return err
	}

	return env.Env.Model.InsertUser(surname, name, middleName, passportData, birthdayDate, email, address)
}

func GetAllUsers() ([]*models.User, error) {
	return env.Env.Model.GetAllUsers()
}

func GetUserByID(id int) (*models.User, []*models.Book, error) {
	return env.Env.Model.GetUserByID(id)
}
