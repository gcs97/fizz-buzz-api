package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
	t.Run("Valid request", func(t *testing.T) {
		reqURL := "/?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
		req := httptest.NewRequest(http.MethodGet, reqURL, nil)
		rr := httptest.NewRecorder()

		FizzBuzzHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}

		var resp FizzBuzzResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

		if len(resp.Result) != len(expected) {
			t.Fatalf("Expected result length %d, got %d", len(expected), len(resp.Result))
		}

		for i, v := range expected {
			if resp.Result[i] != v {
				t.Errorf("Expected result[%d] = %s, got %s", i, v, resp.Result[i])
			}
		}

		contentType := rr.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}
	})

	t.Run("Invalid request", func(t *testing.T) {
		reqURL := "/?int1=abc&int2=5&limit=15&str1=fizz&str2=buzz"
		req := httptest.NewRequest(http.MethodGet, reqURL, nil)
		rr := httptest.NewRecorder()

		FizzBuzzHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", rr.Code)
		}

		var resp FizzBuzzResponse
		json.NewDecoder(rr.Body).Decode(&resp)

		if len(resp.Errors) == 0 {
			t.Error("Expected validation errors, got none")
		}
	})

	t.Run("Invalid parameters", func(t *testing.T) {
		reqURL := "/?int1=3&int2=5"
		req := httptest.NewRequest(http.MethodGet, reqURL, nil)
		rr := httptest.NewRecorder()

		FizzBuzzHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", rr.Code)
		}

		var resp FizzBuzzResponse
		json.NewDecoder(rr.Body).Decode(&resp)

		expectedErrors := []string{
			"limit must be a positive integer",
			"str1 must be a non-empty string",
			"str2 must be a non-empty string",
		}

		for _, expected := range expectedErrors {
			found := false
			for _, actual := range resp.Errors {
				if strings.Contains(actual, expected) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error containing '%s' not found", expected)
			}
		}
	})
}
