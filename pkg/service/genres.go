package service

import (
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/models"
)

func GetAllGenres() ([]*models.Genre, error) {
	return env.Env.Model.GetAllGenres()
}
