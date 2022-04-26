package rice

import (
	"reflect"
	"testing"
)

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

func TestSliceSet(t *testing.T) {
	type args struct {
		s1 []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{name: "1", args: args{[]int64{2, 3, 5, 4, 3, 0, 0, 0, 2, 1}}, want: []int64{2, 3, 5, 4, 0, 1}},
		{name: "2", args: args{[]int64{0, 0, 0, 0, 0, 0}}, want: []int64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSet(tt.args.s1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
