package rice

import (
	"testing"
)

func TestLocalHostname(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalHostname()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalHostname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("LocalHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalAddr(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalAddr()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("LocalAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveInvisibleChars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
