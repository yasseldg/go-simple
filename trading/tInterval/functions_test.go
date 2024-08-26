package tInterval

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		input    string
		expected Interval
	}{
		{"M1", Interval_1m},
		{"M5", Interval_5m},
		{"M15", Interval_15m},
		{"H1", Interval_1h},
		{"H4", Interval_4h},
		{"D", Interval_D},
		{"W", Interval_W},
		{"unknown", DEFAULT},
	}

	for _, tc := range testCases {
		if result := Get(tc.input); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestGetMult(t *testing.T) {
	testCases := []struct {
		input            string
		expectedInterval Interval
		expectedMult     int
	}{
		{"M1*2", Interval_1m, 2},
		{"M5*3", Interval_5m, 3},
		{"H1*4", Interval_1h, 4},
		{"H4*5", Interval_4h, 5},
		{"D*6", Interval_D, 6},
		{"W*7", Interval_W, 7},
	}

	for _, tc := range testCases {
		resultInterval, resultMult := GetMult(tc.input)
		if resultInterval != tc.expectedInterval || resultMult != tc.expectedMult {
			t.Errorf("Expected %v and %v, got %v and %v", tc.expectedInterval, tc.expectedMult, resultInterval, resultMult)
		}
	}
}

func TestIsSameMonth(t *testing.T) {
	testCases := []struct {
		ts1, ts2 int64
		expected bool
	}{
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC).Unix(), true},
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC).Unix(), false},
		{time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC).Unix(), false},
		{time.Date(2022, 5, 30, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 5, 31, 0, 0, 0, 0, time.UTC).Unix(), true},
	}

	for _, tc := range testCases {
		if result := IsSameMonth(tc.ts1, tc.ts2); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestStartOfMonth(t *testing.T) {
	testCases := []struct {
		ts       int64
		expected int64
	}{
		{time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 4, 30, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC).Unix()},
	}

	for _, tc := range testCases {
		if result := startOfMonth(tc.ts); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestStartOfNextMonth(t *testing.T) {
	testCases := []struct {
		ts       int64
		expected int64
	}{
		{time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC).Unix()},
		{time.Date(2022, 4, 30, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC).Unix()},
	}

	for _, tc := range testCases {
		if result := startOfNextMonth(tc.ts); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
