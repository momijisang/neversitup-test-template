package persistence

import (
	"io"
	"net/http"
	"strings"

	"errors"
)

type AnotherAPIRepository struct{}

var anotherAPIRepository *AnotherAPIRepository

func AnotherAPI() *AnotherAPIRepository {
	if anotherAPIRepository == nil {
		anotherAPIRepository = &AnotherAPIRepository{}
	}
	return anotherAPIRepository
}

func (r AnotherAPIRepository) CallHttp(url string, method string, data string) (string, error) {
	if method == "get" {
		return r.HttpGet(url)
	} else if method == "post" {
		return r.HttpPostRawJson(url, data)
	}
	return "", errors.New("method not found")
}

func (r AnotherAPIRepository) HttpGet(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		Log().AddApiLog(url, 0, "", err.Error(), false)
		return "", err
	}
	if (res.StatusCode < 200) || (res.StatusCode >= 300) {
		Log().AddApiLog(url, res.StatusCode, "", res.Status, false)
		return "", errors.New(res.Status)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		Log().AddApiLog(url, res.StatusCode, "", res.Status, false)
		return "", err
	}
	Log().AddApiLog(url, res.StatusCode, "", string(body), true)
	return string(body), nil
}

func (r AnotherAPIRepository) HttpPostRawJson(url string, data string) (string, error) {
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		Log().AddApiLog(url, 0, data, err.Error(), false)
		return "", err
	}
	if (res.StatusCode < 200) || (res.StatusCode >= 300) {
		Log().AddApiLog(url, res.StatusCode, data, res.Status, false)
		return "", errors.New(res.Status)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		Log().AddApiLog(url, res.StatusCode, data, res.Status, false)
		return "", err
	}
	Log().AddApiLog(url, res.StatusCode, data, string(body), true)
	return string(body), nil
}
