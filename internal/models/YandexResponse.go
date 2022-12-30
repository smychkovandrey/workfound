package models

import (
	"strconv"
)

type result struct{
	Id int
	Title string
}

type YandexResponse struct{
	Results []result
}
func (yr *YandexResponse) GetJobs () []Job {
	jobs:=make([]Job,0,len(yr.Results))
	for _, item:= range yr.Results{
		jobs = append(jobs, Job{Url: "https://yandex.ru/jobs/vacancies/"+ strconv.Itoa(item.Id), Name: item.Title})
	}	
	return jobs
}
