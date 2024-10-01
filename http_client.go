package rice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
)

type DefaultResponse map[string]any

var defaultHttpClient = &HttpClient{client: http.DefaultClient, logger: slog.Default()}

type HttpClient struct {
	client *http.Client
	logger *slog.Logger
}

func DefaultHttpClient() *HttpClient {
	return defaultHttpClient
}

func NewHttpClient(client *http.Client, logger *slog.Logger) *HttpClient {
	return &HttpClient{client: client, logger: logger}
}

func Do[R any](hc *HttpClient, request *http.Request) (R, error) {
	var result R

	response, err := hc.client.Do(request)
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

func Get[R any](hc *HttpClient, url string, params url.Values) (R, error) {
	var result R

	if len(params) > 0 {
		url += url + "?" + params.Encode()
	}

	response, err := hc.client.Get(url)
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

// PostJson application/json
func PostJson[T, R any](hc *HttpClient, url string, data T) (R, error) {
	var result R

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	response, err := hc.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
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

// PostForm application/x-www-form-urlencoded
func PostForm[R any](hc *HttpClient, url string, data url.Values) (R, error) {
	var result R

	response, err := hc.client.PostForm(url, data)
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

type FormFile struct {
	Fieldname string
	Filename  string
	Data      io.Reader
}

// PostMultipartForm multipart/form-data
func PostMultipartForm[R any](hc *HttpClient, url string, fieldValues url.Values, files []FormFile) (R, error) {
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

	response, err := hc.client.Post(url, w.FormDataContentType(), reqBody)
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
