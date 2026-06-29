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

var current_user User

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	for _, look := range users {

		if look.Gitea == login || look.Email == login {

			if look.Password != password {
				temp.Execute(w, User_Pagedata{
					Pass_Errors: "Invalid password",
					Gitea:       login,
					Email:       login,
				})
				return
			}

			// Login successful
			current_user = look

			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	// User not found
	temp.Execute(w, User_Pagedata{
		Em_Errors: "Invalid Gitea username or Email",
		Gitea:     login,
		Email:     login,
	})
}
