package tExchange

type Inter interface {
	// Name returns the exchange name
	Name() string
	IsValid() bool
}
