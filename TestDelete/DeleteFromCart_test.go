package TestDelete

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestDeleteCartNotExists(t *testing.T) {

	data := []byte(`{"product_id":2, "ref":"axxyy"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletefromcart", bytes.NewBuffer(data))
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

	//expected := "{\"message\":\"Product is not found in your cart"}

	/*bodyBytes, err := io.ReadAll(resp.Body)

		//if string(bodyBytes) != expected {
			t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
		}

	}
	*/
}

func TestDeleteCartExists(t *testing.T) {

	data := []byte(`{"product_id":111, "reference_id":"1f45bb50-3f65-423d-b9c9-8daf85b29e3b"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletecart/", bytes.NewBuffer(data))
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
	expected := ""

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
