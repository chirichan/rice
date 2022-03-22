package rice

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
)

type Task func() error

func TimerRun(ctx context.Context, tm time.Time, task ...Task) error {

	timer := time.NewTimer(tm.Sub(time.Now()))
	defer timer.Stop()

	select {
	case <-timer.C:
		for _, tk := range task {
			err := tk()
			if err != nil {
				return err
			}
		}
		return nil
	case <-ctx.Done():
		return nil
	}

}

func TickerRun(ctx context.Context, d time.Duration, task ...Task) error {

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, tk := range task {
				err := tk()
				if err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func CronRun(ctx context.Context, corn string, task ...func()) error {

	c := cron.New(cron.WithSeconds())
	defer c.Stop()

	for _, tk := range task {
		_, err := c.AddFunc(corn, tk)
		if err != nil {
			return err
		}
	}

	c.Start()

	select {
	case <-ctx.Done():
		return nil
	}
}
