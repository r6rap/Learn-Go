package function

import "fmt"

//recursive function: function yang memanggil dirinya sendiri

func FaktorialRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * FaktorialRecursive(n-1)
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func FaktorialLoop(n int) {
	h := 1
	for i := n; i > 0; i--{
		h *= i //sama seperti h = h*i
	}
	fmt.Println(h)
}