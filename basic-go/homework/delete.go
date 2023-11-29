package main

import "fmt"

func Delete[T any](idx int, vals []T) []T {
	if idx < 0 || idx >= len(vals) {
		panic("idx不合法")
	}

	vals = append(vals[:idx], vals[idx+1:]...)

	if len(vals) <= cap(vals)/2 {
		res := make([]T, len(vals))
		copy(res, vals)
		return res
	} else {
		return vals
	}
}

func main() {
	var vals = []int{1, 2, 3, 4, 5}
	res := Delete(1, vals)
	fmt.Printf("%v\n", res)
	fmt.Printf("len=%v,cap=%v\n", len(res), cap(res))

	res = Delete(1, res)
	fmt.Printf("%v\n", res)
	fmt.Printf("len=%v,cap=%v\n", len(res), cap(res))

	res = Delete(1, res)
	fmt.Printf("%v\n", res)
	fmt.Printf("len=%v,cap=%v\n", len(res), cap(res))
}
