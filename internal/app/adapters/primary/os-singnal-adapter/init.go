package os_singnal_adapter

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type OsSignalAdapter struct{}

func New() *OsSignalAdapter {
	return &OsSignalAdapter{}
}

func (a *OsSignalAdapter) Start(ctx context.Context) error {
	osSignCh := make(chan os.Signal, 1)

	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-ctx.Done():
		err := ctx.Err()

		return err
	case sig := <-osSignCh:
		err := fmt.Errorf("получен сигнал [%s]", sig.String())

		return err
	}
}
