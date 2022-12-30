package models

import (
	"strconv"
)

type CodeAreaSberTech string
type vacancie struct {
	InternalId       int
	Title            string
	SpecializationId CodeAreaSberTech
}
type data struct {
	Vacancies []vacancie
	Total     int
}
type SberTechResponse struct {
	Data    data
	Success bool
}

func (str *SberTechResponse) GetJobs(filter_areas []CodeAreaSberTech) []Job {
	jobs := make([]Job, 0, len(str.Data.Vacancies))
	if filter_areas == nil || len(filter_areas) == 0 {
		for _, item := range str.Data.Vacancies {
			jobs = append(jobs, Job{Url: "https://rabota.sber.ru/search/" + strconv.Itoa(item.InternalId), Name: item.Title})
		}
	} else {
		for _, item := range str.Data.Vacancies {
			for _, area := range filter_areas {
				if item.SpecializationId == area {
					jobs = append(jobs, Job{Url: "https://rabota.sber.ru/search/" + strconv.Itoa(item.InternalId), Name: item.Title})
					break
				}
			}
		}
	}
	return jobs
}
