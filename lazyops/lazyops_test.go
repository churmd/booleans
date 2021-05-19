package lazyops_test

import (
	"testing"

	"github.com/churmd/booleans/lazyops"
	"github.com/stretchr/testify/assert"
)

func alwaysTrue() bool { return true }

func alwaysFalse() bool { return false }

func TestPredicateConstruction(t *testing.T) {
	var a lazyops.Predicate = func() bool {return true}
	var b lazyops.Predicate = func() bool {
		x := 2
		y := 3
		z := 5
		return x + y == z
	}

	assert.True(t, a())
	assert.True(t, b())
}

func TestAndTrueVarients(t *testing.T) {
	trueOptions := [][]lazyops.Predicate{{}, {alwaysTrue}}

	for _, bools := range trueOptions {
		assert.True(t, lazyops.And(bools...))
	}
}

func TestAndFalseVarients(t *testing.T) {
	falseOptions := [][]lazyops.Predicate{{alwaysFalse}}

	for _, bools := range falseOptions {
		assert.False(t, lazyops.And(bools...))
	}
}
