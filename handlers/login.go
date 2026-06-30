package handlers

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseFiles("template/login.html"))

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
			sessionID := Generate_Login_ID()
			sessions[sessionID] = look

			cookies := &http.Cookie{
				Name:  "session_id",
				Value: sessionID,
				Path:  "/",
			}

			http.SetCookie(w, cookies)

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
