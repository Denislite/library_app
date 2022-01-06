package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
	"net/http"
)

func ValidateAuthor(r *http.Request) error {
	surname := r.FormValue("surname")
	name := r.FormValue("name")
	middleName := r.FormValue("middle_name")
	imagePath := UploadImage("authors", r)
	env.Env.Model.InsertAuthor(surname, name, middleName, imagePath)
	return nil
}

func GetAllAuthors() ([]*models.Author, error) {
	return env.Env.Model.GetAllAuthors()
}

func GetAuthorByID(id int) (*models.Author, error) {
	return env.Env.Model.GetAuthorByID(id)
}
