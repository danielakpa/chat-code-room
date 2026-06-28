package handlers

import (
	"net/http"
)

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")

	if roomID == "" {
		http.Error(w, "roomid required", http.StatusBadRequest)
		return
	}

	_, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "room not found", http.StatusNotFound)
		return

	} else {
		http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)
	}
}
