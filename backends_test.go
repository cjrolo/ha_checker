package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setUpBackend() *BaseBackend {
	bb, _ := NewBaseBackend("127.0.0.1", 8000)
	return bb
}

func TestCreateBaseBackend(t *testing.T) {
	bb, err := NewBaseBackend("127.0.0.1", 8000)
	assert.NoError(t, err)
	assert.Equal(t, bb, setUpBackend())
}

func TestCreateBaseBackendWrongValues(t *testing.T) {
	bb, err := NewBaseBackend("555.555.332.221", -8000)
	assert.EqualError(t, err, "IP and Port can't have negative values!")
	assert.Nil(t, bb)
}

func TestIsBackendDown(t *testing.T) {
	bb := setUpBackend()
	bb.Down = true
	down := bb.IsDown()
	assert.Equal(t, down, true)
}

func TestIsBackendUp(t *testing.T) {
	bb := setUpBackend()
	down := bb.IsDown()
	assert.Equal(t, down, false)
}

func TestVerifyBackendifDown(t *testing.T) {
	bb := setUpBackend()
	bb.Verify()
	bb.Down = true
	down := bb.IsDown()
	assert.Equal(t, down, true)
}

func TestVerifyBackendifUp(t *testing.T) {
	bb := setUpBackend()
	bb.Verify()
	bb.Down = false
	down := bb.IsDown()
	assert.Equal(t, down, false)
}
