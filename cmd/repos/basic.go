package repos

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type InterBasic interface {
	String() string
	Log()

	CUuid() string
}

type ModelBasic struct {
	Uuid string `bson:"uuid" json:"uuid"`
}

func (m *ModelBasic) String() string {
	return fmt.Sprintf("uuid: %s", m.Uuid)
}

func (m *ModelBasic) Log() {
	sLog.Info("Model: %s", m.String())
}

func (m *ModelBasic) CUuid() string {
	return m.Uuid
}

//  models

type InterModelss interface {
	rMongo.InterModel

	InterBasic
}

type Modelss struct {
	rMongo.DefaultModel `bson:",inline"`
	ModelBasic          `bson:",inline"`
}

type Model_A struct {
	Modelss `bson:",inline"`
	Aaa     string `bson:"aaa" json:"aaa"`
}

type Model_B struct {
	Modelss `bson:",inline"`
	Bbb     string `bson:"bbb" json:"bbb"`
}
