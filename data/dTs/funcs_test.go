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

func TestGetInters(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  []Inter
	}{
		{
			name:  "Empty slice",
			input: []interface{}{},
			want:  []Inter{},
		},
		{
			name: "Slice with nil values",
			input: []interface{}{
				nil,
				nil,
			},
			want: []Inter{},
		},
		{
			name: "Slice with valid Inter values",
			input: []interface{}{
				&mockInter{ts: 1},
				&mockInter{ts: 2},
			},
			want: []Inter{
				&mockInter{ts: 1},
				&mockInter{ts: 2},
			},
		},
		{
			name: "Mixed slice with nil and valid Inter values",
			input: []interface{}{
				nil,
				&mockInter{ts: 1},
				nil,
				&mockInter{ts: 2},
			},
			want: []Inter{
				&mockInter{ts: 1},
				&mockInter{ts: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInters(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInters() = %v, want %v", got, tt.want)
			}
		})
	}
}
