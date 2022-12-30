package models

import (
	"testing"
)


func TestGetJobsYandex(t *testing.T) {
	or := YandexResponse {Results:[]result{{Id: 73035636, Title:"Старший разработчик Go, Отдел разработки ядра Банка Ozon"},
									 {Id: 72818483, Title:"Старший инженер по автоматизации тестирования (frontend), Телеметрия"}	}}
	got:= or.GetJobs()
	want:=[]Job{{Url: "https://yandex.ru/jobs/vacancies/73035636", Name:"Старший разработчик Go, Отдел разработки ядра Банка Ozon"},
				{Url: "https://yandex.ru/jobs/vacancies/72818483", Name:"Старший инженер по автоматизации тестирования (frontend), Телеметрия"}}
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