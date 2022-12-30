package models

import (
	"strconv"
	"sync"
	//"sync"
)



type ParamForOzon func(ozon) ozon

func SetParamCityForOzon(city string) ParamForOzon {
	return func(o ozon) ozon {
		o.link.AddQueryParam("city", city)
		return o
	}
}

func SetParamDepartmentForOzon(department string) ParamForOzon {
	return func(o ozon) ozon {
		o.link.AddQueryParam("department", department)
		return o
	}
}

func SetParamLevelForOzon(level string) ParamForOzon {
	return func(o ozon) ozon {
		o.link.AddQueryParam("level", level)
		return o
	}
}

type ozon struct {
	link *Link
	name string
}

func InitOzon(params ...ParamForOzon) *ozon {

	ozn := ozon{link: &Link{Host:  "job-api.ozon.ru",
							  	   Scheme: "https", 
							  	   Path: "vacancy"},
				name: "Ozon"}

	ozn.link.AddQueryParam("limit", "50")
	for _, param := range params {
		ozn = param(ozn)
	}
	return &ozn
}

func (o *ozon) GetName() string {
	return o.name
}

func (o *ozon) GetJobs() ([]Job, error) {
	or, err := DoRequest[OzonResponse](o.link.GetFullURL()+"&page=1")
	if err != nil {
		return nil, err
	}
	var jobs = or.GetJobs()

	if or.Meta.TotalPages > 1 {
		for jmr := range o.get_json_results(or.Meta.TotalPages){
			if jmr.err != nil {
				return nil, err
			}
			jobs = append(jobs, jmr.jobs...)
		}
	}

	return jobs, nil
}

func (o *ozon) get_json_results(limit int) chan json_model_response {
	out := make(chan json_model_response, limit-1)
	var wg sync.WaitGroup
	wg.Add(limit - 1)
	for i := 2; i <= limit; i++ {
		go func(i int) {
			defer wg.Done()
			or, err := DoRequest[OzonResponse](o.link.GetFullURL()+"&page="+strconv.Itoa(int(i)))
			jmr := json_model_response{}
			if err != nil {
				jmr.err = err
			} else {
				jmr.jobs = or.GetJobs()
			}
			out <- jmr
		}(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

