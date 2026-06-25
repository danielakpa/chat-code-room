package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login() ([]Users, error) {
	data, err := os.ReadFile("users.json")
	if err != nil {
		return nil, err
	}
	var user []Users

	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request", 400)
		return
	}

	if r.Method == http.MethodPost {

	}

}
