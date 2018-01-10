package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
}

type F func(F) func(int) int

func Y(le func(func(int) int) func(int) int) func(int) int {
	return func(f F) func(int) int {
		return f(f)
	}(func(f F) func(int) int {
		return le(func(x int) int {
			return f(f)(x)
		})
	})
}

var fibonacci = Y(func(fib func(int) int) func(int) int {
	return func(x int) int {
		switch x {
		case 0, 1:
			return 1
		default:
			return fib(x-1) + fib(x-2)
		}
	}
})
