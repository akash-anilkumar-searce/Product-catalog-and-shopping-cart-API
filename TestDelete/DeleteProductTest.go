package TestDelete

import (
	"io"
	"net/http"
	"testing"
)

func TestDeleteProductNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteproduct/121", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := "The value of Product_Id Does not exist,enter a Valid ID"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteProductExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteproduct/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Check the response body, if necessary
	// ...

	expected := ""

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
