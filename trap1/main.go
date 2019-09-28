package main

import (
	"fmt"
	"sync"
)

func main() {

	vs := serialInts(6)

	trap(vs) // print last value of array
	fmt.Println()

	solution1(vs)
	fmt.Println()

	solution2(vs)
	fmt.Println()
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func trap(vs []int) {
	var wg sync.WaitGroup
	wg.Add(len(vs))
	for _, v := range vs {
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func solution1(vs []int) {
	var wg sync.WaitGroup
	wg.Add(len(vs))
	for _, v := range vs {
		x := v
		go func() {
			fmt.Println(x)
			wg.Done()
		}()
	}
	wg.Wait()
}

func solution2(vs []int) {
	var wg sync.WaitGroup
	wg.Add(len(vs))
	for _, v := range vs {
		go func(x int) {
			fmt.Println(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
