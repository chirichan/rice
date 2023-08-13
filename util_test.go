package rice

import (
	"fmt"
	"os"
	"path/filepath"
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

func TestIsDir(t *testing.T) {
	base := ".vscode"

	// f, _ := os.Open(".vscode")
	// finfo, _ := f.Readdir(-1)

	// for _, v := range finfo {
	// 	fmt.Printf("v.Name(): %v\n", v.Name())
	// }

	de, err := os.ReadDir(base)
	_ = err
	for _, v := range de {

		fmt.Printf("v.Name(): %v\n", filepath.Join(base, v.Name()))

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
