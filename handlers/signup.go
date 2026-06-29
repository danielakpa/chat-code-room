package handlers

import (
	"html/template"
	"net/http"
)

type User_Pagedata struct {
	Pass_Errors string
	Em_Errors   string
	Name        string
	Gitea       string
	Email       string
	Password    string
}

type User struct {
	Name     string `json:"name"`
	Gitea    string `json:"gitea"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

var tmp = template.Must(template.ParseFiles("template/signup.html"))

func Signup(w http.ResponseWriter, r *http.Request) {
	Loaddata()
	if r.Method == "GET" {
		tmp.Execute(w, nil)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user := r.FormValue("name")
	gitea := r.FormValue("gitea")
	email := r.FormValue("email")
	password := r.FormValue("password")
	Confrim_password := r.FormValue("confirm_password")

	if password != Confrim_password {
		err_msg := User_Pagedata{
			Pass_Errors: "password do not match",
			Name:        user,
			Gitea:       gitea,
			Email:       email,
		}
		tmp.Execute(w, err_msg)
		return
	}

	for _, e := range users {
		if e.Email == email {
			pageData := User_Pagedata{
				Em_Errors: "Email already exists",
				Name:      user,
				Gitea:     gitea,
				Email:     email,
			}

			tmp.Execute(w, pageData)
			return
		}
	}

	users = append(users, User{
		Name:     user,
		Gitea:    gitea,
		Email:    email,
		Password: password,
	})

	Savedata()

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
