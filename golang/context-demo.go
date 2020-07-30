package main // EXAMPLE 1

mport "fmt"
mport "context"

type keyType string

func main() {
	key := keyType("Name")
	ctx := context.WithValue(context.Background(), key, "Tobyy")
	exampleContext(ctx, key)
}

func exampleContext(ctx context.Context, k keyType){
	value := ctx.Value(k)
	if value != nil {
		fmt.Print("The context value is :", value)
		return
	}
	fmt.Print("Ooooops, unable to find the context value")
}

package main // EXAMPLE 2

import  "fmt"
import  "context"
import  "time"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(2*time.Second, cancel)

	sayMyName(ctx, 5*time.Second, "Toby")
}

func sayMyName(ctx context.Context, d time.Duration, name string){
	select {
	case <- time.After(d):
		fmt.Print("Your name is ", name)
	case <-ctx.Done():
		err := ctx.Err()
        fmt.Print(err)
	}
}

package main //EXAMPLE 3

import "context"
import "fmt"
import "time"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	sayMyName(ctx, 5*time.Second, "Toby")
}

func sayMyName(ctx context.Context, d time.Duration, name string) {
	select {
	case <-time.After(d):
		fmt.Print("Your name is ", name)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Print(err)
	}
}
