package handlers

import "net/http"

type Message struct {
	Users string
	Text  string
}

type Comment struct {
	User string
	Text string
}

type Room struct {
	ID         string
	Code       string
	CodeOwner  string
	CodeLocked bool
	Messages   []Message
	Comments   []Comment
}

var rooms = make(map[string]*Room)

func sendMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")
	MessageText := r.FormValue("message")

	roomss, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "room ID required", http.StatusNotFound)
		return
	}

	msg := Message{
		Users: "",
		Text:  MessageText,
	}
	roomss.Messages = append(roomss.Messages, msg)

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)
}
