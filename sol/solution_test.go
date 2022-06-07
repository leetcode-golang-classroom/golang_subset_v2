package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	nums := []int{1, 2, 2}
	for idx := 0; idx < b.N; idx++ {
		subsetsWithDup(nums)
	}
}
func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [1,2,2]",
			args: args{nums: []int{1, 2, 2}},
			want: [][]int{{1, 2, 2}, {1, 2}, {1}, {2, 2}, {2}, {}},
		},
		{
			name: "nums = [0]",
			args: args{nums: []int{0}},
			want: [][]int{{0}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
