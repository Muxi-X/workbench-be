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
	Query       map[string]string
	BodyData    map[string]string
	ContentType string
}

type Response struct {
	Code    int
	Message string
	Data    map[string]interface{}
}

func SendHTTPRequest(requestURL, method string, data *RequestData) ([]byte, error) {
	if len(data.Query) != 0 {
		requestURL += "?"
	}
	for key, value := range data.Query {
		requestURL += fmt.Sprintf("%s=%s&", key, value)
	}
	fmt.Println(requestURL)

	var payload string

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
	} else {
		return nil, errors.New("content type error")
	}

	fmt.Println(payload)

	req, err := http.NewRequest(method, requestURL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", data.ContentType)

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

	fmt.Println(string(body))

	return body, nil
}

func MarshalBody(body []byte) (map[string]interface{}, error) {
	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func MarshalBodyForCustomData(body []byte, data interface{}) error {
	return json.Unmarshal(body, &data)
}
