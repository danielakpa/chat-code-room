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
	user, ok := Get_home_user(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// The room must already be claimed
	if !room.CodeLocked {
		http.Error(w, "no active code review", http.StatusForbidden)
		return
	}

	// Only the person who claimed the code can edit it
	if room.CodeOwner != user.Gitea {
		http.Error(w, "you are not allowed to edit this code", http.StatusForbidden)
		return
	}

	room.RoomCode = code

	SaveRoom()

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
