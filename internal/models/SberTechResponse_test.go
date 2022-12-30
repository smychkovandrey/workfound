package models

import (
	"testing"
)


func TestGetJobsSberTech(t *testing.T) {
	or := SberTechResponse{
		Data: data{
			Vacancies: []vacancie{
				{ InternalId: 4118682, Title: "Senior/Lead iOS developer"	},
				{ InternalId: 4121157, Title: "Java team lead"	},
			},
			Total:     0,
		},
		Success: false,
	}
	got:= or.GetJobs(nil)
	want:=[]Job{{Url: "https://rabota.sber.ru/search/4118682", Name:"Senior/Lead iOS developer"},
				{Url: "https://rabota.sber.ru/search/4121157", Name:"Java team lead"}}
	if len(got)!=len(want){
		t.Errorf("Wrong lenght. got %q, wanted %q", len(got),len(want))
	}

	for i := range got {
		if got[i]!=want[i]{
			t.Errorf("Wrong elmnts. got %q, wanted %q", got, want)
			break
		}
		
	}	
}