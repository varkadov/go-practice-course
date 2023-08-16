package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandler_PostMetricHandler(t *testing.T) {
	type want struct {
		status int
		body   string
	}

	tests := []struct {
		name    string
		url     string
		storage *storage
		want    want
	}{
		{
			name:    "Should return 400 status",
			url:     "/?metricType=gauge&metricName=alloc&metricValue=1",
			storage: newStorage("", errors.New("400")),
			want: want{
				status: http.StatusBadRequest,
				body:   "400\n",
			},
		},
		{
			name:    "Should return 200 status",
			url:     "/?metricType=gauge&metricName=alloc&metricValue=1",
			storage: newStorage("", nil),
			want: want{
				status: http.StatusOK,
				body:   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				storage: tt.storage,
			}

			router := chi.NewRouter()
			router.Post("/", h.PostMetricHandler)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, tt.url, nil)

			handler := http.HandlerFunc(router.ServeHTTP)
			handler.ServeHTTP(w, r)

			res := w.Result()
			body, err := io.ReadAll(res.Body)
			defer res.Body.Close()

			assert.NoError(t, err)
			assert.Equal(t, tt.want.status, res.StatusCode)
			assert.Equal(t, tt.want.body, string(body))
		})
	}
}
