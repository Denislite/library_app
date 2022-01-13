package env

import (
	"database/sql"
	"github.com/Denislite/library_app/pkg/models/dao"
	"html/template"
	"log"
	"os"
)

type Environment struct {
	Model         *dao.Model
	TemplateCache map[string]*template.Template
	Email         EmailConfig
}

type EmailConfig struct {
	Email string
	Pass  string
	Host  string
	Port  string
}

var Env *Environment

func NewEnv(db *sql.DB) *Environment {

	templateCache, err := newTemplateCache("./template/html/")
	if err != nil {
		log.Fatal(err)
	}

	env := &Environment{
		Model:         &dao.Model{DB: db},
		TemplateCache: templateCache,
		Email: EmailConfig{
			Email: os.Getenv("EMAIL"),
			Pass:  os.Getenv("EMAIL_PASS"),
			Host:  os.Getenv("EMAIL_HOST"),
			Port:  os.Getenv("EMAIL_PORT"),
		},
	}

	return env
}
