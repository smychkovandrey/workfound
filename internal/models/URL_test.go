package models

import "testing"

func TestGetFullUrl(t *testing.T) {
	link:=Link{Host:"test", Scheme: "tts",Path:"var/1/2"}
	link.DeleteQueryParam("t1")
	got := link.GetFullURL()
	want := "tts://test/var/1/2"
	if got != want {
		t.Errorf("Wrong URL. got %q, wanted %q", got, want)
	}
	link.AddQueryParam("t1","test test")
	link.AddQueryParam("var1","1")
	link.AddQueryParam("var1","34")
	got = link.GetFullURL()
	want = "tts://test/var/1/2?t1=test+test&var1=1&var1=34"
	if got != want {
		t.Errorf("Wrong URL. got %q, wanted %q", got, want)
	}
	link.SetQueryParam("var2","1")
	link.SetQueryParam("var1","2")
	link.DeleteQueryParam("t1")
	link.DeleteQueryParam("t2")
	got = link.GetFullURL()
	want = "tts://test/var/1/2?var1=2&var2=1"
	if got != want {
		t.Errorf("Wrong URL. got %q, wanted %q", got, want)
	}
}
func TestGetQueryParam(t *testing.T){
	link:=Link{Host:"test", Scheme: "tts",Path:"var/1/2"}
	got,res := link.GetQueryParam("test")
	want := ""
	if got != want || res {
		t.Errorf("Wrong GetParam. got %q, wanted %q", got, want)
	}
	link.AddQueryParam("var1","1")
	got,res = link.GetQueryParam("var1")
	want = "1"
	if got != want || !res {
		t.Errorf("Wrong GetParam. got %q, wanted %q", got, want)
	}
}