package models


type resultLamoda struct{
	Id int
	Title_ru string
	Slug string
}

type LamodaResponse struct{
	Results []resultLamoda
}
func (la *LamodaResponse) GetJobs () []Job {
	var jobs []Job
	for _, item:= range la.Results{
		jobs = append(jobs, Job{Url: "https://job.lamoda.ru/vacancies/"+ item.Slug, Name: item.Title_ru})
	}	
	return jobs
}