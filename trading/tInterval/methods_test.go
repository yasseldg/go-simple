package tInterval

import (
	"testing"
)

func TestIsDefault(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want bool
	}{
		{"Default", Interval_DEFAULT, true},
		{"Not Default", Interval_1m, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsDefault(); got != tt.want {
				t.Errorf("IsDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want string
	}{
		{"Default", Interval_DEFAULT, ""},
		{"Not Default", Interval_1m, "1m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinutes(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want int64
	}{
		{"1m", Interval_1m, 1},
		{"5m", Interval_5m, 5},
		{"15m", Interval_15m, 15},
		{"1h", Interval_1h, 60},
		{"4h", Interval_4h, 240},
		{"D", Interval_D, 1440},
		{"W", Interval_W, 10080},
		{"Default", Interval_DEFAULT, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Minutes(); got != tt.want {
				t.Errorf("Minutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeconds(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want int64
	}{
		{"1m", Interval_1m, 60},
		{"5m", Interval_5m, 300},
		{"15m", Interval_15m, 900},
		{"1h", Interval_1h, 3600},
		{"4h", Interval_4h, 14400},
		{"D", Interval_D, 86400},
		{"W", Interval_W, 604800},
		{"Default", Interval_DEFAULT, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Seconds(); got != tt.want {
				t.Errorf("Seconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMilliSeconds(t *testing.T) {
	tests := []struct {
		name string
		i    Interval
		want int64
	}{
		{"1m", Interval_1m, 60000},
		{"5m", Interval_5m, 300000},
		{"15m", Interval_15m, 900000},
		{"1h", Interval_1h, 3600000},
		{"4h", Interval_4h, 14400000},
		{"D", Interval_D, 86400000},
		{"W", Interval_W, 604800000},
		{"Default", Interval_DEFAULT, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.MilliSeconds(); got != tt.want {
				t.Errorf("MilliSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
