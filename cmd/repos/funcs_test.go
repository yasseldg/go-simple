package repos

import (
	"os"
	"testing"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tSide"
)

func TestRun(t *testing.T) {
	mongo := rMongo.New()
	Run(mongo)
}

func TestConfig(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
	}
}

func TestTestEnv(t *testing.T) {
	testEnv()
	if os.Getenv("CONN_strategies") != "READ" {
		t.Errorf("Expected CONN_strategies to be READ, got %s", os.Getenv("CONN_strategies"))
	}
	if os.Getenv("DB_strategies") != "bot_test_1" {
		t.Errorf("Expected DB_strategies to be bot_test_1, got %s", os.Getenv("DB_strategies"))
	}
	if os.Getenv("COLL_strategies") != "strat_test_1" {
		t.Errorf("Expected COLL_strategies to be strat_test_1, got %s", os.Getenv("COLL_strategies"))
	}
}

func TestModelA(t *testing.T) {
	model_A()
}

func TestModelB(t *testing.T) {
	model_B()
}

func TestLoad(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
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
		t.Errorf("load() error = %v", err)
	}
}

func TestFind(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
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

	_, err = find(strategie)
	if err != nil {
		t.Errorf("find() error = %v", err)
	}
}

func TestCreate(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
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

	err = create(strategie)
	if err != nil {
		t.Errorf("create() error = %v", err)
	}
}

func TestUpsert(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
		return
	}

	strategie := &Model_A{
		Modelss: Modelss{
			ModelBasic: ModelBasic{
				Uuid: "uuid_4"}},
		Aaa: "aaa_4444",
	}

	err = upsert(strategie)
	if err != nil {
		t.Errorf("upsert() error = %v", err)
	}
}
