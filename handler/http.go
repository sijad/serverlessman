package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func doHTTP(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func newReq(url, method string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func getJSON(url string, headers map[string]string) (*http.Response, error) {
	req, err := newReq(url, http.MethodGet, headers, nil)
	if err != nil {
		return nil, err
	}

	res, err := doHTTP(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func sendJSON(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := newReq(url, method, headers, b)
	if err != nil {
		return nil, err
	}

	res, err := doHTTP(req)
	if err != nil {
		return nil, err
	}

	// if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
	// 	return errors.New("send json error: " + res.Status + " status code")
	// }

	return res, nil
}
