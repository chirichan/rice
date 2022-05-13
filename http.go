package rice

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func HttpGet[T any](url string) (T, error) {

	var (
		data T
		err  error
	)

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)

	return data, err
}

func UnWrapperURL(r *http.Request) map[string]any {

	var urlValues = make(map[string]any)

	if err := r.ParseForm(); err != nil {
		log.Printf("err: %+v", err)
		return nil
	}

	for k, v := range r.Form {
		if len(v) == 1 {
			urlValues[k] = v[0]
		}
	}

	return urlValues
}

func UnWrapperBody[T any](r *http.Request) (*T, error) {

	var (
		data T
		err  error
	)

	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {
			log.Printf("UnWrapperBody-err: %+v", err)
			return
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	return &data, err
}

func ErrorWrapper(w http.ResponseWriter, code int, err error, msg ...string) error {

	w.Header().Set("Content-Type", "application/json")

	log.Printf("err: %+v", err)

	d := ApiWrapper{
		Code: code,
	}

	if len(msg) > 0 {
		d.Msg = strings.Join(msg, ",")
	} else {
		d.Msg = err.Error()
	}

	b, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))

	return err
}

func Wrapper(w http.ResponseWriter, data any, msg ...string) error {

	w.Header().Set("Content-Type", "application/json")

	d := ApiWrapper{
		Code: 0,
		Data: data,
	}

	if len(msg) > 0 {
		d.Msg = strings.Join(msg, ",")
	}

	b, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))

	return err
}

type ApiWrapper struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
