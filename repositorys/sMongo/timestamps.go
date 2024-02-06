package sMongo

import (
	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

type TsModel struct {
	mgm.DefaultModel `bson:",inline"`
	UnixTs           int64 `bson:"ts" json:"ts"`
}

// First, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for "ts" first object,
func (c Collection) First(tsFrom, tsTo int64, obj mgm.Model) error {
	c.sort.TsAsc()
	c.filter.Ts(tsFrom, tsTo)
	return c.FindOne(obj)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for first
func (c Collection) FirstTs(tsFrom, tsTo int64) int64 {
	var obj TsModel
	err := c.First(tsFrom, tsTo, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.UnixTs)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for "ts" Last object,
func (c Collection) Last(tsFrom, tsTo int64, obj mgm.Model) error {
	c.sort.TsDesc()
	c.filter.Ts(tsFrom, tsTo)
	return c.FindOne(obj)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for last
func (c *Collection) LastTs(tsFrom, tsTo int64) int64 {
	var obj TsModel
	err := c.Last(tsFrom, tsTo, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.UnixTs)
}

func (c Collection) GetTss() ([]int64, error) {
	c.pipeline = *pipelineTss(c.filter, c.sort, c.limit)

	var docs []TsModel
	err := c.Agregates(&docs)
	if err != nil {
		return nil, err
	}

	tss := make([]int64, len(docs))
	for i, doc := range docs {
		tss[i] = doc.UnixTs
	}

	return tss, nil
}

func pipelineTss(filter Filter, sort Sort, limit int64) *Pipeline {
	p := Pipelines()
	p.Append("$match", filter.Fields)
	p.Append("$sort", sort.Fields)
	p.Append("$limit", limit)
	p.Append("$project", bson.D{{"ts", 1}, {"_id", 0}})
	return p
}
