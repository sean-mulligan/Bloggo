package hello

import (
	 "fmt"
    "net/http"
    "github.com/hoisie/mustache"
)

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	out := mustache.RenderFileInLayout("mustache/index.html.mustache", "mustache/layout.html.mustache", nil)
	fmt.Fprint(w, out)
}
