package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Home(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading the response body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestAbout(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/about", nil)
	w := httptest.NewRecorder()

	About(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading the response body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}