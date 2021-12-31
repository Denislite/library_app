package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
	"net/http"
)

func ValidateAuthor(surname, name, middleName, imagePath string, r *http.Request) error {
	env.Env.Model.InsertAuthor(surname, name, middleName, imagePath)
	return nil
}

func GetAllAuthors() ([]*models.Author, error) {
	return env.Env.Model.GetAllAuthors()
}

func GetAuthorByID(id int) (*models.Author, error) {
	return env.Env.Model.GetAuthorByID(id)
}
