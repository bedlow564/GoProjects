package main

import (
	"go_cmd/hello"
	"testing"
)

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},

		{
			items: []string{"Brandyn"},
			result: "Hello, Brandyn!",
		},
		{
			items: []string{"Aylah"},
			result: "Hello, Aylah!",
		},
	}

	for _, st := range subtests {
		if s := hello.Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
