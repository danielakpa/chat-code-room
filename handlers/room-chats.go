package handlers

import (
	"encoding/json"
	"os"
)

func Saveroom() {
	chat, err := json.MarshalIndent(rooms, "", " ")
	if err != nil {
		return
	}

	os.WriteFile("storage/chat.json", chat, 0666)

}

func Loadroom() {
	chat, err := os.ReadFile("storage/chat.json")
	if err != nil {
		return
	}

	json.Unmarshal(chat, &rooms)

}
