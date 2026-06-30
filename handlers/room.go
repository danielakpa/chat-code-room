package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/room.html"))

type RoomPageData struct {
	User User
	Room *Rooms
}

func Room(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	roomID := r.URL.Query().Get("id")

	if roomID == "" {
		http.Error(w, "room ID required", http.StatusBadRequest)
		return
	}

	roomss, exist := rooms[roomID]
	if !exist {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}
	user, Exist := Get_home_user(r)
	if !Exist {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := RoomPageData{
		User: user,
		Room: roomss,
	}

	fmt.Println("Current user:", current_user.Gitea)
	fmt.Println("Room:", roomss.Name)

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
