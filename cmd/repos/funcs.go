package repos

import (
	"context"
	"fmt"
	"os"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tSide"
)

var (
	_coll rMongo.InterColl
)

func Run(mongo rMongo.Inter) {

	testEnv()

	mongo.SetDebug(true)

	err := config(mongo)
	if err != nil {
		sLog.Error("repos.Run(): config(): %s", err)
		return
	}

	strategie := &Model{
		ModelBasic: ModelBasic{
			Uuid: "uuid_7"},
		Name:   "name",
		Code:   "code",
		Symbol: "symbol",
		Side:   tSide.Buy,
	}

	err = load(strategie)
	if err != nil {
		sLog.Error("load(): %s", err)
	}

	model_A()
	model_B()
}

func config(_mongo rMongo.Inter) error {

	ctx := context.Background()

	var err error
	_coll, err = _mongo.GetColl(ctx, "strategies", "WRITE", "bot_test", "strat_test", Indexes()...)
	if err != nil {
		return fmt.Errorf("GetColl(): %s", err)
	}
	_coll.Log()

	return nil
}

func testEnv() {
	// set env

	os.Setenv("CONN_strategies", "READ")
	os.Setenv("DB_strategies", "bot_test_1")
	os.Setenv("COLL_strategies", "strat_test_1")
}

func model_A() {
	strategie_a := &Model_A{
		Modelss: Modelss{
			ModelBasic: ModelBasic{
				Uuid: "uuid_4"}},
		Aaa: "aaa_4444",
	}

	err := upsert(strategie_a)
	if err != nil {
		sLog.Error("upsert(): %s", err)
	}
}

func model_B() {
	strategie_a := &Model_B{
		Modelss: Modelss{
			ModelBasic: ModelBasic{
				Uuid: "uuid_4"}},
		Bbb: "bbb_4444",
	}

	err := upsert(strategie_a)
	if err != nil {
		sLog.Error("upsert(): %s", err)
	}
}

func load(strategie InterModel) error {

	doc, err := find(strategie)
	if err != nil {
		sLog.Error("find(): %s", err)
		return create(strategie)
	}

	sLog.Warn("strategie already exists: %s", doc.String())

	return nil
}

func find(strategie InterModel) (*Model, error) {

	filter := NewFilter().Code(strategie.CCode()).Symbol(strategie.CSymbol()).
		State("active").Side(strategie.CSide()).Uuid(strategie.CUuid())

	sLog.Warn("find strategie: %s", filter.String())

	var doc Model
	err := _coll.Filters(filter).FindOne(&doc)
	if err != nil {
		return nil, fmt.Errorf("coll.FindOne(): %s", err)
	}

	sLog.Warn("found strategie: %s", doc.String())

	return &doc, nil
}

func create(strategie InterModel) error {

	sLog.Warn("create strategie: %s", strategie.String())

	strategie.SetState("active")

	strategie.Log()

	_coll.Log()

	err := _coll.Create(strategie)
	if err != nil {
		return fmt.Errorf("coll.Create(): %s", err)
	}

	strategie.Log()

	return nil
}

func upsert(strategie InterModelss) error {

	sLog.Warn("create strategie: %s", strategie.String())

	strategie.Log()

	_coll.Log()

	filter := NewFilter().Uuid(strategie.CUuid())

	err := _coll.Filters(filter).Upsert(strategie)
	if err != nil {
		return fmt.Errorf("coll.Upsert(): %s", err)
	}

	strategie.Log()

	return nil
}
