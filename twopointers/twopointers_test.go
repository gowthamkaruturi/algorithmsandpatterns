package twopointers

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid pallindrome string",
			args: args{"raceacar"},
			want: false,
		},
		{
			name: "invalid pallindrome string",
			args: args{"racecar"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfThreeNumbers(t *testing.T) {
	type args struct {
		nums      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: " valid test case",
			args: args{[]int{1, -1, 0}, 0},
			want: []int{-1, 0, 1},
		},
		{name: " invalid test case",
			args: args{[]int{3, 7, 1, 2, 8, 4, 5}, 21},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfThreeNumbers(tt.args.nums, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfThreeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainerWithMostWater(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainerWithMostWater(tt.args.a); got != tt.want {
				t.Errorf("ContainerWithMostWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductOfArrayOtherThanItself(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "valid ",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: []int{120, 60, 40, 30, 24},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductOfArrayOtherThanItself(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductOfArrayOtherThanItself() = %v, want %v", got, tt.want)
			}
		})
	}
}
