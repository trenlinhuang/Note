package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var Version = "1.0"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	} else if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	// str, _ := os.Getwd()
	log.Println(os.Getenv("STORAGE_ROOT"))
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer f.Close()
	io.Copy(w, f)
}
