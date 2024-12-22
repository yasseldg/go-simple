package rFilter

import (
	"testing"
	"time"

	"github.com/yasseldg/go-simple/trading/tSide"
)

func TestClone(t *testing.T) {
	f := &Filters{}
	cloned := f.Clone()
	if cloned == nil {
		t.Errorf("Expected cloned filter, got nil")
	}
}

func TestOper(t *testing.T) {
	f := &Filters{}
	if f.Oper() == nil {
		t.Errorf("Expected InterOper instance, got nil")
	}
}

func TestTsField(t *testing.T) {
	f := &Filters{}
	f.TsField(1625097600, 1627689600, "timestamp")
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestTs(t *testing.T) {
	f := &Filters{}
	f.Ts(1625097600, 1627689600)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestTsIn(t *testing.T) {
	f := &Filters{}
	f.TsIn(1625097600, 1627689600)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestStates(t *testing.T) {
	f := &Filters{}
	f.States("active", "inactive")
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestNotStates(t *testing.T) {
	f := &Filters{}
	f.NotStates("active", "inactive")
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

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

func TestSides(t *testing.T) {
	f := &Filters{}
	f.Sides(tSide.Buy, tSide.Sell)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestObjectId(t *testing.T) {
	mockID := &MockInterID{}
	f := &Filters{}
	f.ObjectId("field", mockID)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestObjectId_in(t *testing.T) {
	mockID := &MockInterID{}
	f := &Filters{}
	f.ObjectId_in("field", mockID, mockID)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

func TestObjectId_gt(t *testing.T) {
	mockID := &MockInterID{}
	f := &Filters{}
	f.ObjectId_gt("field", mockID)
	if len(f.InterOper.Filters()) == 0 {
		t.Errorf("Expected filter to be added")
	}
}

// MockInterID is a mock implementation of the InterID interface for testing purposes
type MockInterID struct{}

func (m *MockInterID) GetID() interface{} { return "mockID" }
