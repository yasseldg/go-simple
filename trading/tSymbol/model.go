package tSymbol

import "github.com/yasseldg/go-simple/repos/rMongo"

// Model

type model struct {
	Common `bson:",inline"`

	M_precision   int     `bson:"p" json:"p"`
	M_launch_time int64   `bson:"l_t" json:"l_t"`
	M_min_order   float64 `bson:"m_o" json:"m_o"`
	M_location    string  `bson:"l" json:"l"`

	M_config rMongo.M `bson:"cfg" json:"cfg"`
}

func (b *model) Precision() int {
	return b.M_precision
}

func (b *model) LaunchTime() int64 {
	return b.M_launch_time
}

func (m *model) MinOrder() float64 {
	return m.M_min_order
}

func (m *model) Location() string {
	return m.M_location
}

func (m *model) GetConfig(config any) error {
	return rMongo.BsonUnmarshal(m.M_config, config)
}

// set methods

func (s *model) SetPrecision(prec int) {
	s.M_precision = prec
}

func (s *model) SetLaunchTime(launch_time int64) {
	s.M_launch_time = launch_time
}

func (s *model) SetMinOrder(min_order float64) {
	s.M_min_order = min_order
}

func (s *model) SetLocation(location string) {
	s.M_location = location
}

func (m *model) SetConfig(config any) error {

	bson_config, err := rMongo.BsonMarshal(config)
	if err != nil {
		return err
	}
	m.M_config = bson_config

	return nil
}
