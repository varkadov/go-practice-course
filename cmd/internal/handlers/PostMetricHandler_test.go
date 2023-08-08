package handlers

import (
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

type want struct {
	statusCode int
}

func TestPostMetricHandler(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		method string
		want   want
	}{
		{
			name:   "Invalid metric type",
			url:    "/update/some-invalid-metric/1/1",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Invalid gauge metric",
			url:    "/update/gauge/x/x",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Invalid counter metric",
			url:    "/update/counter/x/x",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Valid gauge metric",
			url:    "/update/gauge/1/1",
			method: http.MethodPost,
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name:   "Valid counter metric",
			url:    "/update/counter/1/1",
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
			s := storage.NewMemStorage()

			PostMetricHandler(s)(w, r)

			res := w.Result()
			_ = res.Body.Close()

			if res.StatusCode != tt.want.statusCode {
				t.Errorf("Status code expected: %d, actual: %d", tt.want.statusCode, res.StatusCode)
			}
		})
	}
}
