package repos

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/trading/tSide"
)

var (
	_coll rMongo.Collection
)

func Run(mongo *rMongo.Manager) {
	err := config(mongo)
	if err != nil {
		sLog.Error("repos.Run(): config(): %s", err)
		return
	}

	strategie := Model{
		Uuid:   "uuid",
		Name:   "name",
		Code:   "code",
		Symbol: "symbol",
		Side:   tSide.Buy,
	}

	err = load(strategie)
	if err != nil {
		sLog.Error("load(): %s", err)
	}
}

func config(_mongo *rMongo.Manager) error {
	var err error
	_coll, err = _mongo.GetColl("strategies", "WRITE", "bot", "strategies", Indexes()...)
	if err != nil {
		return fmt.Errorf("GetColl(): %s", err)
	}
	_coll.Log()

	return nil
}

func load(strategie Model) error {

	doc, err := find(strategie)
	if err != nil {
		sLog.Error("find(): %s", err)
		return create(strategie)
	}

	sLog.Warn("strategie already exists: %s", doc.String())

	return nil
}

func find(strategie Model) (*Model, error) {

	filter := NewFilter().Code(strategie.Code).Symbol(strategie.Symbol).State("active").Side(strategie.Side)

	sLog.Warn("find strategie: %s", filter.String())

	var doc Model
	err := _coll.Filters(filter.Filters).FindOne(&doc)
	if err != nil {
		return nil, fmt.Errorf("coll.FindOne(): %s", err)
	}

	sLog.Warn("found strategie: %s", doc.String())

	return &doc, nil
}

func create(strategie Model) error {

	sLog.Warn("create strategie: %s", strategie.String())

	strategie.SetState("active")

	strategie.Log()

	_coll.Log()

	err := _coll.Create(&strategie)
	if err != nil {
		return fmt.Errorf("coll.Create(): %s", err)
	}

	strategie.Log()

	return nil
}
