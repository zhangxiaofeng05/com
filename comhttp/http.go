// Package comhttp provide function about http standard library
package comhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get get method. return json data
func Get(ctx context.Context, url string, header map[string]string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	errorBuf, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("HTTP GET url:%s response status: %v, err:%v", url, resp.Status, errorBuf)
}

// Post post method. return json data
func Post(ctx context.Context, url string, header map[string]string, data []byte, result any) error {
	body := bytes.NewBuffer(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	errorBuf, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("HTTP POST url:%s response status: %v, err:%v", url, resp.Status, errorBuf)
}
