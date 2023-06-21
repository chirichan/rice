package rice

import "testing"

func TestCountWeek(t *testing.T) {
	type args struct {
		TimeFormat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{TimeFormat: "2023-06-21 20:34:00"}, 4},
		{"", args{TimeFormat: "2023-06-30 20:34:00"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWeek(tt.args.TimeFormat); got != tt.want {
				t.Errorf("CountWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
