package indicator

type Base struct {
	c int
}

func NewBase() *Base {
	return &Base{
		c: 0,
	}
}

func (b *Base) Count() int {
	return b.c
}

func (b *Base) Increase() {
	b.c++
}
