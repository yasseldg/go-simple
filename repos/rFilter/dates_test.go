package rFilter

import (
	"testing"
	"time"
)

func TestCreatedAt_gt(t *testing.T) {
	f := &Filters{}
	tm := time.Now()
	f.CreatedAt_gt(tm)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestCreatedAt_lt(t *testing.T) {
	f := &Filters{}
	tm := time.Now()
	f.CreatedAt_lt(tm)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestUpdatedAt_gt(t *testing.T) {
	f := &Filters{}
	tm := time.Now()
	f.UpdatedAt_gt(tm)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestUpdatedAt_lt(t *testing.T) {
	f := &Filters{}
	tm := time.Now()
	f.UpdatedAt_lt(tm)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}
