package bloggo

import (
	 "fmt"
    "net/http"
    "appengine"
    "github.com/hoisie/mustache"
)

func init() {
	 http.HandleFunc("/blog", blog)
	 http.HandleFunc("/contact", contact)
	 http.HandleFunc("/resume", resume)
	 http.HandleFunc("/bio", bio)
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	//For now, I am going to redirect the home page to the blog. However, in the future
	//I might flesh out the home page with other projects.
	http.Redirect(w, r, "/blog", http.StatusTemporaryRedirect)
}

func blog(w http.ResponseWriter, r *http.Request) {
	out := mustache.RenderFileInLayout("mustache/blog.html.mustache", "mustache/layout.html.mustache", nil)
	fmt.Fprint(w, out)
}

func resume(w http.ResponseWriter, r *http.Request) {
	out := mustache.RenderFileInLayout("mustache/resume.html.mustache", "mustache/layout.html.mustache", nil)
	fmt.Fprint(w, out)
}

func contact(w http.ResponseWriter, r *http.Request) {
	var submitted string
	if r.Method == "POST" {
		c := appengine.NewContext(r)
		addr := r.FormValue("email")
		info := r.FormValue("info")
		c.Infof("email %s, info %s", addr, info)
		submitted = "Contact information submitted!"
	}
	out := mustache.RenderFileInLayout("mustache/contact.html.mustache", "mustache/layout.html.mustache", map[string]string{"submitted":submitted})
	fmt.Fprint(w, out)
}

func bio(w http.ResponseWriter, r *http.Request) {
	out := mustache.RenderFileInLayout("mustache/bio.html.mustache", "mustache/layout.html.mustache", nil)
	fmt.Fprint(w, out)
}