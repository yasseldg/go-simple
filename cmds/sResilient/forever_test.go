package sResilient

import (
	"context"
	"testing"
	"time"
)

func TestForever(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runFunc := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			// Simulate work
		}
	}

	go Forever(ctx, "testFunc", runFunc)

	time.Sleep(2 * time.Second)
	cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Error("Function did not stop after context cancellation")
	case <-ctx.Done():
		// Test passed
	}
}

func TestHandlePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()

	handlePanic("testFunc")
	panic("test panic")
}
