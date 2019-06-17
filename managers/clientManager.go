//Package managers holds bussiness logic of manager
package managers

import (
	"encoding/json"

	models "github.com/geomonitor34/bffTestWs/models"
	"github.com/gorilla/websocket"
)

//Client fields
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

func (c *Client) Read(manager *ClientManager) {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			manager.Unregister <- c
			c.Socket.Close()
			break
		}

		jsonMessage, _ := json.Marshal(&models.Message{Sender: c.ID, Content: string(message)})
		manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write(manager *ClientManager) {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
