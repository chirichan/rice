package rice

import (
	"fmt"
	"testing"
)

func TestWirteFile(t *testing.T) {
	err := WriteFileIfNotExists("abc/a.txt", []byte("ddd"))
	fmt.Printf("err: %v\n", err)
}

func TestWriteFileWhatever(t *testing.T) {
	type args struct {
		name string
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				name: "a/bc/d.txt",
				data: []byte("aaa"),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				name: "a/bc/d.txt",
				data: []byte("bbb"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFileWhatever(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFileWhatever() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppendFileIfExists(t *testing.T) {
	type args struct {
		name string
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				name: "a/bc/d.txt",
				data: []byte("aaa"),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				name: "a/bc/d.txt",
				data: []byte("bbb"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AppendFileIfExists(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AppendFileIfExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
