package handlers

import "net/http"

func Comment(w http.ResponseWriter, r *http.Request) {
	Loadroom()
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")
	comment := r.FormValue("comment")

	room, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "roomid required", http.StatusNotFound)
		return
	}

	com := Comments{
		User: current_user.Gitea,
		Text: comment,
	}
	Saveroom()

	room.Comments = append(room.Comments, com)

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
