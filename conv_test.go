package rice

import (
	"reflect"
	"testing"
)

func TestStrconvParseInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "test1", args: args{"1645270804"}, want: 1645270804},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := StrconvParseInt(tt.args.s); got != tt.want {
				t.Errorf("StrconvParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrconvFormatInt(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{1645270804}, want: "1645270804"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrconvFormatInt(tt.args.i); got != tt.want {
				t.Errorf("StrconvFormatInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrconvParseFloat(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test1", args: args{85}, want: 0.85},
		{name: "test2", args: args{0}, want: 0.00},
		{name: "test3", args: args{1}, want: 0.01},
		{name: "test4", args: args{100}, want: 1.00},
		{name: "test5", args: args{80}, want: 0.80},
		{name: "test6", args: args{20}, want: 0.20},
		{name: "test7", args: args{99}, want: 0.99},
		{name: "test8", args: args{999999999}, want: 9999999.99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := StrconvParseFloat(tt.args.i); got != tt.want {
				t.Errorf("StrconvParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringByte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"hel, lo"}, want: []byte("hel, lo")},
		{name: "test1", args: args{" hel, lo"}, want: []byte(" hel, lo")},
		{name: "test1", args: args{"hel, lo "}, want: []byte("hel, lo ")},
		{name: "test1", args: args{"hel ;lk😄lo"}, want: []byte("hel ;lk😄lo")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringByte(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceRemoveIndexUnOrder(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{s: []int{2, 4, 8, 9}, i: 2}, want: []int{2, 4, 9}},
		{name: "test2", args: args{s: []int{9, 99, 14, 345345, 453, 90, 239, 29832}, i: 3}, want: []int{9, 99, 14, 453, 90, 239, 29832}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceRemoveIndexUnOrder(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceRemoveIndexUnOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
