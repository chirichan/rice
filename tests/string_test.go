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
