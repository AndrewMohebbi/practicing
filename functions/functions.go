//Package functions provides some functions to use
package functions

//Flip reverses a string
func Flip(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
