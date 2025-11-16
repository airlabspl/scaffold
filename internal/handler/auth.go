package handler

import (
	"net/http"
	"scaffold/html"
	"scaffold/internal/form"
)

type loginPageData struct {
	Title  string
	Form   form.LoginForm
	Errors map[string]string
}

func LoginPage() http.HandlerFunc {
	t := html.New("base_layout.html", "login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, loginPageData{
			Title:  "Login",
			Form:   form.LoginForm{},
			Errors: make(map[string]string),
		})
	}
}

func LoginForm() http.HandlerFunc {
	t := html.New("login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		f := form.LoginForm{
			Email: r.FormValue("email"),
		}

		if errors := f.Validate(); len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			t.ExecuteTemplate(w, "login_form", loginPageData{
				Form:   f,
				Errors: errors,
			})
			return
		}
	}
}
