package solutions

import "testing"

func Test_solutionOne(t *testing.T) {
	type args struct {
		LIMIT int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pre", args{10}, 23},
		{"solution", args{1000}, 233168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionOne(tt.args.LIMIT); got != tt.want {
				t.Errorf("SolutionOne(%v) = %v, want %v", tt.args.LIMIT, got, tt.want)
			} else {
				t.Logf("SolutionOne(%v) = %v", tt.args.LIMIT, got)
			}
		})
	}
}
