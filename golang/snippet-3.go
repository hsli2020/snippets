// 类型判断
var val interface{}
val = "foo"
if str, ok := val.(string); ok {
    fmt.Println(str)
}
----------------------------------------
for range time.Tick(time.Second) {
    // do it once a second
}
----------------------------------------
func sum(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

sum(1,2,3)

nums := []int{1,2,3,4}
sum(nums...)
----------------------------------------
func outer() (func() int, int) {
    outer_val := 2
    inner := func() int {
        outer_val += 99
        return outer_val
    }
    return inner, outer_val
}
----------------------------------------
ch := make(chan int)
ch <- 42             // Send message to chan
v := <-ch            // Read message from chan
v, ok := <-ch        // Read, ok is false if closed
for i := range ch {  // Read until closed
    fmt.Println(i)
}
----------------------------------------
ch := make(chan string)
go func() { ch <- "ping" }()
msg := <-ch
----------------------------------------
func ping(ch chan<-string, msg string) {
    ch <- msg
}
func pong(pings <-chan string, pongs chan<-string) {
    msg := <-pings
    pongs <- msg
}
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "message")
    pong(pings, pongs)
    fmt.Println(<-poings)
}
----------------------------------------
c1 := make(chan string)
c2 := make(chan string)

go func() {
    time.Sleep(1*time.Second)
    c1 <- "one"
}()

go func() {
    time.Sleep(1*time.Second)
    c2 <- "two"
}()

for i:=0; i<2; i++ {
    select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
    }
}
----------------------------------------
