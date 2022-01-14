package service

import (
	"errors"
	"fmt"
	"github.com/Denislite/library_app/env"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	EMAIL  = `^[a-zA-Z0-9]+@[a-zA-Z]+.[a-zA-Z]+$`
	NAME   = `^[A-Z][a-z]+$`
	FLOAT  = `^[1-9][0-9]*\.?[0-9]*$`
	NUMBER = `^[1-9][0-9]*$`
	PASS   = `^[A-Z]{2}[0-9]{7}$`
	RATING = `^[0-5]$`
)

func UploadImage(imgType string, r *http.Request) (imgLink string) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(imgType)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(imgType)

	tempFile, err := ioutil.TempFile("./template/static/"+imgType, "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tempFile.Write(fileBytes)

	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	imgLink = tempFile.Name()[8:]

	return imgLink
}

func Render(w http.ResponseWriter, r *http.Request, name string, td *env.TemplateData) {
	ts, ok := env.Env.TemplateCache[name]
	if !ok {
		fmt.Println("!!!")
		return
	}
	err := ts.Execute(w, td)
	if err != nil {
		return
	}
}

func DataValidating(pattern, data string) error {
	matched, _ := regexp.Match(pattern, []byte(data))
	if matched == false {
		return errors.New("syntax error")
	}
	return nil
}

func YearValidating(data string) error {
	currentTime := time.Now()
	year, err := strconv.Atoi(data)
	if err != nil {
		return err
	}

	if year > currentTime.Year() {
		return errors.New("syntax error")
	}

	return nil
}

func BirthdayDateValidating(date string) error {
	dateValue, _ := time.Parse("2006-01-02", date)
	if dateValue.After(time.Now()) {
		return errors.New("wrong date")
	}
	return nil
}

func DateValidating(date string) error {
	dateValue, _ := time.Parse("2006-01-02", date)
	if dateValue.Before(time.Now()) || dateValue.After(time.Now().AddDate(0, 0, 30)) {
		return errors.New("wrong date")
	}
	return nil
}
