package handlers

import "net/http"

func SendMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")
	MessageText := r.FormValue("message")

	room, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "room ID required", http.StatusNotFound)
		return
	}

	user, Exist := Get_home_user(r)
	if !Exist {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	msg := Message{
		Users: user.Gitea,
		Text:  MessageText,
	}
	room.Messages = append(room.Messages, msg)

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)
}
