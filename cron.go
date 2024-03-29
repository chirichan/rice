package rice

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type Task func() error
type TaskContext func(ctx context.Context) error
type contextKey string

func TimerRun(ctx context.Context, tm time.Time, task ...Task) error {
	timer := time.NewTimer(time.Until(tm))
	defer timer.Stop()

	var err error
	select {
	case <-timer.C:
		for _, tk := range task {
			if err = tk(); err != nil {
				err = errors.Join(err)
			}
		}
		return err
	case <-ctx.Done():
		return err
	}
}

func TickerRun(ctx context.Context, d time.Duration, task ...Task) error {
	for _, tk := range task {
		err := tk()
		if err != nil {
			return err
		}
	}

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

// TickerRunWithStartTimeContext 到达 tm 时间之后，开始以 tikcer 的方式执行 task
func TickerRunWithStartTimeContext(ctx context.Context, wg *sync.WaitGroup, tm time.Time, d time.Duration, task ...TaskContext) error {
	now := time.Now()
	for ; now.After(tm); tm = tm.Add(d) {
		ctx = context.WithValue(ctx, contextKey("tm"), tm)
		for _, tk := range task {
			err := tk(ctx)
			if err != nil {
				return fmt.Errorf("%v have an error - %w", tm, err)
			}
		}
	}

	if wg != nil {
		wg.Done()
	}

	timer := time.NewTimer(time.Until(tm))
	defer timer.Stop()

	select {
	case <-timer.C:
		ctx := context.WithValue(ctx, contextKey("tm"), time.Now())
		err := TickerRunContext(ctx, d, task...)
		return err
	case <-ctx.Done():
		return nil
	}
}

// TickerRunContext 立即开始以 ticker 的方式执行 task
func TickerRunContext(ctx context.Context, d time.Duration, task ...TaskContext) error {
	for _, tk := range task {
		err := tk(ctx)
		if err != nil {
			return err
		}
	}
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx := context.WithValue(ctx, contextKey("tm"), time.Now())
			for _, tk := range task {
				err := tk(ctx)
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
	<-ctx.Done()
	return nil
}

func CronRunM(ctx context.Context, corn string, task ...func()) error {
	c := cron.New()
	defer c.Stop()
	for _, tk := range task {
		_, err := c.AddFunc(corn, tk)
		if err != nil {
			return err
		}
	}
	c.Start()
	<-ctx.Done()
	return nil
}
