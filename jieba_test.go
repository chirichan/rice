package rice

import (
	"testing"
)

func TestJieba(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"", args{s: ""}, []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Jieba(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jieba() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
