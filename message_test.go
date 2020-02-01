// Copyright 2020 Vladislav Smirnov

package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	m := Message{
		Status: true,
		Data:   "Hello",
	}

	s := fmt.Sprint(m.Data)
	if m.Status != true || strings.Compare("Hello", s) != 0 {
		t.Fatal("Message test failed.")
	}
}
