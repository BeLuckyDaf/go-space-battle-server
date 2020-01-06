package main

// Message is used as a general structure of a message
type Message struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
