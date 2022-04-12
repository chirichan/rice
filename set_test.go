package rice

import "testing"

func TestSliceIn(t *testing.T) {
	type args struct {
		e int
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{e: 0, s: []int{0}}, want: true},
		{name: "2", args: args{e: 0, s: []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIn(tt.args.e, tt.args.s); got != tt.want {
				t.Errorf("SliceIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
