package mz_request

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
GetJson sends an HTTP GET request to the specified URL and unmarshals the response body into the provided data.

Parameters:

- url: the URL to send the GET request to.

- data: a pointer to the variable where the response body will be unmarshaled into.

Returns:

- error: an error if the GET request fails or if there is an issue unmarshaling the response body.
*/
func GetJson(url string, data any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}
