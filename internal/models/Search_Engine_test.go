package models

import (
	"sort"
	"strings"
	"testing"
)

type req1 struct {}

func (req1) GetName() string {
	return "req1"
}

func (req1) GetJobs() ([]Job,error){
	return []Job {
		{Url: "req1/1", Name: "req1 1 "},
		{Url: "req1/2", Name: "req1 2 "},
		{Url: "req1/3", Name: "req1 3 "},
		{Url: "req1/4", Name: "req1 4 "},
	},nil
}

type req2 struct {}

func (req2) GetName() string {
	return "req2"
}

func (req2) GetJobs() ([]Job,error){
	return []Job {
		{Url: "req2/1", Name: "req2 1 "},
		{Url: "req2/2", Name: "req2 2 "},
		{Url: "req2/3", Name: "req2 3 "},
		{Url: "req2/4", Name: "req2 4 "},
	},nil
}

type req3 struct {}
func (req3) GetName() string {
	return "req3"
}

func (req3) GetJobs() ([]Job,error){
	return []Job {
		{Url: "req3/1", Name: "req3 1 "},
		{Url: "req3/2", Name: "req3 2 "},
		{Url: "req3/3", Name: "req3 3 "},
		{Url: "req3/4", Name: "req3 4 "},
	},nil
}


func TestGetGetResult(t *testing.T) {
	
	se:=search_Engine {Requsters: []JobRequester {req1{},req2{},req3{}},
					Filter: func(j Job) bool {return !strings.Contains(j.Name, " 3")},}

	results:=se.GetResults()
	got:= make([]Results,0)
	for result:=range results{
		sort.SliceStable(result.Jobs,func(i, j int) bool {
			return result.Jobs[i].Name < result.Jobs[j].Name
		})
		got = append(got, result)
	}
	sort.SliceStable(got,func(i, j int) bool {
		return got[i].Name < got[j].Name
	})
	want:=[]Results{{Name: "req1",
					Error: nil, 
					Jobs:[]Job{{Url: "req1/1", Name: "req1 1 "},
						       {Url: "req1/2", Name: "req1 2 "},
						       {Url: "req1/4", Name: "req1 4 "},
								}},
					{Name: "req2",
					Error: nil, 
					Jobs:[]Job{{Url: "req2/1", Name: "req2 1 "},
						       {Url: "req2/2", Name: "req2 2 "},
						       {Url: "req2/4", Name: "req2 4 "},
								}},
					{Name: "req3",
					 Error: nil, 
					 Jobs:[]Job{{Url: "req3/1", Name: "req3 1 "},
								{Url: "req3/2", Name: "req3 2 "},
								{Url: "req3/4", Name: "req3 4 "},
								}},
							}
	
	if len(got)!=len(want){
		t.Errorf("Wrong lenght. got %q, wanted %q", len(got),len(want))
	}
						
	for i := range got {
		if got[i].Error!=want[i].Error{
			t.Errorf("Error is not equvalent. got %q, wanted %q", got[i].Error, want[i].Error)
			break
		}
		if got[i].Name!=want[i].Name{
			t.Errorf("Name is not equvalent. got %q, wanted %q", got[i].Name, want[i].Name)
			break
		}

		for j:= range got[i].Jobs{
			if got[i].Jobs[j]!=want[i].Jobs[j]{
				t.Errorf("Wrong Job. got %q, wanted %q", got[i].Jobs[j], want[i].Jobs[j])
				break
			}
		}
								
	}	

}