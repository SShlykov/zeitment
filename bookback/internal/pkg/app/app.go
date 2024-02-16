package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	ctx      context.Context
	closeCtx func()
}

func NewApp(configPath string) (*App, error) {
	ctx, closeCtx := context.WithCancel(context.Background())
	_ = configPath
	app := &App{
		ctx:      ctx,
		closeCtx: closeCtx,
	}

	return app, nil
}

func (app *App) Run() error {
	go handleSignals(app.closeCtx)

	if err := startServer(app.ctx); err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context) error {
	laddr, err := net.ResolveTCPAddr("tcp", ":7709")
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}

	defer func(l *net.TCPListener) {
		err = l.Close()
	}(l)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err = l.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}

			_, err = l.Accept()
			if err != nil && !os.IsTimeout(err) {
				return err
			}
		}
	}
}

func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cancel()
			fmt.Println("Interrupted")
			return
		}
	}
}
