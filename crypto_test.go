package rice

import (
	"fmt"
	"testing"
	"time"
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

func TestLocation(t *testing.T) {

	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Error(err)
	}
	t.Errorf("l: %v\n", l)

	time13 := time.Date(2022, 3, 11, 18, 46, 0, 0, time.Local)
	fmt.Printf("time13: %v\n", time13)
	fmt.Printf("time13.Unix(): %v\n", time13.Unix())

	time14 := time.Date(2022, 3, 11, 18, 46, 0, 0, time.UTC)
	fmt.Printf("time14: %v\n", time14)
	fmt.Printf("time14.Unix(): %v\n", time14.Unix())

	fmt.Printf("time.Now(): %v\n", time.Now())

	fmt.Printf("time.Now().UTC(): %v\n", time.Now().UTC())

	fmt.Printf("time.Now().UTC().Local(): %v\n", time.Now().UTC().Local())

	fmt.Printf("time.Now().Local(): %v\n", time.Now().Local())

}
