package tSide

import (
	"testing"
)

func TestGet(t *testing.T) {
	if Get("buy") != Buy {
		t.Error("Expected Buy, got ", Get("buy"))
	}
	if Get("sell") != Sell {
		t.Error("Expected Sell, got ", Get("sell"))
	}
	if Get("unknown") != DEFAULT {
		t.Error("Expected DEFAULT, got ", Get("unknown"))
	}
}

func TestIsBuy(t *testing.T) {
	if !Buy.IsBuy() {
		t.Error("Expected true, got ", Buy.IsBuy())
	}
	if Sell.IsBuy() {
		t.Error("Expected false, got ", Sell.IsBuy())
	}
}

func TestIsSell(t *testing.T) {
	if !Sell.IsSell() {
		t.Error("Expected true, got ", Sell.IsSell())
	}
	if Buy.IsSell() {
		t.Error("Expected false, got ", Buy.IsSell())
	}
}

func TestString(t *testing.T) {
	if Buy.String() != "Buy" {
		t.Error("Expected Buy, got ", Buy.String())
	}
	if Sell.String() != "Sell" {
		t.Error("Expected Sell, got ", Sell.String())
	}
	if DEFAULT.String() != "Default" {
		t.Error("Expected Default, got ", DEFAULT.String())
	}
}

func TestPosition(t *testing.T) {
	if Buy.Position() != "Long" {
		t.Error("Expected Long, got ", Buy.Position())
	}
	if Sell.Position() != "Short" {
		t.Error("Expected Short, got ", Sell.Position())
	}
	if DEFAULT.Position() != "Default" {
		t.Error("Expected Default, got ", DEFAULT.Position())
	}
}
