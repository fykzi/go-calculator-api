package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fykzi/go-calculator-api/internal/application"
)

func sanitizeJSON(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name             string
		method           string
		body             interface{}
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Valid request",
			method:           http.MethodPost,
			body:             map[string]string{"expression": "2 + 2 * 2"},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"result":"6"}`,
		},
		{
			name:             "Wrong Method",
			method:           http.MethodGet,
			body:             nil,
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: `{"error":"Invalid request method"}`,
		},
		{
			name:             "Invalid Body",
			method:           http.MethodPost,
			body:             "invalid body",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Bad request"}`,
		},
		{
			name:             "Invalid Expression",
			method:           http.MethodPost,
			body:             map[string]string{"expression": "2 + 2 * 2)"},
			expectedStatus:   http.StatusUnprocessableEntity,
			expectedResponse: `{"error":"Expression is not valid"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestBody []byte
			if tt.body != nil {
				var err error
				requestBody, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatal(err)
				}
			}

			reqPath := "/api/v1/calculate"

			req := httptest.NewRequest(tt.method, reqPath, bytes.NewBuffer(requestBody))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(application.CalcHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Got wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedResponse != "" {
				if sanitizedBody := sanitizeJSON(rr.Body.String()); sanitizedBody != sanitizeJSON(tt.expectedResponse) {
					t.Errorf("Got unexpected response body: got '%v' want '%v'", rr.Body.String(), tt.expectedResponse)
				}
			}
		})
	}
}

