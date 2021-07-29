# Boolean Algrebra

Set of basic boolean operations to make code more human readable. No more hidden !

They come in two forms:

-   Standard operations which take in and return booleans
-   Lazy operations which take in and return predicates (functions that return a boolean)

## Examples

```golang
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/churmd/booleans/lazyops"
	"github.com/churmd/booleans/ops"
)

func main() {

	// Standard bool operations

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

	// Lazy bool operations

	if lazyops.Not(lazyops.And(alwaysFalse, neverEnding))() {
		fmt.Println("Lazy And short circuit example")
	}

	if lazyops.Or(alwaysTrue, neverEnding)() {
		fmt.Println("Lazy And short circuit example")
	}

	if lazyops.Not(lazyops.Xor(alwaysTrue, alwaysTrue))() {
		fmt.Println("Lazy Not / Xor cannot be short circuited but are included to make chaining calls easy")
	}
}

func alwaysTrue() bool  { return true }

func alwaysFalse() bool { return false }

func neverEnding() bool {
	for i := 0; i < 1000; i++ {
		fmt.Println("never ending work")
		time.Sleep(2 * time.Second)
	}

	return true
}
```

## Should you use this?

Probably not. Using the early return pattern can be applied in most cases where you would want to consider using something like this, rather than setting up all the code beforehand and then only needing some of it.

This little project was initially created becuase I glossed over too many `!` in if statements one day, really wanted to have a `not` keyword and then came across another if statement that chained 7 bool results together with `&&`'s.

The lazy versions were dreamt up later for a bit of fun and to mimic the short circuit feature that `&& / ||` have, which some code relies on.
