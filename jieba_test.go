package rice

import (
	"reflect"
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
		{"", args{s: "我的生活"}, []string{"我的", "我的生活", "的生活", "生活"}, false},
		{"", args{s: "my life"}, []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Jieba(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jieba() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Jieba() got = %v", got)
		})
	}
}

func TestPinyin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "我的"}, "wode"},
		{"", args{s: ""}, ""},
		{"", args{s: "！@#￥%……&*（）——+=-😀"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pinyin(tt.args.s); got != tt.want {
				t.Errorf("Pinyin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJiebaAndPinyin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"", args{s: "我的"}, []string{"我的", "wode"}, false},
		{"", args{s: ""}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JiebaAndPinyin(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("JiebaAndPinyin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JiebaAndPinyin() = %v, want %v", got, tt.want)
			}
		})
	}
}
