package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	FormData = "application/x-www-form-urlencoded"
	JsonData = "application/json"
)

type RequestData struct {
	Query       map[string]string `json:"query"`
	BodyData    map[string]string `json:"body_data"`
	Header      map[string]string `json:"header"`
	ContentType string            `json:"content_type"`
}

type Response struct {
	BasicResponse
	Data map[string]interface{}
}

type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendHTTPRequest(requestURL, method string, data *RequestData) ([]byte, error) {
	if len(data.Query) != 0 {
		requestURL += "?"
	}
	for key, value := range data.Query {
		requestURL += fmt.Sprintf("%s=%s&", key, value)
	}

	var payload string

	if len(data.BodyData) != 0 {
		if data.ContentType == JsonData {
			body, err := json.Marshal(data.BodyData)
			if err != nil {
				return nil, err
			}
			payload = string(body)
		} else if data.ContentType == FormData {
			body := url.Values{}
			for key, value := range data.BodyData {
				body.Set(key, value)
			}
			payload = body.Encode()
		}
	}

	req, err := http.NewRequest(method, requestURL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", data.ContentType)

	for key, value := range data.Header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func MarshalBody(body []byte) (map[string]interface{}, error) {
	var rp Response
	if err := json.Unmarshal(body, &rp); err != nil {
		return nil, err
	}
	if rp.Code != 0 {
		return nil, errors.New(rp.Message)
	}
	return rp.Data, nil
}

func MarshalBodyForCustomData(body []byte, rp interface{}) error {
	return json.Unmarshal(body, &rp)
}
