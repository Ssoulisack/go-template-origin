package outsource

import (
	"bytes"
	"fmt"
	"go-fiber/core/logs"
	"io/ioutil"
	"net/http"
	"time"
)

var HttpClient = http.Client{
	Timeout: 5 * time.Minute,
}

func Do(log_key string, req *http.Request) (*http.Response, error) {
	logs.Info(fmt.Sprintf("key:%v Method %v : %v", log_key, req.Method, req.URL))
	// Log the request payload
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logs.Info(fmt.Sprintf("Error reading request body: %v", err))
		} else {
			logs.Info(fmt.Sprintf("key:%v Request: %s", log_key, body))
			// Reset the request body for subsequent use
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
	req.Close = true
	// Send the request and get the response
	resp, err := HttpClient.Do(req)

	// Log the response for non-GET and non-OPTIONS requests
	if err != nil {
		logs.Info(fmt.Sprintf("Error sending request: %v", err))
	} else if req.Method != "GET" && req.Method != "OPTIONS" {
		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Info(fmt.Sprintf("Error reading response body: %v", err))
		} else {
			logs.Info(fmt.Sprintf("key:%v Response: %s", log_key, responseBody))
			// Reset the response body for subsequent use
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
		}
	}

	return resp, err
}
