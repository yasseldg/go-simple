package repos

import (
	"testing"

	"github.com/yasseldg/go-simple/trading/tSide"
)

func TestModel_String(t *testing.T) {
	model := &Model{
		ModelBasic: ModelBasic{Uuid: "test-uuid"},
		Name:       "test-name",
		Code:       "test-code",
		Symbol:     "test-symbol",
		Side:       tSide.Buy,
		State:      "active",
	}
	expected := "uuid: test-uuid .. name: test-name .. code: test-code .. symbol: test-symbol .. side: Buy .. state: active"
	if model.String() != expected {
		t.Errorf("Expected %s, got %s", expected, model.String())
	}
}

func TestModel_CCode(t *testing.T) {
	model := &Model{Code: "test-code"}
	expected := "test-code"
	if model.CCode() != expected {
		t.Errorf("Expected %s, got %s", expected, model.CCode())
	}
}

func TestModel_CSymbol(t *testing.T) {
	model := &Model{Symbol: "test-symbol"}
	expected := "test-symbol"
	if model.CSymbol() != expected {
		t.Errorf("Expected %s, got %s", expected, model.CSymbol())
	}
}

func TestModel_CSide(t *testing.T) {
	model := &Model{Side: tSide.Buy}
	expected := tSide.Buy
	if model.CSide() != expected {
		t.Errorf("Expected %v, got %v", expected, model.CSide())
	}
}
