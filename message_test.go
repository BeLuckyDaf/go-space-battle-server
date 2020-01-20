package main

import "testing"

func TestMessage(t *testing.T) {
	m := Message{
		Status: true,
		Data:   "Hello",
	}

	if m.Status != true {
		t.Fatal("Status should have been true.")
	}
}
