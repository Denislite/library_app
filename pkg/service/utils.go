package service

import (
	"fmt"
	"github.com/Denislite/library_app/env"
	"io/ioutil"
	"net/http"
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
