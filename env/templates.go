package env

import (
	"github.com/Denislite/library_app/pkg/models"
	"html/template"
	"path/filepath"
)

type TemplateData struct {
	Author    *models.Author
	Authors   []*models.Author
	Book      *models.Book
	Books     []*models.Book
	User      *models.User
	Users     []*models.User
	Genres    []*models.Genre
	Orders    []*models.Order
	UserBooks []*models.UserBooks
	Error     error
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
