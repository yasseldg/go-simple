package tSide

import (
	"fmt"
	"strings"
)

type Side int

const (
	Buy     = Side(1)
	Sell    = Side(-1)
	DEFAULT = Side(0)
)

func Get(s string) Side {
	switch strings.ToLower(s) {
	case "buy", "long", "1":
		return Buy
	case "sell", "short", "-1":
		return Sell
	default:
		return DEFAULT
	}
}

func (s Side) IsBuy() bool {
	return s == Buy
}

func (s Side) IsSell() bool {
	return s == Sell
}

func (s Side) String() string {
	switch s {
	case Buy:
		return "Buy"
	case Sell:
		return "Sell"
	default:
		return "Default"
	}
}

func (s Side) Position() string {
	switch s {
	case Buy:
		return "Long"
	case Sell:
		return "Short"
	default:
		return "Default"
	}
}

func (s Side) ForLog() string {
	return fmt.Sprintf(" ( %s ) ", s.String())
}

func (s Side) ForLogPosition() string {
	return fmt.Sprintf(" ( %s ) ", s.Position())
}
