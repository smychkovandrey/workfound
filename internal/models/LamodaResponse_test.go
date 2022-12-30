package models

import (
	"testing"
)


func TestGetJobsLamoda(t *testing.T) {
	la := LamodaResponse {Results:[]resultLamoda{{Id: 4411, Title_ru: "Customer Research analyst", Slug: "customer-research-analyst-2"},
									 {Id: 4413, Title_ru: "Senior product analyst", Slug: "senior-product-analyst-2"}	}}
	got:= la.GetJobs()
	want:=[]Job{{Url: "https://job.lamoda.ru/vacancies/customer-research-analyst-2", Name:"Customer Research analyst"},
				{Url: "https://job.lamoda.ru/vacancies/senior-product-analyst-2", Name:"Senior product analyst"}}
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