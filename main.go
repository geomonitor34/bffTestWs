package main

import (
	"fmt"
	"net/http"

	"github.com/geomonitor34/bffTestWs/handlers"
	managers "github.com/geomonitor34/bffTestWs/managers"
)

func main() {
	fmt.Println("Starting application...")

	wsManager := managers.NewWebsocketManager()
	wsHandler := handlers.NewWebsocketHandler(&wsManager)

	go wsManager.Start()

	http.HandleFunc("/ws", wsHandler.ChatPage)

	http.ListenAndServe(":12345", nil)
}
