package main

func main() {
	println(f(1))
	println(g(1))
}

func f(x int) (_, __ int) {
	_, __ = x, x
	return
}

func g(x int) (_, _ int) {
	_, _ = x, x
	return
}
