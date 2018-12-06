package main
/*
Ruby之父松本行弘 《代码的未来》中关于go的一个例子

这个例子的算法及目的是通过通道将多个goroutine连接起来，这些goroutine分别将值加1，并传给下一个goroutine。
向最开始的通道写入0，则返回由goroutine链所生成的goroutine个数。这个程序中我们生成了10万个goroutine。
*/
import "fmt"

const ngoroutine = 100000;

func f(left, right chan int) {
    left <- 1 + <- right
}

func main() {
    leftmost := make(chan int);

    var left, right chan int = nil, leftmost;
    for i := 0; i < ngoroutine; i++ {
        left, right = right, make(chan int);
        go f(left, right);
    }

    right <- 0;
    x := <-leftmost;
    fmt.Println(x);
}
