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

	msg := Message{
		Users: current_user.Gitea,
		Text:  MessageText,
	}
	room.Messages = append(room.Messages, msg)

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)
}
