//Package models holds websocket request, response and other type declerations
package models

//Message websocket response fields
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}
