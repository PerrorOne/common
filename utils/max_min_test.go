package utils

import "testing"

func TestMax(t *testing.T) {
	type args struct {
		d []int64
	}
	tests := []struct {
		name      string
		args      args
		wantMax   int64
		wantIndex int
	}{
		{"NOT MAX", args{[]int64{1, 1, 1, 1, 1, 1}}, 1, 0},
		{"MAX", args{[]int64{1, 1, 1, 1, 1, 5}}, 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotIndex := Max(tt.args.d)
			if gotMax != tt.wantMax {
				t.Errorf("Max() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("Max() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
