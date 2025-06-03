package collection

import (
	"context"

	"github.com/yasseldg/mgm/v4"
)

type TsModel struct {
	mgm.DefaultModel `bson:",inline"`
	Ts               int64 `bson:"ts" json:"ts"`
}

// First, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for "ts" first object,
func (c *Full) First(ts_from, ts_to int64, model mgm.Model) error {
	return c.FirstWithCtx(mgm.Ctx(), ts_from, ts_to, model)
}

func (c *Full) FirstWithCtx(ctx context.Context,
	ts_from, ts_to int64, model mgm.Model) error {
	c.sort.TsAsc()
	c.filter.Ts(ts_from, ts_to)
	return c.FindOneWithCtx(ctx, model)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for first
func (c *Full) FirstTs(ts_from, ts_to int64) int64 {
	return c.FirstTsWithCtx(mgm.Ctx(), ts_from, ts_to)
}

func (c *Full) FirstTsWithCtx(ctx context.Context,
	ts_from, ts_to int64) int64 {
	var model TsModel
	err := c.FirstWithCtx(ctx, ts_from, ts_to, &model)
	if err != nil {
		return 0
	}
	return int64(model.Ts)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for "ts" Last object,
func (c *Full) Last(ts_from, ts_to int64, model mgm.Model) error {
	return c.LastWithCtx(mgm.Ctx(), ts_from, ts_to, model)
}

func (c *Full) LastWithCtx(ctx context.Context,
	ts_from, ts_to int64, model mgm.Model) error {
	c.sort.TsDesc()
	c.filter.Ts(ts_from, ts_to)
	return c.FindOneWithCtx(ctx, model)
}

// Last, $gte: ts_from  $lt: ts_to, ts_from = ts_to = 0 for last
func (c *Full) LastTs(ts_from, ts_to int64) int64 {
	return c.LastTsWithCtx(mgm.Ctx(), ts_from, ts_to)
}

func (c *Full) LastTsWithCtx(ctx context.Context,
	ts_from, ts_to int64) int64 {
	var model TsModel
	err := c.LastWithCtx(ctx, ts_from, ts_to, &model)
	if err != nil {
		return 0
	}
	return model.Ts
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
