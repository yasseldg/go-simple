package sResilient

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yasseldg/go-simple/logs/sLog"
)

var (
	defaultSignals = []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGABRT,
		// os.Kill is a signal that cannot be captured.
	}
)

// SignalHandler defines the handler function type for signals.
type SignalHandler func(os.Signal)

// SignalNotifier configures listening for system signals.
type SignalNotifier struct {
	signals  []os.Signal
	handler  SignalHandler
	stopChan chan struct{} // Added a channel to signal shutdown
}

// Notifier creates a new instance of SignalNotifier.
// It allows customizing the signals to listen for and defines an optional handler.
func Notifier(handler SignalHandler, signals ...os.Signal) *SignalNotifier {
	if len(signals) == 0 {
		signals = defaultSignals
	}

	if handler == nil {
		handler = defaultHandler
	}

	return &SignalNotifier{
		signals:  signals,
		handler:  handler,
		stopChan: make(chan struct{}), // Initialize the stop channel
	}
}

// Listen starts listening for the configured signals.
// It stops when the context is canceled.
func (sn *SignalNotifier) Listen(ctx context.Context) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, sn.signals...)

	go func() {
		defer signal.Stop(s)
		sLog.Info("Listening for signals...")

		for {
			select {
			case <-ctx.Done():
				sLog.Debug("Signal listener stopped due to ctx.Done()")
				return

			case sig := <-s:
				if sn.handler != nil {
					sn.handler(sig)
				}

				close(sn.stopChan) // Signal the shutdown

				sLog.Debug("Signal listener stopped due to signal: %v", sig)
				return
			}
		}
	}()
}

// ShutdownChan returns the channel that will be closed when it's time to shut down.
func (sn *SignalNotifier) ShutdownChan() <-chan struct{} {
	return sn.stopChan
}

func defaultHandler(sig os.Signal) {
	sLog.Error("Default handler, signal received: %v", sig)
	// Implement graceful shutdown logic here instead of os.Exit(0)
}
