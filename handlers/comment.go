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

func addComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	roomID := r.FormValue("roomid")
	comment := r.FormValue("comment")

	roomss, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "roomid required", http.StatusNotFound)
		return
	}

	com := Comment{
		User: "",
		Text: comment,
	}

	roomss.Comments = append(roomss.Comments, com)

	http.Redirect(w, r, "/room?id="+roomID, http.StatusSeeOther)

}
