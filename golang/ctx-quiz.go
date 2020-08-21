package main

import (
	"context"
	"fmt"
)

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}

func main() {
	ctx := context.TODO()
	f(ctx)
	fmt.Println(ctx.Value("foo"))
}
