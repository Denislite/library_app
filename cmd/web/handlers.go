package main

import (
	"html/template"
	"net/http"
	"strconv"
)

//home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./template/html/home.page.tmpl",
		"./template/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//all functions for authors
func (app *application) showAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := app.model.GetAllAuthors()
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	files := []string{
		"./template/html/author.page.tmpl",
		"./template/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, authors)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) createAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.errorLog.Println(err.Error())
			return
		}
		surname := r.FormValue("surname")
		name := r.FormValue("name")
		middle_name := r.FormValue("middle_name")
		image_link := r.FormValue("image_link")

		err = app.model.InsertAuthor(surname, name, middle_name, image_link)
		
		if err != nil {
			app.errorLog.Println(err.Error())
			return
		}
		http.Redirect(w, r, "/authors", 301)
	} else {
		files := []string{
			"./template/html/addauthor.page.tmpl",
			"./template/html/base.layout.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.errorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			app.errorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

//all functions for books
func (app *application) showBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

}

func (app *application) showBooks(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", 405)
		return
	}

}

//all functions for users
func (app *application) showUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

}

func (app *application) showUsers(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", 405)
		return
	}

}
