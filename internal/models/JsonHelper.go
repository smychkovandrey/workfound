package models

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)
type json_model_response struct {
	jobs []Job
	err  error
}
type Response interface {
	SberTechResponse | OzonResponse | YandexResponse | LamodaResponse
   }
func DoRequest[T Response](url string) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
	return JsonUnmarshal[T](body)
}

func JsonUnmarshal[T Response](body []byte) (*T, error) {
	var or T
	
	err := json.Unmarshal(body, &or)
	if err != nil {
		return nil, err
	}
	return &or, nil
}