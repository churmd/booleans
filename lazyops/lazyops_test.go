package lazyops_test

import (
	"testing"

	"github.com/churmd/booleans/lazyops"
	"github.com/stretchr/testify/assert"
)

func alwaysTrue() bool { return true }

func alwaysFalse() bool { return false }

func TestPredicateConstruction(t *testing.T) {
	var truePred lazyops.Predicate = func() bool {
		x := 2
		y := 3
		z := 5
		return x + y == z
	}

	assert.True(t, alwaysTrue())
	assert.True(t, truePred())

	var falsePred lazyops.Predicate = func() bool {
		x := 2
		y := 3
		z := 5
		return x + y != z
	}

	assert.False(t, alwaysFalse())
	assert.False(t, falsePred())
}

func TestAndTrueVarients(t *testing.T) {
	trueOptions := [][]lazyops.Predicate{
		{}, 
		{alwaysTrue}, 
		{alwaysTrue, alwaysTrue, alwaysTrue},
	}

	for _, bools := range trueOptions {
		assert.True(t, lazyops.And(bools...))
	}
}

func TestAndFalseVarients(t *testing.T) {
	falseOptions := [][]lazyops.Predicate{
		{alwaysFalse}, 
		{alwaysFalse, alwaysFalse}, 
		{alwaysTrue, alwaysTrue, alwaysFalse},
	}

	for _, bools := range falseOptions {
		assert.False(t, lazyops.And(bools...))
	}
}

func TestOrTrueVarients(t *testing.T) {
	trueOptions := [][]lazyops.Predicate{
		{alwaysTrue}, 
		{alwaysTrue, alwaysFalse}, 
		{alwaysFalse, alwaysTrue}, 
		{alwaysTrue, alwaysTrue, alwaysFalse}, 
		{alwaysFalse, alwaysFalse, alwaysTrue},
	}

	for _, bools := range trueOptions {
		assert.True(t, lazyops.Or(bools...))
	}
}

func TestOrFalseVarients(t *testing.T) {
	falseOptions := [][]lazyops.Predicate{
		{}, 
		{alwaysFalse}, 
		{alwaysFalse, alwaysFalse}, 
		{alwaysFalse, alwaysFalse, alwaysFalse},
	}

	for _, bools := range falseOptions {
		assert.False(t, lazyops.Or(bools...))
	}
}
