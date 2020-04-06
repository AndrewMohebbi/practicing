/*A practice server!*/
package main

import (
	"flag"
	"fmt"
	"practicing/functions"
)

var optional = flag.String("opt", "Hello!", "Tell what to print")

func main() {
	flag.Parse()
	fmt.Println(functions.Flip(*optional))
}
