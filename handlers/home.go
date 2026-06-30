package handlers

import (
	"html/template"
	"net/http"
)

type Home_Data struct {
	Home_user User
	Home_Room []*Rooms
}

func Getroom() []*Rooms {
	var Allroms []*Rooms

	for _, room := range rooms {
		Allroms = append(Allroms, room)
	}
	return Allroms
}

func Home(w http.ResponseWriter, r *http.Request) {
	home_temp, err := template.ParseFiles("template/home.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user,Exist := Get_home_user(r)
	if !Exist {
		http.Redirect(w,r, "/login", http.StatusSeeOther)
		return
	}


	data := Home_Data{
		Home_user: user,
		Home_Room: Getroom(),
	}



	home_temp.Execute(w, data)

}
