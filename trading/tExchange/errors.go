package tExchange

import "errors"

var (
	ErrInvalidExchange = errors.New("invalid exchange")
	ErrEmptySymbols    = errors.New("empty symbols")
)
