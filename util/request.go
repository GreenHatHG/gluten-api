package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Post(url string, body string) (content []byte, err error) {
	req, errHttp := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if errHttp != nil {
		return
	}

	client := http.Client{}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	res, errClient := client.Do(req)
	if errClient != nil {
		return
	}

	defer res.Body.Close()
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

func Get(url string, token string) (content []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
