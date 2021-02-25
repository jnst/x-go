package leetcode_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jnst/x-go/leetcode"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := map[string]struct {
		args args
		want []int
	}{
		"Example 1": {
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		"Example 2": {
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		"Example 3": {
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := leetcode.TwoSum(tt.args.nums, tt.args.target)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("TwoSum() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
