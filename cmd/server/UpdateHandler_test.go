package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type want struct {
	statusCode int
}

func TestUpdateHandler(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		method string
		want   want
	}{
		{
			name:   "GET method",
			url:    "/",
			method: http.MethodGet,
			want: want{
				statusCode: http.StatusNotFound,
			},
		},
		{
			name:   "Not found",
			url:    "/some-invalid-url",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusNotFound,
			},
		},
		{
			name:   "Invalid metric type",
			url:    "/some-invalid-metric/1/1",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Invalid gauge metric",
			url:    "/gauge/x/x",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Invalid counter metric",
			url:    "/counter/x/x",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Valid gauge metric",
			url:    "/gauge/1/1",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name:   "Valid counter metric",
			url:    "/counter/1/1",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.url, nil)
			s := NewMemStorage()

			UpdateHandler(s)(w, r)

			res := w.Result()

			if res.StatusCode != tt.want.statusCode {
				t.Errorf("Status code expected: %d, actual: %d", tt.want.statusCode, res.StatusCode)
			}
		})
	}
}
