package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
		body    string
		storage *storageMock
		want    want
	}{
		{
			name:    "Should return 400 status",
			url:     "/",
			body:    `{"id": "alloc", "type": "gauge", "value": 1}`,
			storage: newStorage(1, errors.New("400")),
			want: want{
				status: http.StatusBadRequest,
				body:   "400\n",
			},
		},
		{
			name:    "Should return 200 status",
			url:     "/",
			body:    `{"id": "alloc", "type": "gauge", "value": 1}`,
			storage: newStorage(1, nil),
			want: want{
				status: http.StatusOK,
				body:   `{"id": "alloc", "type": "gauge", "value": 1}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				storage: tt.storage,
			}

			router := chi.NewRouter()
			router.Post("/", h.MetricsUpdateHandler)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))

			handler := http.HandlerFunc(router.ServeHTTP)
			handler.ServeHTTP(w, r)

			res := w.Result()
			body, err := io.ReadAll(res.Body)
			defer res.Body.Close()

			assert.NoError(t, err)
			assert.Equal(t, tt.want.status, res.StatusCode)
			assert.JSONEq(t, tt.want.body, string(body))
		})
	}
}
