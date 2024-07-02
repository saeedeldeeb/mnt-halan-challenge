package main

func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func square(x int) int {
	return x * x
}

func inc(x int) int {
	return x + 1
}
