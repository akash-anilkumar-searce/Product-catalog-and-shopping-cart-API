package TestDelete

import (
	"io"
	"net/http"
	"testing"
)

func TestDeleteInventoryNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteinventory/312", nil)
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

	expected := "{\"message\":\"The value product_id does not exist, please enter a Valid ID\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %q, want %q", string(bodyBytes), expected)
	}

}

func TestDeleteInventoryExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteinventory/22", nil)
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

	expected := "\"The value is deleted successfully\"\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %q, want %q", string(bodyBytes), expected)
	}

}
