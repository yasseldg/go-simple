package rFilter

// ----- States Filters

func (f *Filters) States(states ...string) Inter { f.String_in("st", states...); return f }

func (f *Filters) NotStates(states ...string) Inter { f.String_nin("st", states...); return f }
