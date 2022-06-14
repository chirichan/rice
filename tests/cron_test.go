package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/woxingliu/rice"
)

// go test -v -timeout 30s -run ^TestCronRun$ github.com/woxingliu/rice
func TestCronRun(t *testing.T) {

	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// 0 8 * * *
	// 0 8 * * 1
	rice.CronRun(withTimeout, "*/1 * * * * *", func() {
		log.Println(time.Now())
	})
}

func TestTickerRun(t *testing.T) {

	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()

	err := rice.TickerRun(withTimeout, 1*time.Second, func() error {
		log.Println(time.Now())
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTimerRun(t *testing.T) {

	err := rice.TimerRun(context.Background(), time.Date(2022, 3, 20, 19, 21, 0, 0, time.Local),
		func() error {
			log.Println(time.Now())
			return nil
		})
	if err != nil {
		t.Error(err)
	}

}

func TestTickerRunWithStartTimeContext(t *testing.T) {

	var ctx = context.Background()
	//var t1 int64 = 1651916258
	var t2 = time.Date(2022, 5, 7, 17, 52, 20, 0, time.Local)
	//sep := time.Duration(3600000000000)

	err := rice.TickerRunWithStartTimeContext(ctx, nil, t2, 2*time.Second, func(ctx context.Context) error {

		log.Println(time.Now(), "A")
		return nil
	})

	if err != nil {
		t.Error(err)
	}

}
