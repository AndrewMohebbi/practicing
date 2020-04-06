//Package pelicans defines the pelican class
package pelicans

import (
	"fmt"
)

//Pelican is a pelican!
type Pelican struct {
	call string
	name string
}

//Pelicaner does what pelicaners do
type Pelicaner interface {
	Speak() string
}

//NewPelican makes a new pelican
func NewPelican(name, call string) *Pelican {
	return &Pelican{name: name, call: call}
}

//Speak returns the pelican's call
func (p Pelican) Speak() string {
	return p.call
}

//String replaces the Stringer interface function, for Print()
func (p Pelican) String() string {
	return fmt.Sprintf("%s: %s!!", p.name, p.call)
}
