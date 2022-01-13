package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Denislite/library_app/env"
	"github.com/Denislite/library_app/pkg/autoupdate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8000", "HTTP-address")
	flag.Parse()

	if err := godotenv.Load("config.env"); err != nil {
		log.Print("No .env file found")
	}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")

	db, err := openDB(user, pass)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	env.Env = env.NewEnv(db)

	go autoupdate.UserChecker()

	log.Printf("Starting server at 127.0.0.1%s", *addr)
	err = http.ListenAndServe(*addr, routes())
	log.Fatal(err)
}

func openDB(user, pass string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+pass+"@/librarydb")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
