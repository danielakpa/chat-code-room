package main

import (
	"chat-code-room/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Signup)
	http.HandleFunc("/login", handlers.Login)
	// http.HandleFunc("/logout", handlers.Logout)

	http.HandleFunc("/home", handlers.Home)

	http.HandleFunc("/test", handlers.CreateRoom)
	http.HandleFunc("/join-room", handlers.JoinRoom)
	http.HandleFunc("/room", handlers.Room)

	http.HandleFunc("/send-message", handlers.SendMessage)
	http.HandleFunc("/comment", handlers.Comment)

	http.HandleFunc("/claim-code", handlers.ClaimCode)
	http.HandleFunc("/update-code", handlers.UpdateCode)
	http.HandleFunc("/finish-review", handlers.FinishReview)

	fmt.Println("server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
