//Package managers holds bussiness logic of manager
package managers

import (
	"encoding/json"
	"fmt"

	models "github.com/geomonitor34/bffTestWs/models"
)

//ClientManager fields
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

//NewWebsocketManager manager create
func NewWebsocketManager() (manager ClientManager) {
	manager = ClientManager{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
	return
}

//Start func
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&models.Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			fmt.Println("<- manager.unregister")
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&models.Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}
