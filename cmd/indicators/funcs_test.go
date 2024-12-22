package indicators

import (
	"testing"

	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
)

type mockMongo struct{}

func (m *mockMongo) GetColl(ctx context.Context, db, collType, symbol, collName string) (rMongo.InterRepo, error) {
	return &mockRepo{}, nil
}

func (m *mockMongo) SetDebug(debug bool) {}

func (m *mockMongo) Log() {}

type mockRepo struct{}

func (r *mockRepo) Log() {}

func (r *mockRepo) FindOne(filter interface{}, result interface{}) error {
	return nil
}

func (r *mockRepo) Create(doc interface{}) error {
	return nil
}

func (r *mockRepo) Upsert(filter interface{}, doc interface{}) error {
	return nil
}

func (r *mockRepo) Filters(filters ...interface{}) rMongo.InterRepo {
	return r
}

func (r *mockRepo) Sorts(sorts ...interface{}) rMongo.InterRepo {
	return r
}

func (r *mockRepo) Clone() rMongo.InterRepo {
	return r
}

func (r *mockRepo) Find(result interface{}) error {
	return nil
}

func (r *mockRepo) FindById(id rMongo.ObjectID, result interface{}) error {
	return nil
}

func (r *mockRepo) GetColl(ctx context.Context, db, collType, symbol, collName string) (rMongo.InterRepo, error) {
	return r, nil
}

func (r *mockRepo) SetDebug(debug bool) {}

func TestRun(t *testing.T) {
	mongo := &mockMongo{}
	Run(mongo)
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		indicator string
		expected  Indicator
	}{
		{"RSI", "RSI", tIndicator.NewRSIcandle(14)},
		{"BBands", "BBands", tIndicator.NewBBcandle(30, 2)},
		{"AvgATR", "AvgATR", tIndicator.NewAvgATR(10)},
		{"SmATR", "SmATR", tIndicator.NewSmoothedATR(14)},
		{"ADX", "ADX", tIndicator.NewADX(14)},
		{"SuperTrend", "SuperTrend", tIndicator.NewSuperTrend(10, 3, false)},
		{"PriceAction", "PriceAction", tIndicator.NewPriceAction()},
		{"SuperTrendIter", "SuperTrendIter", nil},
		{"Unknown", "Unknown", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := get(tt.indicator)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRunIndicator(t *testing.T) {
	mongo := &mockMongo{}
	indicator := tIndicator.NewRSIcandle(14)
	run(indicator, mongo, "BYBIT_BTCUSDT", tInterval.Interval_D)
}

func TestConfig(t *testing.T) {
	mongo := &mockMongo{}
	err := config(mongo, "BYBIT_BTCUSDT", tInterval.Interval_D)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
