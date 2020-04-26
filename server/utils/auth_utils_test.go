package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandomToken(t *testing.T) {
	// We want to make sure we don't get the same token back twice.
	// I guess this could happen, but it's not likely
	t1 := RandomToken()
	t2 := RandomToken()
	assert.NotEqual(t, t1, t2, "2 randomly generated tokens shouldn't be equal")
}
