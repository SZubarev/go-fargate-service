package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, GetHello(), "Hello from package function")

}
