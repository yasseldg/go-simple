package sResilient

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNotifier(t *testing.T) {
	handlerCalled := false
	handler := func(sig os.Signal) {
		handlerCalled = true
	}

	notifier := Notifier(handler, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go notifier.Listen(ctx)

	// Simulate sending a signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	signalChan <- syscall.SIGINT

	time.Sleep(1 * time.Second)

	if !handlerCalled {
		t.Errorf("Expected handler to be called, but it was not")
	}

	cancel()
	select {
	case <-notifier.ShutdownChan():
		// Test passed
	case <-time.After(1 * time.Second):
		t.Error("Expected notifier to shut down, but it did not")
	}
}

func TestDefaultHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got %v", r)
		}
	}()

	defaultHandler(syscall.SIGINT)
}
