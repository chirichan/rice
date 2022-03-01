package rice

import (
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	type args struct {
		key []byte
		s   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESEncrypt(tt.args.key, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AESEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESDecrypt(t *testing.T) {
	type args struct {
		key []byte
		s   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESDecrypt(tt.args.key, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AESDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESNewGCMEncrypt(t *testing.T) {

	keystring := "972ec8dd995743d981417981ac2f30db"

	s, err := AESNewGCMEncrypt(keystring, "hello i am neko")
	if err != nil {
		t.Error(err)
	}
	t.Error(s)
}

func TestAESNewGCMDecrypt(t *testing.T) {

	keystring := "972ec8dd995743d981417981ac2f30db"
	nonceString := "e8ce8ffdbb0a710ad3999ba2"
	ciphertext := "f116733f86881a30d8a84be3c67e07192e93f121d8d1c9326456a8bb3843b1"

	s, err := AESNewGCMDecrypt(keystring, nonceString, ciphertext)
	if err != nil {
		t.Error(err)
	}

	t.Error(s)
}
