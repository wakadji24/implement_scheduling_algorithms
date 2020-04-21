package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

//Schedule for struct
type Schedule struct {
	Input  [][]int
	Result [][]int
}

// main program
func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", routeIndexGet)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}

}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		var tmpl = template.Must(template.New("form").ParseFiles("views/index.html", "views/_header.html"))

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

		fileLocation := filepath.Join(dir, "assets/files", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, uploadedFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filelocation := "assets/files/" + filename
		in := newTaskFromFile(filelocation)
		quicksort(in)

		result := schedulingProcess(in)
		data := Schedule{
			Input:  in,
			Result: result,
		}

		err = tmpl.ExecuteTemplate(w, "form", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "GET":
		var tmpl = template.Must(template.New("form").ParseFiles("views/index.html", "views/_header.html"))
		var err = tmpl.ExecuteTemplate(w, "form", nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
