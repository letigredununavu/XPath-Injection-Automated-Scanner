package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Send a GET request and returns the response body as string
func SendRequest(target string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(target)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	// Once we are done reading the body, close automatically the HTTP response
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(bodyBytes), nil
}

func SendPostRequest(target string, payload string, contentType string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	body := bytes.NewBufferString(payload)
	resp, err := client.Post(target, contentType, body)

	if err != nil {
		return "", fmt.Errorf("failed to send POST request: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read reponse body: %v", err)
	}

	return string(bodyBytes), nil

}
