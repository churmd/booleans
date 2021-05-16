package ops_test

import (
	"testing"

	ops "github.com/churmd/booleans"

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
	trueOptions := [][]bool{ {false}, {true, false}, {false, true}, {true, true, false}}
	
	for _, bools := range trueOptions {
		assert.False(t, ops.And(bools...))
	}
}
