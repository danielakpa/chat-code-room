package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

var temp = template.Must(template.ParseFiles("template/login.html"))

func login() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp.Execute(w, nil)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "bad request", 400)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	for _, look := range users {
		if look.Gitea != login {
			msg := User_Pagedata{
				Errors:   "invalid Gitea",
				Gitea:    login,
				Email:    login,
				Password: password,
			}
			temp.Execute(w, msg)
			return
		}

	}
	for _, look2 := range users {
		if look2.Email != login {
			msg2 := User_Pagedata{
				Errors:   "invalid email",
				Gitea:    login,
				Email:    login,
				Password: password,
			}
			temp.Execute(w, msg2)
			return
		}

	}
	for _, look3 := range users {
		if look3.Password != password {
			msg3 := User_Pagedata{
				Errors:   " invalid password",
				Gitea:    login,
				Email:    login,
				Password: password,
			}
			temp.Execute(w, msg3)
			return
		}
	}

	data := User{
		Gitea:    login,
		Email:    login,
		Password: password,
	}

	temp.Execute(w, data)

}
