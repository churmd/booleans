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
		return x+y == z
	}

	assert.True(t, alwaysTrue())
	assert.True(t, truePred())

	var falsePred lazyops.Predicate = func() bool {
		x := 2
		y := 3
		z := 5
		return x+y != z
	}

	assert.False(t, alwaysFalse())
	assert.False(t, falsePred())
}

func TestNotTrueVarient(t *testing.T) {
	assert.True(t, lazyops.Not(alwaysFalse)(), "Lazy Not failed expcted true")
}

func TestNotFalseVarient(t *testing.T) {
	assert.False(t, lazyops.Not(alwaysTrue)(), "Lazy Not failed expcted false")
}

func TestAndTrueVarients(t *testing.T) {
	trueOptions := [][]lazyops.Predicate{
		{},
		{alwaysTrue},
		{alwaysTrue, alwaysTrue, alwaysTrue},
	}

	for _, bools := range trueOptions {
		assert.True(t, lazyops.And(bools...), "Lazy And failed for case %+v, expected true", bools)
	}
}

func TestAndFalseVarients(t *testing.T) {
	falseOptions := [][]lazyops.Predicate{
		{alwaysFalse},
		{alwaysFalse, alwaysFalse},
		{alwaysTrue, alwaysTrue, alwaysFalse},
	}

	for _, bools := range falseOptions {
		assert.False(t, lazyops.And(bools...), "Lazy And failed for case %+v, expected false", bools)
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
		assert.True(t, lazyops.Or(bools...), "Lazy Or failed for case %+v, expected true", bools)
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
		assert.False(t, lazyops.Or(bools...), "Lazy Or failed for case %+v, expected false", bools)
	}
}

func TestXorTrueVarients(t *testing.T) {
	trueOptions := [][]lazyops.Predicate{
		{alwaysTrue},
		{alwaysTrue, alwaysFalse},
		{alwaysFalse, alwaysTrue},
		{alwaysFalse, alwaysTrue, alwaysFalse},
		{alwaysFalse, alwaysFalse, alwaysTrue},
		{alwaysTrue, alwaysTrue, alwaysTrue, alwaysFalse},
	}

	for _, bools := range trueOptions {
		assert.True(t, lazyops.Xor(bools...), "Lazy Xor failed for case %+v, expected true", bools)
	}
}

func TestXorFalseVarients(t *testing.T) {
	falseOptions := [][]lazyops.Predicate{
		{},
		{alwaysFalse},
		{alwaysFalse, alwaysFalse},
		{alwaysTrue, alwaysTrue},
		{alwaysTrue, alwaysTrue, alwaysFalse},
		{alwaysFalse, alwaysTrue, alwaysTrue},
		{alwaysTrue, alwaysTrue, alwaysFalse, alwaysFalse},
	}

	for _, bools := range falseOptions {
		assert.False(t, lazyops.Xor(bools...), "Lazy Xor failed for case %+v, expected false", bools)
	}
}
