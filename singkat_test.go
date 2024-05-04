package main

import "testing"

func TestNewShorten(t *testing.T) {

	urls := []struct {
		url string
	}{
		{url: "https://www.google.com/"},
		{url: "https://www.google.com/?name=fairus"},
		{url: "https://github.com/jackc/pgx"},
	}

	for _, test := range urls {
		t.Run("Valid URL: "+test.url, func(t *testing.T) {
			newUrl, err := NewShorten(test.url)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if newUrl == nil {
				t.Error("Unexpected error: New URL not generated")
				return
			}

			if newUrl.URL != test.url {
				t.Error("Unexpected error: New URL nil")
				return
			}
		})
	}

	invalidURLs := []struct {
		url string
	}{
		{url: "this-not-url"},
		{url: "this not url"},
	}

	for _, test := range invalidURLs {
		t.Run("Invalid URL: "+test.url, func(t *testing.T) {
			_, err := NewShorten(test.url)

			if err == nil {
				t.Errorf("Expected error for invalid URL, but no error occurred")
				return
			}

			if err != ErrUrl {
				t.Errorf("Expected error for invalid URL, but false error")
				return
			}

		})
	}

}
