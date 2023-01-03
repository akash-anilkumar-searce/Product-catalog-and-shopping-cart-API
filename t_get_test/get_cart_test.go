package t_get_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCartExists(t *testing.T) {
	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Get("http://localhost:8089/cart/get?ref=2468bb80-74cf-4af7-8096-f3a465540fb2")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	// Read the response from the API
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}

func TestGetCartNotExists(t *testing.T) {
	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Get("http://localhost:8089/cart/get?ref=2468bb80-74cf-4af7-80qwe96-f3a465540f")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)
	var response string

	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response != "NO DATA FOUND!" {
		t.Errorf("Expected Reference_id doesn't exists, got %s", response)
	}

}
