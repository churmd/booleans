package ops_test

import (
	"testing"

	ops "github.com/churmd/booleans/ops"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	assert.False(t, ops.Not(true))
	assert.True(t, ops.Not(false))
}

func TestAndTrueVarients(t *testing.T) {
	trueOptions := [][]bool{{}, {true}, {true, true}, {true, true, true}}

	for _, bools := range trueOptions {
		assert.True(t, ops.And(bools...))
	}
}

func TestAndFalseVarients(t *testing.T) {
	falseOptions := [][]bool{{false}, {true, false}, {false, true}, {true, true, false}}

	for _, bools := range falseOptions {
		assert.False(t, ops.And(bools...))
	}
}

func TestOrTrueVarients(t *testing.T) {
	trueOptions := [][]bool{{true}, {true, false}, {false, true}, {true, true, false}, {false, false, true}}

	for _, bools := range trueOptions {
		assert.True(t, ops.Or(bools...))
	}
}

func TestOrFalseVarients(t *testing.T) {
	falseOptions := [][]bool{{}, {false}, {false, false}, {false, false}}

	for _, bools := range falseOptions {
		assert.False(t, ops.Or(bools...))
	}
}

func TestXorTrueVarients(t *testing.T) {
	trueOptions := [][]bool{
		{true},
		{true, false},
		{false, true},
		{false, true, false},
		{false, false, true},
		{true, true, true, false},
	}

	for _, bools := range trueOptions {
		assert.True(t, ops.Xor(bools...), "Xor failed for case %+v, expected true", bools)
	}
}

func TestXorFalseVarients(t *testing.T) {
	falseOptions := [][]bool{
		{},
		{false},
		{false, false},
		{true, true},
		{true, true, false},
		{false, true, true},
		{true, true, false, false},
	}

	for _, bools := range falseOptions {
		assert.False(t, ops.Xor(bools...), "Xor failed for case %+v, expected false", bools)
	}
}
