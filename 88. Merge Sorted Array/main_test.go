package main

import (
	"reflect"
	"testing"
)

func make(elems ...int) []int {
	a := []int{}
	return append(a, elems...)
}

func Test_merge(t *testing.T) {
	type args struct {
		a1 []int
		m  int
		a2 []int
		n  int
	}
	tests := []struct {
		name    string
		args    args
		expects []int
	}{
		{name: "Case #1", args: args{a1: make(1, 2, 3, 0, 0, 0), m: 3, a2: make(2, 5, 6), n: 3}, expects: make(1, 2, 2, 3, 5, 6)},
		{name: "Case #2", args: args{a1: make(1), m: 1, a2: make(), n: 0}, expects: make(1)},
		{name: "Case #3", args: args{a1: make(0), m: 0, a2: make(1), n: 1}, expects: make(1)},
		{name: "Case #4", args: args{a1: make(1, 0), m: 1, a2: make(2), n: 1}, expects: make(1, 2)},
		{name: "Case #5", args: args{a1: make(2, 0), m: 1, a2: make(1), n: 1}, expects: make(1, 2)},
		{name: "Case #6", args: args{a1: make(-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0), m: 5, a2: make(-1, -1, 0, 0, 1, 2), n: 6}, expects: make(-1, -1, -1, 0, 0, 0, 0, 0, 1, 2, 3)},
		{name: "Case #7", args: args{a1: make(4, 0, 0, 0, 0, 0), m: 1, a2: make(1, 2, 3, 5, 6), n: 5}, expects: make(1, 2, 3, 4, 5, 6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.a1, tt.args.m, tt.args.a2, tt.args.n)

			if !reflect.DeepEqual(tt.args.a1, tt.expects) {
				t.Errorf("Arrays mismatch: Expected `%+v`, got `%+v`", tt.expects, tt.args.a1)
			}
		})
	}
}
