package main

import "fmt"

func main() {
	var (
		m, n   int
		a1, a2 []int
	)

	m, a1 = 1, []int{4, 0, 0, 0, 0, 0}
	n, a2 = 5, []int{1, 2, 3, 5, 6}
	merge(a1, m, a2, n)
	fmt.Println(a1)
}

func merge(a1 []int, m int, a2 []int, n int) {
	if m == 0 {
		copy(a1, a2)
		return
	}

	if n == 0 {
		return
	}

	var j int

	for i := range m + n {
		// prevent j from panicking
		if j >= n {
			continue
		}

		// gets array index value
		a, b := a1[i], a2[j]

		// if second array value is less than the first array value, or
		// the index i is less than the maximum possible walked distance (m + j) then
		// do a index push and shift on the array
		if b < a || (a == 0 && i >= m+j) {
			copy(a1[i:], append([]int{b}, a1[i:]...))
			j++
		}
	}
}
