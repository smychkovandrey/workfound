package models

import "testing"

func TestGetNameOzon(t *testing.T) {
	got := InitOzon().GetName()
	want := "Ozon"
	if got != want {
		t.Errorf("Wrong Name. got %q, wanted %q", got, want)
	}
}
