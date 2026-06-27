package handlers

import "net/http"

func UpdateCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomID")
	code := r.FormValue("code")

	room, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "room id not found", http.StatusNotFound)
		return
	}

	room.RoomCode = code

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
