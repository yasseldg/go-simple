package tSide

type Inter interface {
	Side() Side
	Switch() Side

	IsBuy() bool
	IsSell() bool
	IsDefault() bool

	String() string
	Position() string
	ForLog() string
	ForLogPosition() string
}
