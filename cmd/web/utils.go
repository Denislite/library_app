package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadImage(imgType string, r *http.Request) (imgLink string) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("newFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

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
