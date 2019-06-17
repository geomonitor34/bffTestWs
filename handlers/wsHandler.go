package handlers

import (
	"net/http"

	managers "github.com/geomonitor34/bffTestWs/managers"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

//WsHandler type
type WsHandler struct {
	connectionManager *managers.ClientManager
}

//NewWebsocketHandler create handler
func NewWebsocketHandler(connectionManager *managers.ClientManager) (handler WsHandler) {
	handler = WsHandler{
		connectionManager: connectionManager,
	}
	return
}

//ChatPage handler
func (h *WsHandler) ChatPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}

	uuidGen, _ := uuid.NewV4()

	client := &managers.Client{ID: uuidGen.String(), Socket: conn, Send: make(chan []byte)}

	h.connectionManager.Register <- client

	go client.Read(h.connectionManager)
	go client.Write(h.connectionManager)
}
