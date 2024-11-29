package rFilter

import "time"

// ----- Dates Filters

func (f *Filters) CreatedAt_gt(t time.Time) Inter { f.Time_gt("c_at", t); return f }

func (f *Filters) CreatedAt_lt(t time.Time) Inter { f.Time_lt("c_at", t); return f }

func (f *Filters) UpdatedAt_gt(t time.Time) Inter { f.Time_gt("u_at", t); return f }

func (f *Filters) UpdatedAt_lt(t time.Time) Inter { f.Time_lt("u_at", t); return f }
