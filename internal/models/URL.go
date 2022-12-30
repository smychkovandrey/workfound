package models

import "net/url"

type Link struct {
	Host string
	Scheme string
	Path string
	queryUrl   *url.Values
}
func (link *Link)AddQueryParam(key string, value string){
	if link.queryUrl== nil{
		link.queryUrl=&url.Values{}
	}
	link.queryUrl.Add(key,value)
}
func (link *Link)GetQueryParam(key string) (string,bool){
	if link.queryUrl== nil || !link.queryUrl.Has(key){
		return "", false
	}
	return link.queryUrl.Get(key),true
}
func (link *Link)DeleteQueryParam(key string){
	if link.queryUrl== nil || !link.queryUrl.Has(key){
		return
	}
	link.queryUrl.Del(key)
}
func (link *Link)SetQueryParam(key string, value string){
	if link.queryUrl== nil{
		link.queryUrl=&url.Values{}
	}
	link.queryUrl.Set(key, value)
}
func (link *Link) GetFullURL() string{

	url := url.URL{
		Host:       link.Host,
		Scheme: 	link.Scheme,
		Path:       link.Path,
	}
	if link.queryUrl!=nil {
		url.RawQuery= link.queryUrl.Encode()
	}
	return url.String()
}