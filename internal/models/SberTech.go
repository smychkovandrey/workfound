package models

import (
	"strconv"
	"sync"
)

type CodeLocationSberTech string

const(
	STMoscow CodeLocationSberTech ="0c5b2444-70a0-4932-980c-b4dc0d3f02b5"
	STDev CodeAreaSberTech = "39d29fd5-1a53-43ae-ab6f-1eda9ce5db58"
	STDevArch CodeAreaSberTech = "9b07fbb7-3259-4eb5-839c-ea14b08ad51a"
	STSupportAndDevOps CodeAreaSberTech = "21b7434a-43e7-4841-920d-29433b91e19d"

)

type sbertech struct {
	link *Link
	name string
	take int
	areas []CodeAreaSberTech
}

type ParamForSberTech func(sbertech) sbertech

func SetParamSearchStringForSberTechWithReplace(searchstring string) ParamForSberTech {
	return func(st sbertech) sbertech {
		st.link.SetQueryParam("searchString",searchstring)
		return st
	}
}
func SetParamCityForSberTech(city CodeLocationSberTech) ParamForSberTech {
	return func(st sbertech) sbertech {
		st.link.AddQueryParam("locations", string(city))
		return st
	}
}
func SetParamAreaForSberTech(area CodeAreaSberTech) ParamForSberTech {
	return func(st sbertech) sbertech {
		if st.areas ==nil{
			st.areas =make([]CodeAreaSberTech, 1);
			st.areas[0] = area
		} 
		st.areas = append(st.areas, area)
		return st
	}
}
func InitSberTech(params ...ParamForSberTech) *sbertech {

	st := sbertech{link: &Link{Host:  "rabota.sber.ru",
							  	   Scheme: "https", 
							  	   Path: "public/app-candidate-public-api-gateway/api/v1/publications"},
				name: "SberTech",
			take:50}

	st.link.SetQueryParam("take",strconv.Itoa(int(st.take)))
	for _, param := range params {
		st = param(st)
	}
	return &st
}

func (st *sbertech) GetName() string {
	return st.name
}

func (st *sbertech) GetJobs() ([]Job, error) {
	str, err := DoRequest[SberTechResponse](st.link.GetFullURL())
	if err != nil {
		return nil, err
	}
	var jobs = str.GetJobs(st.areas)

	if str.Data.Total > st.take {
		for jmr := range st.get_json_results(str.Data.Total) {
			if jmr.err != nil {
				return nil, err
			}
			jobs = append(jobs, jmr.jobs...)
		}
	}

	return jobs, nil
}

func (st *sbertech) get_json_results(limit int) <-chan json_model_response {
	out := make(chan json_model_response, limit/st.take)
	var wg sync.WaitGroup
	wg.Add(limit/st.take)
	for i := 1; i <= limit/st.take; i++ {
		go func(i int) {
			defer wg.Done()
			str, err := DoRequest[SberTechResponse](st.link.GetFullURL()+"&skip="+strconv.Itoa(int(st.take*i)))
			jmr := json_model_response{}
			if err != nil {
				jmr.err = err
			} else {
				jmr.jobs = str.GetJobs(st.areas)
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
