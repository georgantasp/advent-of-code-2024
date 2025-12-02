package main

import "testing"

func Test_counts_run(t *testing.T) {
	tests := []struct {
		name      string
		dialStart int
		direction string
		num       int

		part1out int
		part2out int
		dialout  int
	}{
		{
			name:      "simple left",
			dialStart: 0,
			direction: "L",
			num:       20,
			part1out:  0,
			part2out:  0,
			dialout:   -20,
		},
		{
			name:      "simple left",
			dialStart: 1,
			direction: "L",
			num:       20,
			part1out:  0,
			part2out:  1,
			dialout:   81,
		},
		{
			name:      "simple left",
			dialStart: 1,
			direction: "L",
			num:       100,
			part1out:  0,
			part2out:  1,
			dialout:   1,
		},
		{
			name:      "simple right",
			dialStart: 0,
			direction: "R",
			num:       100,
			part1out:  1,
			part2out:  1,
			dialout:   0,
		},
		{
			name:      "simple right",
			dialStart: 1,
			direction: "R",
			num:       100,
			part1out:  1,
			part2out:  1,
			dialout:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counts{}
			got := c.run(tt.dialStart, tt.direction, tt.num)
			if got != tt.dialout {
				t.Errorf("run() = %v, want %v", got, tt.dialout)
			}
			if c.part1 != tt.part1out {
				t.Errorf("part1 = %v, want %v", c.part1, tt.part1out)
			}
			if c.part2 != tt.part2out {
				t.Errorf("part2 = %v, want %v", c.part2, tt.part2out)
			}
		})
	}
}
