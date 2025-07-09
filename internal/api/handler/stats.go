package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mu            sync.Mutex
	requestCounts = make(map[FizzBuzzRequest]int)
)

func recordStats(req FizzBuzzRequest) {
	mu.Lock()
	defer mu.Unlock()
	requestCounts[req]++
}

type statsResponse struct {
	Parameters FizzBuzzRequest `json:"parameters"`
	Hits       int             `json:"hits"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var mostUsed FizzBuzzRequest
	maxHits := 0
	found := false

	for req, hits := range requestCounts {
		if hits > maxHits {
			mostUsed = req
			maxHits = hits
			found = true
		}
	}

	if !found {
		http.Error(w, "no stats available", http.StatusNotFound)
		return
	}

	response := statsResponse{
		Parameters: mostUsed,
		Hits:       maxHits,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
