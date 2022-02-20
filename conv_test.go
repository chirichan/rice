package ha

import (
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
			if got := StrconvParseInt(tt.args.s); got != tt.want {
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
