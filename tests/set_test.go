package tests

import (
	"reflect"
	"testing"

	"github.com/latext/rice"
)

func TestSliceDifferenceBoth(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"1"}, []string{"2"}}, []string{"1", "2"}},
		{"2", args{[]string{}, []string{"2"}}, []string{"2"}},
		{"3", args{[]string{"1"}, []string{}}, []string{"1"}},
		{"4", args{[]string{"1", "1", "2", "3"}, []string{"0", "2", "3", "4"}}, []string{"1", "1", "0", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rice.SliceDifferenceBoth(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceDifferenceBoth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceDifference(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"1"}, []string{"2"}}, []string{"1"}},
		{"2", args{[]string{}, []string{"2"}}, nil},
		{"3", args{[]string{"1"}, []string{}}, []string{"1"}},
		{"4", args{[]string{"1", "1", "2", "3"}, []string{"0", "2", "3", "4"}}, []string{"1", "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rice.SliceDifference(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := rice.SliceIn(tt.args.e, tt.args.s); got != tt.want {
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
			if got := rice.RemoveDuplicates(tt.args.s1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
