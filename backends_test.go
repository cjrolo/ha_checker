package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBaseBackend(t *testing.T) {
	bb, err := NewBaseBackend(1235234, 8000)
	assert.NoError(t, err)
	assert.Equal(t, bb, &BaseBackend{Ip: 1235234, Port: 8000})
}

func TestCreateBaseBackendNegativeValues(t *testing.T) {
	bb, err := NewBaseBackend(-1235234, -8000)
	assert.EqualError(t, err, "IP and Port can't have negative values!")
	assert.Nil(t, bb)
}
