package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
)

func ValidateAuthor(surname, name, middleName, imagePath string) error {
	err := DataValidating(NAME, name)
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

	return env.Env.Model.InsertAuthor(surname, name, middleName, imagePath)
}

func GetAllAuthors() ([]*models.Author, error) {
	return env.Env.Model.GetAllAuthors()
}

func GetAuthorByID(id int) (*models.Author, error) {
	return env.Env.Model.GetAuthorByID(id)
}
