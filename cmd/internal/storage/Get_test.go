package storage

import (
	"testing"
)

func TestMemStorage_Get(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		t string
		n string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return gauge metric value",
			fields: fields{
				gauge:   map[string]float64{"metricName": 4.2},
				counter: nil,
			},
			args: args{
				t: "gauge",
				n: "metricName",
			},
			want:    "4.2",
			wantErr: false,
		},
		{
			name: "should return counter metric value",
			fields: fields{
				gauge:   nil,
				counter: map[string]int64{"metricName": 4},
			},
			args: args{
				t: "counter",
				n: "metricName",
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: "should not return non-existent metric",
			fields: fields{
				gauge:   nil,
				counter: map[string]int64{"metricName": 4},
			},
			args: args{
				t: "nonExistentMetric",
				n: "metricName",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should not return non-existent gauge metric",
			fields: fields{
				gauge:   map[string]float64{"metricName": 4.2},
				counter: nil,
			},
			args: args{
				t: "gauge",
				n: "notExistentMetricName",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should not return non-existent counter metric",
			fields: fields{
				gauge:   nil,
				counter: map[string]int64{"metricName": 4},
			},
			args: args{
				t: "counter",
				n: "notExistentMetricName",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			got, err := s.Get(tt.args.t, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
