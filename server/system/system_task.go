package system

import (
	"context"
	"log/slog"
	"time"
)

func StartTask(ctx context.Context, task func(ctx context.Context) error, interval time.Duration, name string) {
	slog.Info("Starting task", "name", name, "interval", interval)

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("Stopping task", "name", name)
			return
		case <-t.C:
			go func() {
				slog.Info("Running task", "name", name)
				err := task(ctx)
				if err != nil {
					slog.Error("Error in task", "name", name, "error", err)
				}
			}()
		}
	}
}
