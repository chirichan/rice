package rice

import (
	"reflect"
	"testing"
)

func TestRemoveItem(t *testing.T) {
	type args struct {
		items []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{2, 5, 7, 9}, 0}, []int{5, 7, 9}},
		{"", args{[]int{2, 5, 7, 9}, 1}, []int{2, 7, 9}},
		{"", args{[]int{2, 5, 7, 9}, 2}, []int{2, 5, 9}},
		{"", args{[]int{2, 5, 7, 9}, 3}, []int{2, 5, 7}},
		{"", args{[]int{2, 5, 7, 9}, 4}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveItem(tt.args.items, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeDuplicate(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{2, 5, 9, 9}}, []int{2, 5, 9}},
		{"", args{[]int{2}}, []int{2}},
		{"", args{[]int{2, 2}}, []int{2}},
		{"", args{[]int{2, 2, 7, 7}}, []int{2, 7}},
		{"", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeDuplicate(tt.args.items); len(got) != len(tt.want) {
				t.Errorf("DeDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
