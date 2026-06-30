package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

func SaveRoom() {
	room, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return
	}

	os.WriteFile("strorage/room.json", room, 0666)
}

func LoadRoom() {
	room, err := os.ReadFile("storage/room.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(room, &users)
	if err != nil {
		return
	}

}

func Savelogindata() {

	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return
	}

	os.WriteFile("storage/users.json", data, 0666)
}

func Logindata() {

	data, err := os.ReadFile("storage/users.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}
}

func Generate_Login_ID() string {
	id := make([]byte, 3)

	_, err := rand.Read(id)
	if err != nil {
		return ""

	}

	return hex.EncodeToString(id)
}

func Get_home_user(r *http.Request) (User, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return User{}, false
	}

	user, Exist := sessions[cookie.Value]
	if !Exist {
		return User{}, false
	}
	return user, true

}
