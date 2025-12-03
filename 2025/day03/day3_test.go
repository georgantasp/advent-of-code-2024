package main

import "testing"

func Test_getBankMax(t *testing.T) {
	tests := []struct {
		name string
		bank string
		size int
		want int
	}{
		{
			name: "Test 1",
			bank: "4367133344544422433442443532453333443413434346523142232414343334324332325334334344244542343334544237",
			size: 2,
			want: 77,
		},
		{
			name: "Test 2",
			bank: "987654321111111",
			size: 12,
			want: 987654321111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBankMax(tt.bank, tt.size); got != tt.want {
				t.Errorf("getBankMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
