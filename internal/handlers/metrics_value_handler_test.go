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
	"github.com/varkadov/go-practice-course/internal/models"
)

// TODO Move this mock into the common place
type storageMock struct {
	value float64
	err   error
}

func (s *storageMock) Get(metricType, metricName string) (*models.Metrics, error) {
	return &models.Metrics{
		ID:    metricName,
		MType: metricType,
		Value: &s.value,
	}, s.err
}

func (s *storageMock) Set(metricType, metricName, _ string) (*models.Metrics, error) {
	return &models.Metrics{
		ID:    metricName,
		MType: metricType,
		Value: &s.value,
	}, s.err
}

func (s *storageMock) GetAll() []string {
	return make([]string, 0)
}

func newStorage(value float64, err error) *storageMock {
	return &storageMock{value: value, err: err}
}

func TestHandler_GetMetricHandler(t *testing.T) {
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
			name:    "Should return 404 status if metric doesn't exist",
			url:     "/",
			body:    `{"id": "Alloc", "type": "counter"}`,
			storage: newStorage(1, errors.New("404")),
			want: want{
				status: http.StatusNotFound,
				body:   "404\n",
			},
		},
		{
			name:    "Should return 200 status if metric exists",
			url:     "/",
			body:    `{"id": "Alloc", "type": "counter"}`,
			storage: newStorage(1, nil),
			want: want{
				status: http.StatusOK,
				body:   `{"id": "Alloc", "type": "counter", "value": 1}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{storage: tt.storage}

			router := chi.NewRouter()
			router.Post("/", h.MetricsValueHandler)

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
