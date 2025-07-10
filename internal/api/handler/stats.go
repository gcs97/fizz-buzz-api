package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mu            sync.Mutex
	requestCounts = make(map[fizzBuzzRequest]int)
)

func recordStats(req fizzBuzzRequest) {
	mu.Lock()
	defer mu.Unlock()
	requestCounts[req]++
}

type statsResponse struct {
	Parameters fizzBuzzRequest `json:"parameters,omitempty"`
	Hits       int             `json:"hits,omitempty"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var mostUsed fizzBuzzRequest
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
		w.WriteHeader(http.StatusNotFound)
		errMsg := map[string]string{"error": "no stats available"}
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	response := statsResponse{
		Parameters: mostUsed,
		Hits:       maxHits,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
