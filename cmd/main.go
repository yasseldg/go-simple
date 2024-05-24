package main

import (
	"github.com/yasseldg/go-simple/cmd/indicators"
	"github.com/yasseldg/go-simple/cmd/repos"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
)

var (
	_mongo *rMongo.Manager
)

func Init() {
	mongo := rMongo.NewManager()
	_mongo = &mongo
}

func main() {

	clean := sLog.SetByName(sLog.Zap, sLog.LevelInfo, "")
	defer clean()

	Init()

	sLog.Info("Starting...")

	sTime.TimeControl(testIndicators, "Indicators")
}

func testModel() {
	repos.Run(_mongo)
}

func testIndicators() {
	indicators.Run(_mongo)
}
