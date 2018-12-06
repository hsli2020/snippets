// https://github.com/drgrib/iter
// Package iter provides benchmarking that shows the built-in for clause is the best option for performance discovered so far
package iter

func N(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
