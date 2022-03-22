package rice

import (
	"context"
	"log"
	"testing"
	"time"
)

// go test -v -timeout 30s -run ^TestCronRun$ github.com/woxingliu/rice
func TestCronRun(t *testing.T) {

	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	CronRun(withTimeout, "*/1 * * * * *", func() {
		log.Println(time.Now())
	})
}

func TestTickerRun(t *testing.T) {

	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()

	err := TickerRun(withTimeout, 1*time.Second, func() error {
		log.Println(time.Now())
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTimerRun(t *testing.T) {

	err := TimerRun(context.Background(), time.Date(2022, 3, 20, 19, 21, 0, 0, time.Local),
		func() error {
			log.Println(time.Now())
			return nil
		})
	if err != nil {
		t.Error(err)
	}

}
