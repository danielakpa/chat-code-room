package handlers

import (
	"encoding/json"
	"os"
)

func Savedata() {

	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return
	}

	os.WriteFile("storage/users.json", data, 0666)
}

func Loaddata() {

	data, err := os.ReadFile("storage/users.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}
}
