package dTs

import (
	"testing"
)

type mockInter struct {
	ts int64
}

func (m *mockInter) Ts() int64 {
	return m.ts
}

func TestTs(t *testing.T) {
	tests := []struct {
		name  string
		input Inter
		want  int64
	}{
		{
			name:  "Valid Ts",
			input: &mockInter{ts: 1234567890},
			want:  1234567890,
		},
		{
			name:  "Zero Ts",
			input: &mockInter{ts: 0},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Ts(); got != tt.want {
				t.Errorf("Ts() = %v, want %v", got, tt.want)
			}
		})
	}
}
