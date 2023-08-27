package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/varkadov/go-practice-course/internal/models"
)

// TODO Move this mock into the common place
type storage struct {
	value float64
	err   error
}

func (s *storage) Get(metricType, metricName string) (*models.Metrics, error) {
	return &models.Metrics{
		ID:    metricName,
		MType: metricType,
		Value: &s.value,
	}, s.err
}

func (s *storage) Set(metricType, metricName, metricValue string) (*models.Metrics, error) {
	return &models.Metrics{
		ID:    metricName,
		MType: metricType,
		Value: &s.value,
	}, s.err
}

func (s *storage) GetAll() []string {
	return make([]string, 0)
}

func newStorage(value string, err error) *storage {
	return &storage{value: value, err: err}
}

func TestHandler_GetMetricHandler(t *testing.T) {
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
			name:    "Should return 404 status if metric doesn't exist",
			url:     "/?metricType=counter&metricName=Alloc",
			storage: newStorage("", errors.New("404")),
			want: want{
				status: http.StatusNotFound,
				body:   "404\n",
			},
		},
		{
			name:    "Should return 200 status if metric exists",
			url:     "/?metricType=counter&metricName=Alloc",
			storage: newStorage("value", nil),
			want: want{
				status: http.StatusOK,
				body:   "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{storage: tt.storage}

			router := chi.NewRouter()
			router.Get("/", h.MetricsValueHandler)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

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
