package main

import (
	"fmt"
	"sort"
)

func main() {
	m := makeMap()
	trap(m) // unsorted printing by key
	fmt.Println()
	solution(m) // sorted printing by key
}

func makeMap() map[int]string {
	return map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}
}

func trap(m map[int]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}

func solution(m map[int]string) {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
