package main

import (
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world"

	wg.Add(1)
	go updateMessage("GoodBye")
	wg.Wait()

	if msg != "GoodBye" {
		t.Error("Incorrect value in msg")
	}
}