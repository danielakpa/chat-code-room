package handlers

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/room.html"))

func Room(w http.ResponseWriter, r *http.Request) {
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
