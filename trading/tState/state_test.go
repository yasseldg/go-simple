package tState

import (
	"testing"
)

func TestGet(t *testing.T) {
	if Get("win") != Win {
		t.Error("Expected Win, got ", Get("win"))
	}
	if Get("loss") != Loss {
		t.Error("Expected Loss, got ", Get("loss"))
	}
	if Get("intrade") != InTrade {
		t.Error("Expected InTrade, got ", Get("intrade"))
	}
	if Get("closets") != CloseTs {
		t.Error("Expected CloseTs, got ", Get("closets"))
	}
	if Get("cancelts") != CancelTs {
		t.Error("Expected CancelTs, got ", Get("cancelts"))
	}
	if Get("cancelprice") != CancelPrice {
		t.Error("Expected CancelPrice, got ", Get("cancelprice"))
	}
	if Get("notentry") != NotEntry {
		t.Error("Expected NotEntry, got ", Get("notentry"))
	}
	if Get("brakeeven") != BrakeEven {
		t.Error("Expected BrakeEven, got ", Get("brakeeven"))
	}
	if Get("unknown") != DEFAULT {
		t.Error("Expected DEFAULT, got ", Get("unknown"))
	}
}

func TestGetWithDifferentInputs(t *testing.T) {
	if Get("1") != Win {
		t.Error("Expected Win, got ", Get("1"))
	}
	if Get("in_trade") != InTrade {
		t.Error("Expected InTrade, got ", Get("in_trade"))
	}
	if Get("close_ts") != CloseTs {
		t.Error("Expected CloseTs, got ", Get("close_ts"))
	}
	if Get("cancel_ts") != CancelTs {
		t.Error("Expected CancelTs, got ", Get("cancel_ts"))
	}
	if Get("cancel_price") != CancelPrice {
		t.Error("Expected CancelPrice, got ", Get("cancel_price"))
	}
	if Get("not_entry") != NotEntry {
		t.Error("Expected NotEntry, got ", Get("not_entry"))
	}
	if Get("brake_even") != BrakeEven {
		t.Error("Expected BrakeEven, got ", Get("brake_even"))
	}
	if Get("random") != DEFAULT {
		t.Error("Expected DEFAULT, got ", Get("random"))
	}
}

func TestStringMethod(t *testing.T) {
	if Win.String() != "Win" {
		t.Error("Expected Win, got ", Win.String())
	}
	if Loss.String() != "Loss" {
		t.Error("Expected Loss, got ", Loss.String())
	}
	if InTrade.String() != "InTrade" {
		t.Error("Expected InTrade, got ", InTrade.String())
	}
	if CloseTs.String() != "CloseTs" {
		t.Error("Expected CloseTs, got ", CloseTs.String())
	}
	if CancelTs.String() != "CancelTs" {
		t.Error("Expected CancelTs, got ", CancelTs.String())
	}
	if CancelPrice.String() != "CancelPrice" {
		t.Error("Expected CancelPrice, got ", CancelPrice.String())
	}
	if NotEntry.String() != "NotEntry" {
		t.Error("Expected NotEntry, got ", NotEntry.String())
	}
	if BrakeEven.String() != "BrakeEven" {
		t.Error("Expected BrakeEven, got ", BrakeEven.String())
	}
	if DEFAULT.String() != "-" {
		t.Error("Expected -, got ", DEFAULT.String())
	}
}

func TestLogMethod(t *testing.T) {
	if Win.Log() != " ( Win ) " {
		t.Error("Expected ' ( Win ) ', got ", Win.Log())
	}
	if Loss.Log() != " ( Loss ) " {
		t.Error("Expected ' ( Loss ) ', got ", Loss.Log())
	}
	if InTrade.Log() != " ( InTrade ) " {
		t.Error("Expected ' ( InTrade ) ', got ", InTrade.Log())
	}
	if CloseTs.Log() != " ( CloseTs ) " {
		t.Error("Expected ' ( CloseTs ) ', got ", CloseTs.Log())
	}
	if CancelTs.Log() != " ( CancelTs ) " {
		t.Error("Expected ' ( CancelTs ) ', got ", CancelTs.Log())
	}
	if CancelPrice.Log() != " ( CancelPrice ) " {
		t.Error("Expected ' ( CancelPrice ) ', got ", CancelPrice.Log())
	}
	if NotEntry.Log() != " ( NotEntry ) " {
		t.Error("Expected ' ( NotEntry ) ', got ", NotEntry.Log())
	}
	if BrakeEven.Log() != " ( BrakeEven ) " {
		t.Error("Expected ' ( BrakeEven ) ', got ", BrakeEven.Log())
	}
	if DEFAULT.Log() != " ( - ) " {
		t.Error("Expected ' ( - ) ', got ", DEFAULT.Log())
	}
}
