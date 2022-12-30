package models

import (
	"WorkFound/internal/fileworker"
	"log"
	"strings"
	"sync"
)

type JobRequester interface{
	GetName() string
	GetJobs() ([]Job, error)
}

type Filter func(Job) bool

type Results struct {
	Name string
	Jobs []Job
	Error error
}

type search_Engine struct{
	Filter Filter
	Requsters []JobRequester
}

func (se *search_Engine) GetResults() <-chan Results {
	out:=make(chan Results, len(se.Requsters))
	var wg sync.WaitGroup
	wg.Add(len(se.Requsters))
	for i, ijr := range se.Requsters {
		go func (i int, ijr JobRequester){
			defer wg.Done()
			results:= Results{Name: ijr.GetName()}
			results.Jobs, results.Error=ijr.GetJobs()
			if se.Filter != nil{
				results.Jobs=se.filter(results.Jobs)
			}
			out<-results
		}(i,ijr)
	}
	go func(){
		wg.Wait()
		close (out)
	}()
	
	return out
}
func (se *search_Engine) filter (jobs []Job) (ret []Job){
	for _, j := range jobs {
		if se.Filter(j){
			ret = append(ret, j)
		}
	}
	return
}
func CreateSeacrhEngine () (*search_Engine){
	ya := InitYandex(SetParamLevelForYandex("chief"), SetParamCityForYandex("moscow"),
		SetParamAreaForYandex(YaDB),
		SetParamAreaForYandex(YaDesktop),
		SetParamAreaForYandex(YaFrontend),
		SetParamAreaForYandex(YaBackend),
		SetParamAreaForYandex(YaFullStack),
		SetParamAreaForYandex(YaML),
		SetParamAreaForYandex(YaMobile),
		SetParamAreaForYandex(YaMobileAndroid),
		SetParamAreaForYandex(YaMobileIOS),
		SetParamAreaForYandex(YaNOC),
		SetParamAreaForYandex(YaSystem),
		SetParamAreaForYandex(YaDevOps),
		//SetParamAreaForYandex(YaTechManager),
	)
	o := InitOzon( SetParamCityForOzon("Москва"),
		SetParamDepartmentForOzon("Ozon Информационные технологии"),
		SetParamDepartmentForOzon("Ozon Fintech"),
		SetParamLevelForOzon("Руководитель"))

	st := InitSberTech(SetParamCityForSberTech(STMoscow),
		SetParamSearchStringForSberTechWithReplace("Руководитель"),
		SetParamAreaForSberTech(STDev),
		SetParamAreaForSberTech(STDevArch),
		SetParamAreaForSberTech(STSupportAndDevOps),
	)

	// la := InitLamoda(SetParamCityForLamoda(Moscow),
	// 	SetParamTitleForLamodaWithReplace("manager"))
	stopwords,err:=fileworker.GetStopWords()
	if err!=nil{
		log.Println(err)
	}
	var se = search_Engine{
		Filter: func(j Job) bool {
			for _, str := range stopwords {
				if strings.Contains(strings.ToLower(j.Name), str) {
					return false
				}
			}
			return true
		},
		Requsters: []JobRequester{
			o,
			ya,
			st,
			//la,
		},
	}
	return &se;
}