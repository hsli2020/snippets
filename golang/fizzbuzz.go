// Written by http://xojoc.pw. Public domain.
package main

func makeChan(delay time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Millisecond * delay)
			c <- struct{}{}
		}
	}()
	return c
}

func main() {
	fizz := makeChan(30)
	buzz := makeChan(50)
	fizzBuzz := makeChan(150)
	number := makeChan(10)

	time.Sleep(time.Millisecond * 20)

	for i := 1; i <= 100; i++ {
		select {
		case <-fizzBuzz:
			fmt.Println("FizzBuzz")
			<-fizz
			<-buzz
			<-number
			continue
		default:
		}
		select {
		case <-buzz:
			fmt.Println("Buzz")
			<-number
			continue
		default:
		}
		select {
		case <-fizz:
			fmt.Println("Fizz")
			<-number
			continue
		default:
		}
		select {
		case <-number:
			fmt.Println(i)
		}
	}
}

// Written by http://xojoc.pw. Public domain.
package main

func main() {
	fizz := time.Tick(30 * time.Millisecond)
	buzz := time.Tick(50 * time.Millisecond)
	fizzBuzz := time.Tick(150 * time.Millisecond)
	number := time.Tick(10 * time.Millisecond)

	time.Sleep(20 * time.Millisecond)

	for i := 1; i <= 100; i++ {
		select {
		case <- fizzBuzz:
			fmt.Println("FizzBuzz")
			<-fizz
			<-buzz
			<-number
			continue
		default:
		}
		select {
		case <- fizz:
			fmt.Println("Fizz")
			<-number
			continue
		default:
		}
		select {
		case <- buzz:
			fmt.Println("Buzz")
			<-number
			continue
		default:
		}
		select {
		case <- number:
			fmt.Println(i)
		}
	}
}

// Based on http://xojoc.pw/justcode/6/fizzbuzz.go (public domain)
// Public domain too, since why the heck not.

package main

type Signal struct{}

func main() {
	send, rcv := Multiplex(2)
	fizz := SignalOn(3, rcv[0])
	buzz := SignalOn(5, rcv[1])

	for i := 1; i <= 100; i++ {
		send <- Signal{}

		fizzed := <-fizz
		if fizzed {
			fmt.Print("Fizz")
		}

		buzzed := <-buzz
		if buzzed {
			fmt.Print("Buzz")
		}

		if !(fizzed || buzzed) {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

// Multiplex generates takes a signal on a channel and replicates it out
// across n other channels.
func Multiplex(n int) (chan<- Signal, []<-chan Signal) {
	input := make(chan Signal)
	sendOuts := make([]chan Signal, 0, n)
	outputs := make([]<-chan Signal, 0, n) // Hooray type safety
	for i := 0; i < n; i++ {
		out := make(chan Signal)
		sendOuts = append(sendOuts, out)
		outputs = append(outputs, out)
	}
	go func() {
		// Not going to bother cleaning up, since this goroutine will last the life
		// of the program adding bookkeeping so that the OS doesn't have to clean up
		// our mess is... not the exciting part of fizzbuzz, honestly.
		for {
			<-input
			go func() {
				for _, out := range sendOuts {
					out <- Signal{}
				}
			}()
		}
	}()

	return input, outputs
}

// SignalOn generates a channel that will signal every nth time a signal is sent to start.
func SignalOn(n int, start <-chan Signal) <-chan bool {
	signalOut := make(chan bool, 1)

	go func() {
		var accum *chan Signal
		tmp := make(chan Signal, n-1)
		accum = &tmp
		for {
			<-start
			select {
			case (*accum) <- Signal{}: // Accumulate if possible
			default: // Otherwise signal out and grab a new accumulator
				close(*accum)
				(*accum) = make(chan Signal, n-1)
				signalOut <- true
				continue
			}
			signalOut <- false
		}
	}()

	return signalOut
}
