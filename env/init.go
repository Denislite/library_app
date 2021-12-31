package env

import (
	"database/sql"
	"github.com/Denislite/library_app/pkg/models/dao"
	"html/template"
	"log"
	"os"
)

type Environment struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Model         *dao.Model
	TemplateCache map[string]*template.Template
}

var Env *Environment

func NewEnv(db *sql.DB) *Environment {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./template/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	env := &Environment{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Model:         &dao.Model{DB: db},
		TemplateCache: templateCache,
	}

	return env
}
