package models

import "testing"

func TestGetNameYandex(t *testing.T) {
	got := InitYandex().GetName()
	want := "Yandex"
	if got != want {
		t.Errorf("Wrong Name. got %q, wanted %q", got, want)
	}
}