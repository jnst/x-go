package leetcode_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jnst/x-go/leetcode"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "index out of range",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.args.nums1
			leetcode.Merge(got, tt.args.m, tt.args.nums2, tt.args.n)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Merge() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
