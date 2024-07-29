package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "simple test",
			value: -3.14,
			want:  3.14,
		},
		{
			name:  "with positive number",
			value: 3.1,
			want:  3.1,
		},
		{
			name: "with negative zero",
			// сравнение и так -0 == 0
			value: -0,
			want:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := Abs(test.value); res != test.want {
				t.Errorf("Abs() = %f, want %f", res, test.want)
			}
		})
	}
}
