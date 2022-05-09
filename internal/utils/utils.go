package utils

import (
	"io/ioutil"
	"fmt"
	"os"
	"bytes"
	"math/rand"
	"time"
	"strconv"
	"net/http"
)

const (
	USE_BINDING      = false
	HTTP_URL         = "https://jsonplaceholder.typicode.com"
	HTTP_BINDING_URL = "http://127.0.0.1:3500/v1.0/bindings/jsonbinding"
	QUERY_TEMPLATE   = `{
		"operation": "get",
		"metadata": {
			"path": "/%s"
		}
	}`
)

var client = &http.Client{}
var PostType = []string{"story", "job", "poll"}

func Request(path string) ([]byte, string, error) {
	var (
		url       string
		err       error
		reqMethod string
		reqData   bytes.Buffer
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	useBinding, err := strconv.ParseBool(os.Getenv("USE_BINDING"))
	if err != nil {
		return nil, "", err
	}
	switch useBinding {
	case false:
		url = os.Getenv("HTTP_URL")
		if url == "" {
			url = HTTP_URL
		}
		url = url + "/" + path
		reqMethod = "GET"
	default:
		url = os.Getenv("HTTP_BINDING_URL")
		if url == "" {
			url = HTTP_BINDING_URL
		}
		reqMethod = "POST"
		q := fmt.Sprintf(QUERY_TEMPLATE, path)
		reqData = *bytes.NewBuffer([]byte(q))
	}
	if req, err := http.NewRequest(reqMethod, url, &reqData); err == nil {
		req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("dapr-app-id", "gql")
		if resp, err := client.Do(req); err == nil {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				return body, url, nil
			}
		}
	}
	return nil, "", err
}