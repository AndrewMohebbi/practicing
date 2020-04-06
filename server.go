//A practice server!
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	//Do the concurrency stuff
	fmt.Println(do())

	//Do the defer, panic and recover stuff
	doToo()

	//Here we go, the actual server stuff
	http.HandleFunc("/", myHandler)
	http.Handle("/lol/", logWrapper(myHandler))
	http.HandleFunc("/poop/", logWrapper2(myHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func logWrapper(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("We got a funny guy")
		h2 := http.HandlerFunc(h)
		h2.ServeHTTP(w, r)
	})
}

func logWrapper2(h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("We got a sad guy")
		h(w, r)
	}
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
		*sum++
	}

	mutex.Unlock()

	ch <- 0 //Tell do() that you're done
	return
}

func doToo() {
	defer fmt.Println("print last")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!", r)
		}
	}()

	doPanic()

	fmt.Println("print first")
}

func doPanic() {
	panic(fmt.Sprintf("aaaarrgggg"))
}
