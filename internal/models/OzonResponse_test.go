package models

import (
	"testing"
)


func TestGetJobsOzonResponse(t *testing.T) {
	or := OzonResponse {Items:[]item{{Hhid: 73035636, Title:"Старший разработчик Go, Отдел разработки ядра Банка Ozon"},
									 {Hhid: 72818483, Title:"Старший инженер по автоматизации тестирования (frontend), Телеметрия"}	}}
	got:= or.GetJobs()
	want:=[]Job{{Url: "https://job.ozon.ru/vacancy/73035636", Name:"Старший разработчик Go, Отдел разработки ядра Банка Ozon"},
				{Url: "https://job.ozon.ru/vacancy/72818483", Name:"Старший инженер по автоматизации тестирования (frontend), Телеметрия"}}
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