package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

//Result for struct
type Schedule struct {
	Result [][]int
}

func main() {
	runtime.GOMAXPROCS(2)

	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
	var err = tmpl.ExecuteTemplate(w, "form", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	filelocation := "files/" + filename
	in := newTaskFromFile(filelocation)
	quicksort(in)

	result := schedulingProcess(in)
	data := Schedule{
		Result: result,
	}

	err = tmpl.ExecuteTemplate(w, "result", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
