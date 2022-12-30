package models

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)
type CodeLocationForLamoda string
const(
	Moscow CodeLocationForLamoda = "1"
)
type ParamForLamoda func(lamoda) lamoda

func SetParamCityForLamoda(city CodeLocationForLamoda) ParamForLamoda {
	return func(la lamoda) lamoda {
		cities,res:=la.link.GetQueryParam("cities")
		if res{
			la.link.SetQueryParam("cities",cities+","+string(city))
		} else{
			la.link.AddQueryParam("cities",string(city))
		}
		return la
	}
}
func SetParamTitleForLamodaWithReplace(title string) ParamForLamoda {
	return func(la lamoda) lamoda {
		la.link.SetQueryParam("searchString",title)
		return la
	}
}

type lamoda struct {
	link *Link
	name string
}

func InitLamoda(params...ParamForLamoda) *lamoda {

	la := lamoda{ link: &Link{Host:  "job.lamoda.ru",
	 								 Scheme: "https", 
	 								 Path: "api/vacancies"},
					name: "Lamoda"}
	la.link.AddQueryParam("status","1") //открытая вакансия
	la.link.AddQueryParam("category_id","15") //подразделение разработки
	for _, param := range params {
		la = param(la)
	}	
	return &la
}

func (la *lamoda) GetName() string {
	return la.name
}

func (la *lamoda) GetJobs() ([]Job, error) {
	yar, err := la.jsonconvert()
	if err != nil {
		return nil, err
	}
	var jobs = yar.GetJobs()

	return jobs, nil
}
func (la *lamoda) jsonconvert() (*LamodaResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, la.link.GetFullURL(), nil)
	req.Header.Add("Referer","https://job.lamoda.ru/vacancies")
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if  len(body)<14 { //пустой ответ у ламоды "[]" + названия поле
		return nil, errors.New("пустой ответ "+ string(body))
	}

	//убираем квадратые скобки
	jsn:=make([]byte,len(body)-2)
	for i := 1; i < len(body)-2; i++ {
		jsn[i-1]=body[i]
	}

	return JsonUnmarshal[LamodaResponse](jsn)
}