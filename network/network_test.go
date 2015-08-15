package network

import (
	"testing"
)

func TestPing(t *testing.T) {
	target := "http://envy.blamestross.com:8000/"
	want := "\"PONG\""
	nw := new(Networking)
	got := nw.Ping(target)
	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}
}
