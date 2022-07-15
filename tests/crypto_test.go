package tests

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/woxingliu/rice"
)

func TestBCryptGenerateFromPassword(t *testing.T) {

	pwd := "a😀"
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			s, err := rice.BCryptGenerateFromPassword(pwd)

			if err != nil {
				t.Error(err)
			}

			t.Log(s, rice.BCryptCompareHashAndPassword(pwd, s))

		}()

	}

	wg.Wait()

}

// go test -run ^TestBCryptGenerateFromPassword$ github.com/woxingliu/rice/tests -v -count=1

func TestCheckPassword(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", args{"123456abc123456abc123456abc"}, false},
		{"2", args{"abcdef123456ADLDFIO&*^^%)"}, false},
		{"3", args{"aeafhrtyrrhjghdYUOOJJ^)&%$#$#$&*())ha😀🆒😀🆒😀🆒😀🆒wr#$#$"}, false},
		{"4", args{"%%$"}, false},
		{"5", args{"GGfggsd"}, false},
		{"6", args{"ERRETt4645"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rice.CheckPassword(tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test -timeout 30s -run ^TestCheckPassword$ github.com/woxingliu/rice/tests -v -count=1

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

func TestCTREncrypt(t *testing.T) {

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	byteUnsafe := rice.StringByteUnsafe("u.4%,:oV0xyg8poZAO$?d6KYU0=S?/1,")
	t.Log(len(byteUnsafe))

	encodeToString := hex.EncodeToString(iv)

	t.Log(encodeToString)

	key := encodeToString
	plaintext := "l"

	a := rice.NewCTRCrypt()

	s, err := a.Encrypt(key, plaintext)
	if err != nil {
		t.Error(err)
	}

	t.Log(s)

	s2, _ := a.Decrypt(key, s)

	t.Log(s2)

	generateFromPassword, _ := rice.BCryptGenerateFromPassword(encodeToString)

	t.Log(generateFromPassword)
}

func TestFullPassword(t *testing.T) {

	for i := 0; i < 200; i++ {
		s, err := rice.FullPassword(4, 32)
		if err != nil {
			t.Error(err)
		}

		if s == "" {
			t.Log("null")
		} else {
			t.Log(s)
		}

	}
}
