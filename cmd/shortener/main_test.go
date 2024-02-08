package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRedirect(t *testing.T) {
	shortener := idToURLMap{
		links: make(map[string]string),
	}
	shortener.id = "123"
	shortener.links[shortener.id] = "https://practicum.yandex.ru/"
	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortener.id)
	req, err := http.NewRequest("GET", shortenedURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.handleRedirect)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

	expectedLocation := shortener.links[shortener.id]
	if location := rr.Header().Get("Location"); location != expectedLocation {
		t.Errorf("handler returned unexpected location header: got %v want %v",
			location, expectedLocation)
	}
}

func TestHandleShortenURL(t *testing.T) {
	shortener := idToURLMap{
		links: make(map[string]string),
	}
	originalURL := "https://practicum.yandex.ru/"
	body := strings.NewReader("https://practicum.yandex.ru/")
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}
	shortener.id = "123"
	rr := httptest.NewRecorder()
	shortener.handleShortenURL(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}

	expectedContentType := "text/plain"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
			contentType, expectedContentType)
	}

	expectedURL := fmt.Sprintf("http://localhost:8080/%s", shortener.id)
	bodyBytes := rr.Body.Bytes()
	if string(bodyBytes) != expectedURL {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyBytes), expectedURL)
	}

	if url := shortener.links[shortener.id]; url != originalURL {
		t.Errorf("handler failed to add URL to map: got %v want %v",
			url, originalURL)
	}
}
