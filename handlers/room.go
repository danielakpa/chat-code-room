package handlers

import (
	"html/template"
	"net/http"
)

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

var tmpl = template.Must(template.ParseFiles("template/room.html"))

func Roomm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	roomID := r.URL.Query().Get("id")

	if roomID == "" {
		http.Error(w, "room ID required", http.StatusNotFound)
		return
	}

	roomss, Exist := rooms[roomID]
	if !Exist {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	} else {
		tmpl.Execute(w, roomss)
	}

}
