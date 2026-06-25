package main

import "net/http"

func main() {
	http.HandleFunc("/", loginPage)

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.HandleFunc("/home", home)

	http.HandleFunc("/create-room", createRoom)
	http.HandleFunc("/join-room", joinRoom)
	http.HandleFunc("/room", Room)

	http.HandleFunc("/send-message", sendMessage)
	http.HandleFunc("/add-comment", addComment)

	http.HandleFunc("/claim-code", claimCode)
	http.HandleFunc("/update-code", updateCode)
	http.HandleFunc("/finish-review", finishReview)

	http.ListenAndServe(":8080", nil)
}
