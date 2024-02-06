package tState

import (
	"fmt"
	"strings"
)

type TradeState int

const (
	Win         = TradeState(1)
	Loss        = TradeState(-1)
	InTrade     = TradeState(2)
	CloseTs     = TradeState(3) // CloseTs: Max In Trade Time permitted from Entry
	CancelTs    = TradeState(4) // CancelTs: Max Excursion Time permitted from Trigger
	CancelPrice = TradeState(5) // CancelPrice: Max Favorable Excursion permitted from Entry
	NotEntry    = TradeState(6) // NotEntry: Never reached Entry Price
	BrakeEven   = TradeState(7) // BrakeEven: Never reached Brake Even Price
	DEFAULT     = TradeState(0)
)

func Get(s string) TradeState {
	switch strings.ToLower(s) {
	case "win", "1":
		return Win
	case "loss", "-1":
		return Loss
	case "intrade", "in_trade", "2":
		return InTrade
	case "closets", "close_ts", "3":
		return CloseTs
	case "cancelts", "cancel_ts", "4":
		return CancelTs
	case "cancelprice", "cancel_price", "5":
		return CancelPrice
	case "notentry", "not_entry", "6":
		return NotEntry
	case "brakeeven", "brake_even", "7":
		return BrakeEven

	default:
		return DEFAULT
	}
}

func (s TradeState) String() string {
	switch s {
	case Win:
		return "Win"
	case Loss:
		return "Loss"
	case InTrade:
		return "InTrade"
	case CloseTs:
		return "CloseTs"
	case CancelTs:
		return "CancelTs"
	case CancelPrice:
		return "CancelPrice"
	case NotEntry:
		return "NotEntry"
	case BrakeEven:
		return "BrakeEven"
	default:
		return "-"
	}
}

func (s TradeState) Log() string {
	return fmt.Sprintf(" ( %s ) ", s.String())
}
