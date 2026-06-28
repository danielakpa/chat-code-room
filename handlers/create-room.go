package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

func GenerateId() string {
	id := make([]byte, 3)
	_, err := rand.Read(id)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(id)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	room_name := r.FormValue("roomname")

	if room_name == "" {
		http.Error(w, "room required", http.StatusBadRequest)
		return
	}

	if len(room_name) > 15 {
		http.Error(w, "name to long", http.StatusBadRequest)
		return
	}

	for _, rm := range rooms {
		if rm.Name == room_name {
			http.Error(w, "room Exist already", http.StatusNotFound)
			return
		}
	}
	roomID := GenerateId()

	room := &Rooms{
		ID:         roomID,
		Name:       room_name,
		Owner:      current_user.Gitea,
		Code:       "",
		CodeOwner:  "",
		CodeLocked: false,
		Messages:   []Message{},
		Comments:   []Comments{},
	}

	rooms[roomID] = room

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
