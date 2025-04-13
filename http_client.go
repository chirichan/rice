package rice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type DefaultResponse map[string]any

func Do[R any](hc *http.Client, request *http.Request) (R, error) {
	var result R

	response, err := hc.Do(request)
	if err != nil {
		return result, err
	}

	defer func(body io.ReadCloser) { _ = body.Close() }(response.Body)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	if len(b) == 0 {
		return result, errors.New("empty response body")
	}

	err = json.Unmarshal(b, &result)
	return result, err
}

func Get[R any](hc *http.Client, url string, params url.Values) (R, error) {
	var result R

	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	return Do[R](hc, req)
}

// PostJson application/json
func PostJson[T, R any](hc *http.Client, url string, data T) (R, error) {
	var result R

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/json")

	return Do[R](hc, req)
}

// PostForm application/x-www-form-urlencoded
func PostForm[R any](hc *http.Client, url string, data url.Values) (R, error) {
	var result R

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return Do[R](hc, req)
}

// PutForm application/x-www-form-urlencoded
func PutForm[R any](hc *http.Client, url string, data url.Values) (R, error) {
	var result R

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(data.Encode()))
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return Do[R](hc, req)
}

// DeleteForm application/x-www-form-urlencoded
func DeleteForm[R any](hc *http.Client, url string, data url.Values) (R, error) {
	var result R

	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(data.Encode()))
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return Do[R](hc, req)
}

type FormFile struct {
	Fieldname string
	Filename  string
	Data      io.Reader
}

// PostMultipartForm multipart/form-data
func PostMultipartForm[R any](hc *http.Client, url string, fieldValues url.Values, files []FormFile) (R, error) {
	var result R

	reqBody := &bytes.Buffer{}

	w := multipart.NewWriter(reqBody)

	for k, v := range fieldValues {
		for _, v2 := range v {
			w.WriteField(k, v2)
		}
	}

	for _, v := range files {
		fw, err := w.CreateFormFile(v.Fieldname, v.Filename)
		if err != nil {
			return result, fmt.Errorf("error create form file %s: %w", v.Fieldname, err)
		}

		_, err = io.Copy(fw, v.Data)
		if err != nil {
			return result, fmt.Errorf("error copying file: %w", err)
		}
	}

	_ = w.Close()

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	return Do[R](hc, req)
}

// PostMultipartFormRaw multipart/form-data
func PostMultipartFormRaw(hc *http.Client, url string, fieldValues url.Values, files []FormFile) (*http.Response, error) {

	reqBody := &bytes.Buffer{}

	w := multipart.NewWriter(reqBody)

	for k, v := range fieldValues {
		for _, v2 := range v {
			w.WriteField(k, v2)
		}
	}

	for _, v := range files {
		fw, err := w.CreateFormFile(v.Fieldname, v.Filename)
		if err != nil {
			_ = w.Close()
			return nil, fmt.Errorf("error create form file %s: %w", v.Fieldname, err)
		}

		_, err = io.Copy(fw, v.Data)
		if err != nil {
			_ = w.Close()
			return nil, fmt.Errorf("error copying file: %w", err)
		}
	}

	_ = w.Close()

	response, err := hc.Post(url, w.FormDataContentType(), reqBody)
	if err != nil {
		return nil, err
	}

	return response, nil
}
