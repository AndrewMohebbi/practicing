/*A practice server!*/
package main

import (
	"flag"
	"fmt"
)

var optional = flag.String("optional", "Hello!", "Tell what to print")

func main() {
	flag.Parse()
	fmt.Println(*optional)
}
