package collection

import (
	"github.com/yasseldg/mgm/v4"
)

type TsModel struct {
	mgm.DefaultModel `bson:",inline"`
	Ts               int64 `bson:"ts" json:"ts"`
}

// First, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for "ts" first object,
func (c *Full) First(ts_from, ts_to int64, obj mgm.Model) error {
	c.sort.TsAsc()
	c.filter.Ts(ts_from, ts_to)
	return c.FindOne(obj)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for first
func (c *Full) FirstTs(ts_from, ts_to int64) int64 {
	var obj TsModel
	err := c.First(ts_from, ts_to, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.Ts)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for "ts" Last object,
func (c *Full) Last(ts_from, ts_to int64, obj mgm.Model) error {
	c.sort.TsDesc()
	c.filter.Ts(ts_from, ts_to)
	return c.FindOne(obj)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for last
func (c *Full) LastTs(ts_from, ts_to int64) int64 {
	var obj TsModel
	err := c.Last(ts_from, ts_to, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.Ts)
}

// func (c *Full) GetTss() ([]int64, error) {
// 	p, err := pipelineTss(c.filter, c.sort, c.limit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	c.pipeline = *p

// 	var docs []TsModel
// 	err = c.Agregates(&docs)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tss := make([]int64, len(docs))
// 	for i, doc := range docs {
// 		tss[i] = doc.Ts
// 	}

// 	return tss, nil
// }

// func pipelineTss(filter rFilter.Filters, sort rSort.Sorts, limit int64) (*Pipeline, error) {

// 	f, err := getFilter(filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := getSort(sort)
// 	if err != nil {
// 		return nil, err
// 	}

// 	p := Pipelines()
// 	p.Append("$match", f.getFields)
// 	p.Append("$sort", s.getFields)
// 	p.Append("$limit", limit)
// 	p.Append("$project", bson.D{{"ts", 1}, {"_id", 0}})
// 	return p, nil
// }
