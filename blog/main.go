package hello

import (
    "html/template"
    "net/http"
)

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
    t, err := template.New("index").ParseFiles("templates/index.html")
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }
    err = t.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }
}
