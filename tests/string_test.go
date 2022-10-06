package tests

import (
	"sync"
	"testing"

	"github.com/chirichan/rice"
)

func TestNextId(t *testing.T) {

	var wg sync.WaitGroup
	for i := 0; i < 600000; i++ {
		wg.Add(1)
		go func() {
			t.Error(rice.NextStringId())
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestGetHostname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"a", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rice.GetHostname(); got != tt.want {
				t.Errorf("GetHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}
