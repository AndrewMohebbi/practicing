//Test for pelicans!
package pelicans

import "testing"

func TestNewPelican(t *testing.T) {

	name := "bob"
	call := "cacaaaaa"

	p := NewPelican(name, call)

	if p.name != name || p.call != call {
		t.Errorf("Something's wrong with those pelicans")
	}

}
