package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// using fake STORE API as dummy database for now at least
func Request[T any](method, path string, data interface{}) (*T, error) {
	var client = &http.Client{}
	BaseURL := "https://fakestoreapi.com/"
	url := BaseURL + path
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(err)
		return nil, readErr
	}

	var resp T
	json.Unmarshal(body, &resp)
	return &resp, nil
}
