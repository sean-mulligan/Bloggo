package bloggo

import (
	 "fmt"
    "net/http"
    "appengine"
    "appengine/mail"
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
		name := r.FormValue("name")
		email := r.FormValue("email")
		info := r.FormValue("info")
		msg := &mail.Message{
			Sender: fmt.Sprintf("%s (%s)", name, email),
			To: []string{"sean.mulligan.cs@gmail.com"},
			Subject: fmt.Sprintf("Website Contact - %s", name),
			Body: info,
		}
		if err := mail.Send(c, msg); err != nil {
			c.Errorf("Could not send email: %v", err)
			submitted = "Your information could not be sent. Apologies!"		
		} else {
			submitted = "Your information has been sent. I'll get back to you as soon as possible!"
		}
		c.Infof("Contact submitted: name=%s, email=%s, info=%s", name, email, info)
		
	}
	out := mustache.RenderFileInLayout("mustache/contact.html.mustache", "mustache/layout.html.mustache", map[string]string{"submitted":submitted})
	fmt.Fprint(w, out)
}

func bio(w http.ResponseWriter, r *http.Request) {
	out := mustache.RenderFileInLayout("mustache/bio.html.mustache", "mustache/layout.html.mustache", nil)
	fmt.Fprint(w, out)
}