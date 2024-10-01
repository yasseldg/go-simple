package iters

import (
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/files/fCsv"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func Csv() {

	t := time.Now()
	iterB()

	s1 := time.Since(t)

	println()

	t = time.Now()
	iterF()

	s2 := time.Since(t)

	sLog.Info("Time B: %s", s1)
	sLog.Info("Time F: %s", s2)
}

func iterB() {

	iter, err := fCsv.NewIterBatch("../backend/uploads/bybit/10000LADYSUSDT2023-05-11.csv", 10, ',')
	if err != nil {
		sLog.Error("NewIter: %s", err.Error())
		return
	}
	defer iter.CloseFile()

	iter.Log("Csv")

	c := 0
	for iter.Next() {
		c++

		sLog.Info("Item %-4d: %v", c, iter.Item)

		if c > 105 {
			break
		}
	}

	iter.Log("Csv")
}

func iterF() {
	accu := accu{}

	iterFunc, err := fCsv.NewIterFunc("../backend/uploads/bybit/10000LADYSUSDT2023-05-11.csv", ',', accu.Add)
	if err != nil {
		sLog.Error("NewIterFunc: %s", err.Error())
		return
	}
	defer iterFunc.CloseFile()

	iterFunc.Log("Csv")

	iterFunc.Run()

	iterFunc.Log("Csv")
}

type accu struct {
	c int
}

func (a *accu) Add(line []string) error {

	a.c++

	sLog.Info("Line %-4d: %v", a.c, line)

	if a.c < 108 {
		return nil
	}
	return fmt.Errorf("Stop")
}
