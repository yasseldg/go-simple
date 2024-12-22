package rFilter

import (
	"testing"
)

func TestInt(t *testing.T) {
	f := &Filters{}
	f.Int("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_in(t *testing.T) {
	f := &Filters{}
	f.Int_in("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_nin(t *testing.T) {
	f := &Filters{}
	f.Int_nin("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gt(t *testing.T) {
	f := &Filters{}
	f.Int_gt("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gte(t *testing.T) {
	f := &Filters{}
	f.Int_gte("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_lt(t *testing.T) {
	f := &Filters{}
	f.Int_lt("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_lte(t *testing.T) {
	f := &Filters{}
	f.Int_lte("field", 123)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gt_lt(t *testing.T) {
	f := &Filters{}
	f.Int_gt_lt("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gt_lte(t *testing.T) {
	f := &Filters{}
	f.Int_gt_lte("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gte_lt(t *testing.T) {
	f := &Filters{}
	f.Int_gte_lt("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestInt_gte_lte(t *testing.T) {
	f := &Filters{}
	f.Int_gte_lte("field", 123, 456)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}
