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
	for id, room := range rooms {
		if room.Name == roomID {
			http.Redirect(w, r, "/room?id="+id, http.StatusSeeOther)
		} else {
			http.Error(w, "roon not found", http.StatusNotFound)
			return
		}
	}
}
