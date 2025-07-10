package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func resetStats() {
	mu.Lock()
	defer mu.Unlock()
	requestCounts = make(map[fizzBuzzRequest]int)
}

func TestStatsHandler(t *testing.T) {
	t.Run("Valid request", func(t *testing.T) {
		resetStats()

		fbReq := fizzBuzzRequest{
			Int1:  3,
			Int2:  5,
			Limit: 15,
			Str1:  "fizz",
			Str2:  "buzz",
		}

		for range 5 {
			recordStats(fbReq)
		}

		req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", nil)
		rr := httptest.NewRecorder()

		StatsHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", rr.Code)
		}

		var resp statsResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if resp.Hits != 5 {
			t.Errorf("Expected hits to be 5, got %d", resp.Hits)
		}

		if resp.Parameters != fbReq {
			t.Errorf("Expected parameters %+v, got %+v", fbReq, resp.Parameters)
		}
	})

	t.Run("No stats available", func(t *testing.T) {
		resetStats()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", nil)
		rr := httptest.NewRecorder()

		StatsHandler(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected 404 Not Found, got %d", rr.Code)
		}
	})
}
