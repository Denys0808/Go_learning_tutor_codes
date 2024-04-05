package main

import (
	"fmt"
)

type IntArray []int

func (arr IntArray) String() string {
	res := ""

	for i := range arr {
		// res = res + strconv.Itoa(arr[i])
		res = fmt.Sprintf("%v%d", res, arr[i])

		if i < len(arr)-1 {
			res = res + ","
		}
	}

	return res
}

func main() {
	a := IntArray{1, 2, 3}
	b := IntArray{4, 5, 6}

	a = append(a, b...)

	fmt.Println(a)
}
