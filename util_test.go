package rice

import (
	"testing"
)

func TestRemoveInvisibleChars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//lint:ignore ST1018 ignore
		{"", args{"bushi​"}, "bushi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInvisibleChars(tt.args.s); got != tt.want {
				t.Errorf("RemoveInvisibleChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitNString(t *testing.T) {
	type args struct {
		s     string
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:     "",
				index: 10,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitNString(tt.args.s, tt.args.index); got != tt.want {
				t.Errorf("SplitNString() = %v, want %v", got, tt.want)
			}
		})
	}
}
