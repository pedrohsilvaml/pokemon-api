package request

import (
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	Body       []byte
	StatusCode int
}

func Get(url string) (*HttpResponse, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{Body: body, StatusCode: req.StatusCode}, nil
}
