package handlers

import "net/http"

func FinishReview(w http.ResponseWriter, r *http.Request) {
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

	room.CodeOwner = ""
	room.CodeLocked = false
	SaveRoom()

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
