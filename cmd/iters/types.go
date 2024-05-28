package iters

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
)

func Tests() {
	floats()
	ints()
}

func floats() {
	run(sFloats.NewBaseIter(0, 11, 1.1, 12.3, 45, 6.6), "v1")

	run(sFloats.NewBaseIter(3, 3, 1.2, 12, 45.7, 7.2), "v2")

	run(sFloats.NewBaseIter(3, 0, 1, 12, 45, 6), "v3")

	run(sFloats.NewBaseIter(0, 0, 0.1, 12, 45, 6), "v4")

	run(sFloats.NewBaseIter(0, 10, 1), "v5")

	run(sFloats.NewBaseIter(0, 10, 0), "v6")

	run(sFloats.NewBaseIter(0, 0, 0), "v7")

	run(sFloats.NewBaseIter(0, 0, 0, 0), "v8")

	run(sFloats.NewBaseIter(10, 0, 1), "v9")

	run(sFloats.NewBaseIter(10, 0, -0.15), "v10")
}

func ints() {
	run(sFloats.NewBaseIter(0, 10, 1, 12, 45, 6), "v1")

	run(sFloats.NewBaseIter(3, 3, 1, 12, 45, 6), "v2")

	run(sFloats.NewBaseIter(3, 0, 1, 12, 45, 6), "v3")

	run(sFloats.NewBaseIter(0, 0, 1, 12, 45, 6), "v4")

	run(sFloats.NewBaseIter(0, 10, 1), "v5")

	run(sFloats.NewBaseIter(0, 10, 0), "v6")

	run(sFloats.NewBaseIter(0, 0, 0), "v7")

	run(sFloats.NewBaseIter(0, 0, 0, 0), "v8")
}

func run(iter dIter.Inter, name string) {
	for iter.Next() {
		iter.Log(name)
	}
	println()
}
