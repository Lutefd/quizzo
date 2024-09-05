package commons

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		msg      string
		expected string
	}{
		{"Client Error", 400, "Bad Request", `{"error":"Bad Request"}`},
		{"Server Error", 500, "Internal Server Error", `{"error":"Internal Server Error"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			RespondWithError(w, tt.code, tt.msg)

			if w.Code != tt.code {
				t.Errorf("expected status code %d, got %d", tt.code, w.Code)
			}

			if w.Body.String() != tt.expected {
				t.Errorf("expected body %q, got %q", tt.expected, w.Body.String())
			}

			if w.Header().Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type to be application/json")
			}
		})
	}
}

func TestRespondWithJSON(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		payload  interface{}
		expected string
	}{
		{"Simple Object", 200, map[string]string{"message": "success"}, `{"message":"success"}`},
		{"Array", 200, []int{1, 2, 3}, `[1,2,3]`},
		{"Nested Object", 201, map[string]interface{}{"user": map[string]string{"name": "John"}}, `{"user":{"name":"John"}}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			RespondWithJSON(w, tt.code, tt.payload)

			if w.Code != tt.code {
				t.Errorf("expected status code %d, got %d", tt.code, w.Code)
			}

			if w.Body.String() != tt.expected {
				t.Errorf("expected body %q, got %q", tt.expected, w.Body.String())
			}

			if w.Header().Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type to be application/json")
			}

			var response interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("failed to unmarshal response: %v", err)
			}
		})
	}
}

func TestRespondWithJSONInvalidPayload(t *testing.T) {
	w := httptest.NewRecorder()
	invalidPayload := make(chan int)
	RespondWithJSON(w, 200, invalidPayload)

	if w.Code != 500 {
		t.Errorf("expected status code 500, got %d", w.Code)
	}

	if w.Body.String() != "" {
		t.Errorf("expected empty body, got %q", w.Body.String())
	}
}
