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

func TestAnd(t *testing.T) {
	assert.True(t, ops.And(true, true))
}
  
