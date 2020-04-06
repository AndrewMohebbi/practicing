//A practice server!
package main

import (
	"flag"
	"fmt"
	"practicing/functions"
	"practicing/pelicans"
	"sync"
)

var mutex = &sync.Mutex{}

var optional = flag.String("opt", "Hello!", "Tell what to print")

func main() {
	//Print the statement
	flag.Parse()
	fmt.Println(functions.Flip(*optional))

	//Make the pelicans speak
	p := pelicans.NewPelican("karen", "Kaaaahh")
	fmt.Println(p.Speak())
	fmt.Println(*p)

	fmt.Println(do())
}

func do() int {

	ch1 := make(chan int)
	ch2 := make(chan int)
	var sum int

	go adder(ch1, &sum)
	go adder(ch2, &sum)

	done1, done2 := <-ch1, <-ch2
	fmt.Println(done1, done2)
	return sum
}

func adder(ch chan int, sum *int) {
	mutex.Lock()

	for i := 0; i < 10000; i++ {
		*sum += 1
	}

	mutex.Unlock()

	ch <- 0 //Tell do() that you're done
	return
}
