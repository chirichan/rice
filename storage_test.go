package rice

import (
	"testing"
)

func TestFileStorage_Save(t *testing.T) {
	type args struct {
		name string
		data []byte
	}
	tests := []struct {
		name    string
		fs      *FileStorage
		args    args
		wantErr bool
	}{
		{"1", NewFileStorage("", "20220519"), args{"test.txt", []byte("123456789")}, false},
		{"2", NewFileStorage("", "20220620"), args{"test1.txt", []byte("89474984")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.Save(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FileStorage.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileStorage_Append(t *testing.T) {
	type args struct {
		name string
		data any
	}
	tests := []struct {
		name    string
		fs      *FileStorage
		args    args
		wantErr bool
	}{
		{"1", NewFileStorage("", "20220519"), args{"test.json", []string{"2", "3"}}, false},
		{"2", NewFileStorage("", "20220619"), args{"test.json", []string{"1", "2"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.Append("", tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FileStorage.Append() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
