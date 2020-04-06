/*A practice server!*/
package main

import (
	"flag"
	"fmt"
	"practicing/functions"
	"practicing/pelicans"
)

var optional = flag.String("opt", "Hello!", "Tell what to print")

func main() {
	flag.Parse()
	fmt.Println(functions.Flip(*optional))

	p := pelicans.NewPelican("karen", "Kaaaahh")
	fmt.Println(p.Speak())
	//p2 := Stringer(*p)
	fmt.Println(*p)
}
