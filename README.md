# Boolean Algrebra

Set of basic boolean operations to make code more human readable.

Ever wasted 30 minutes of your life trying to understand how a function is working? Only to go through it with someone else and realise you have simply been missing the negation of an if condition `if !someCoolFunc() { ... }`. Or come across a complex series of nested and's/or's, tried to refactor them out into separate conditions only to end up with failing tests and less of an understanding then when you started?

This library helps to ease these pains by providing named versions of the basic boolean operations so they are harder to gloss over and easier to read.

## Examples

```golang
package main

import (
	"fmt"
	"strings"

	ops "github.com/churmd/booleans/ops"
)

func alwaysTrue() bool  { return true }
func alwaysFalse() bool { return false }

func main() {
	if ops.Not(false) {
		fmt.Println("Not example")
	}

	if ops.And(true, 1+0 == 1, alwaysTrue()) {
		fmt.Println("And example")
	}

	if ops.Or(false, 2+2 == 4, alwaysFalse()) {
		fmt.Println("Or example")
	}

	exampleString := "Hello, world!"
	containsHello := strings.Contains(exampleString, "Hello")
	containsBye := strings.Contains(exampleString, "Bye")
	if ops.And(len(exampleString) == 13, ops.Or(containsHello, containsBye)) {
		fmt.Println("Combining operations example")
	}
}
```
