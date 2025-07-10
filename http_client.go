package rice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var defaultHTTPClient = &http.Client{Timeout: 30 * time.Second}

func SetHttpClient(client *http.Client) {
	defaultHTTPClient = client
}

// HeaderXRequestID is the header key for request ID propagation.
const HeaderXRequestID = "X-Request-ID"

// GetRequestID extracts the request ID from context or returns a default value.
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return NextUUID()
	}
	if v := ctx.Value("request_id"); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return NextUUID()
}

type DefaultResponse map[string]any

type RequestOption func(request *http.Request)

func WithAuthorizationHeader(token string) RequestOption {
	return func(request *http.Request) {
		if token == "" {
			return
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

func WithHeader(key, value string) RequestOption {
	return func(request *http.Request) {
		if key == "" || value == "" {
			return
		}
		request.Header.Add(key, value)
	}
}

func Post[T, R any](ctx context.Context, url string, req T, opts ...RequestOption) (R, error) {

	requestID := GetRequestID(ctx)

	var res R
	postbody, err := json.Marshal(req)
	if err != nil {
		return res, err
	}

	buf := bytes.NewBuffer(postbody)

	request, err := http.NewRequestWithContext(ctx, "POST", url, buf)
	if err != nil {
		return res, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Add(HeaderXRequestID, requestID)
	for _, opt := range opts {
		opt(request)
	}

	resp, err := defaultHTTPClient.Do(request)
	if err != nil {
		return res, fmt.Errorf("post url: %s, params: %v, err: %w", url, req, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read reponse body err: %w, url: %s", err, url)
	}

	slog.DebugContext(ctx, "http post", "url", url, "requset", string(postbody), "response", string(respBody))

	var r R
	if err := json.Unmarshal(respBody, &r); err != nil {
		return res, fmt.Errorf("unmarshal response body: %s, err: %w", string(respBody), err)
	}

	return r, nil
}

func Get[R any](ctx context.Context, url string, params url.Values, opts ...RequestOption) (R, error) {

	var res R

	reqURL := url

	if len(params) > 0 {
		reqURL = url + "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return res, err
	}

	request.Header.Add(HeaderXRequestID, GetRequestID(ctx))
	for _, opt := range opts {
		opt(request)
	}

	resp, err := defaultHTTPClient.Do(request)
	if err != nil {
		return res, fmt.Errorf("get url: %s, params: %v, err: %w", url, params, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read reponse body err: %w", err)
	}

	// slog.InfoContext(ctx, "http get", "url", reqURL, "response", string(respBody))

	var r R
	if err := json.Unmarshal(respBody, &r); err != nil {
		return res, fmt.Errorf("unmarshal response body: %s, err: %w", string(respBody), err)
	}

	return r, nil
}

func Do[R any](hc *http.Client, request *http.Request) (R, error) {
	var result R

	response, err := hc.Do(request)
	if err != nil {
		return result, err
	}

	if response.StatusCode != http.StatusOK {
		return result, fmt.Errorf("http status code %d, status: %s", response.StatusCode, response.Status)
	}

	defer func(body io.ReadCloser) { _ = body.Close() }(response.Body)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	if len(b) == 0 {
		return result, errors.New("empty response body")
	}

	slog.Debug("response body", "body", string(b))

	err = json.Unmarshal(b, &result)
	return result, err
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

func DownloadFile(ctx context.Context, url string, params url.Values, saveDir, filename string, opts ...RequestOption) error {
	reqURL := url

	if len(params) > 0 {
		reqURL = url + "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return err
	}

	request.Header.Add(HeaderXRequestID, GetRequestID(ctx))
	for _, opt := range opts {
		opt(request)
	}

	resp, err := defaultHTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("get url: %s, params: %v, err: %w", url, params, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("create directory failed: %v", err)
	}

	filePath := filepath.Join(saveDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("save file failed: %v", err)
	}

	return nil
}
