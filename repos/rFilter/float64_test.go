package rFilter

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	f := &Filters{}
	f.Float64("field", 1.23)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_in(t *testing.T) {
	f := &Filters{}
	f.Float64_in("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_nin(t *testing.T) {
	f := &Filters{}
	f.Float64_nin("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gt(t *testing.T) {
	f := &Filters{}
	f.Float64_gt("field", 1.23)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gte(t *testing.T) {
	f := &Filters{}
	f.Float64_gte("field", 1.23)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_lt(t *testing.T) {
	f := &Filters{}
	f.Float64_lt("field", 1.23)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_lte(t *testing.T) {
	f := &Filters{}
	f.Float64_lte("field", 1.23)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gt_lt(t *testing.T) {
	f := &Filters{}
	f.Float64_gt_lt("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gt_lte(t *testing.T) {
	f := &Filters{}
	f.Float64_gt_lte("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gte_lt(t *testing.T) {
	f := &Filters{}
	f.Float64_gte_lt("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestFloat64_gte_lte(t *testing.T) {
	f := &Filters{}
	f.Float64_gte_lte("field", 1.23, 4.56)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}
