package rice

import (
	"sync"
	"testing"
)

func TestNextId(t *testing.T) {

	var wg sync.WaitGroup
	for i := 0; i < 600000; i++ {
		wg.Add(1)
		go func() {
			t.Error(NextStringId())
			wg.Done()
		}()
	}
	wg.Wait()
}
