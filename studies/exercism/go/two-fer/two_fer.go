//Package twofer have simple functions that format and return twofer strings
package twofer

import "fmt"

//ShareWith function returns a two-fer string given an name. If name is empty, then its content is replace with 'you'
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
