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

func TestPunycodeEncode(t *testing.T) {
	type args struct {
		chineseDomain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				chineseDomain: "中国",
			},
			want:    "XN--FIQS8S",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PunycodeEncode(tt.args.chineseDomain)
			if (err != nil) != tt.wantErr {
				t.Errorf("PunycodeEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PunycodeEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
