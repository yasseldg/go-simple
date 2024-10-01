package main

import (
	"errors"
	"fmt"

	"github.com/yasseldg/go-simple/cmd/components"
	"github.com/yasseldg/go-simple/cmd/indicators"
	"github.com/yasseldg/go-simple/cmd/iters"
	"github.com/yasseldg/go-simple/cmd/repos"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
)

var (
	_mongo rMongo.Inter
)

func Init() {
	_mongo = rMongo.New()
}

func main() {

	clean := sLog.SetByName(sLog.Zap, sLog.LevelDebug, "")
	defer clean()

	Init()

	sLog.Info("Starting...")

	sTime.TimeControl(iters.Csv, "Iters")
}

func testModel() {
	repos.Run(_mongo)
	// repos.RunConcurr(_mongo)
	// repos.RunIter(_mongo)
}

func testIndicators() {
	indicators.Run(_mongo)
}

func testComponents() {
	components.Run(_mongo)
}

func testErrors() {

	errs := make([]error, 0)
	for i := 0; i < 10; i++ {
		errs = append(errs, fmt.Errorf("error %d", i))
	}

	sLog.Error("Errors: %s", errs)

	sLog.Error("Error Join: \n%s", errors.Join(errs...).Error())
}
