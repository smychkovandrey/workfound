package models

import (
	"strconv"
)


type meta struct{
	Limit uint8
	Page uint16
	TotalPages int
}

type item struct{
	Hhid int
	Title string
}

type OzonResponse struct{
	Items []item
	Meta meta
}
func (or *OzonResponse) GetJobs () []Job {
	jobs:=make([]Job,0,len(or.Items))
	for _, item:= range or.Items{
		jobs = append(jobs, Job{Url: "https://job.ozon.ru/vacancy/"+ strconv.Itoa(item.Hhid), Name: item.Title})
	}	
	return jobs
}