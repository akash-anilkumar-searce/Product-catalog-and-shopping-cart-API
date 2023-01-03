package t_get_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetProductExists(t *testing.T) {
	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Get("http://localhost:8089/product/1")
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

func TestGetProductNotExists(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/product/123")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var response string

	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response != "Product does not exist" {
		t.Errorf("Product does not exist, got %s", response)
	}

}
