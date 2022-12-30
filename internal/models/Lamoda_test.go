package models

import "testing"

func TestGetName(t *testing.T) {
	got := InitLamoda().GetName()
	want := "Lamoda"
	if got != want {
		t.Errorf("Wrong Name. got %q, wanted %q", got, want)
	}
}