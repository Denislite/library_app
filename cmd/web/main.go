package main

import (
	"database/sql"
	"flag"
	"github.com/Denislite/library_app/pkg/models/dao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	model    *dao.Model
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP-address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Errot\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		model:    &dao.Model{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server at 127.0.0.1%s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rootroot@/librarydb")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
