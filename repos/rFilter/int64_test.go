package rFilter

import (
	"testing"
)

func TestInt64(t *testing.T) {
	f := &Filters{}
	f.Int64("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_in(t *testing.T) {
	f := &Filters{}
	f.Int64_in("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_nin(t *testing.T) {
	f := &Filters{}
	f.Int64_nin("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gt(t *testing.T) {
	f := &Filters{}
	f.Int64_gt("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gte(t *testing.T) {
	f := &Filters{}
	f.Int64_gte("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_lt(t *testing.T) {
	f := &Filters{}
	f.Int64_lt("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_lte(t *testing.T) {
	f := &Filters{}
	f.Int64_lte("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gt_lt(t *testing.T) {
	f := &Filters{}
	f.Int64_gt_lt("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gt_lte(t *testing.T) {
	f := &Filters{}
	f.Int64_gt_lte("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gte_lt(t *testing.T) {
	f := &Filters{}
	f.Int64_gte_lt("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt64_gte_lte(t *testing.T) {
	f := &Filters{}
	f.Int64_gte_lte("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}
