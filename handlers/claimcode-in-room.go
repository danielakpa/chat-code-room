package handlers

import "net/http"

func ClaimCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")

	room, exist := rooms[roomID]
	if !exist {
		http.Error(w, "room id not found", http.StatusNotFound)
		return
	}

	if room.CodeLocked {
		http.Error(w, "code is already looked", http.StatusForbidden)
		return
	}

	room.CodeLocked = true

	room.CodeOwner = current_user.Gitea

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)
}
