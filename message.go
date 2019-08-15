package main

type Message struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
}
