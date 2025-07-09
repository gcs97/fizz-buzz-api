package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type FizzBuzzRequest struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func (r *FizzBuzzRequest) Validate() []string {
	var errors []string

	if r.Int1 == 0 {
		errors = append(errors, "int1 must be a valid non-zero integer")
	}
	if r.Int2 == 0 {
		errors = append(errors, "int2 must be a valid non-zero integer")
	}
	if r.Limit <= 0 {
		errors = append(errors, "limit must be a positive integer")
	}
	if r.Str1 == "" {
		errors = append(errors, "str1 must be a non-empty string")
	}
	if r.Str2 == "" {
		errors = append(errors, "str2 must be a non-empty string")
	}

	return errors
}

func (r *FizzBuzzRequest) Compute() []string {
	result := make([]string, 0, r.Limit)
	for i := 1; i <= r.Limit; i++ {
		val := ""
		if i%r.Int1 == 0 {
			val += r.Str1
		}
		if i%r.Int2 == 0 {
			val += r.Str2
		}
		if val == "" {
			val = strconv.Itoa(i)
		}
		result = append(result, val)
	}
	return result
}

type FizzBuzzResponse struct {
	Result []string `json:"result,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	req := FizzBuzzRequest{}
	var err error

	if req.Int1, err = strconv.Atoi(q.Get("int1")); err != nil {
		req.Int1 = 0
	}
	if req.Int2, err = strconv.Atoi(q.Get("int2")); err != nil {
		req.Int2 = 0
	}
	if req.Limit, err = strconv.Atoi(q.Get("limit")); err != nil {
		req.Limit = 0
	}

	req.Str1 = q.Get("str1")
	req.Str2 = q.Get("str2")

	recordStats(req)

	w.Header().Set("Content-Type", "application/json")

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FizzBuzzResponse{Errors: validationErrors})
		return
	}

	result := req.Compute()
	json.NewEncoder(w).Encode(FizzBuzzResponse{Result: result})
}
