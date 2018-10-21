package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFooGood(t *testing.T) {
	assert.Equal(t, foo(), 2)
}

func TestFooFail(t *testing.T) {
	assert.Equal(t, foo(), 2)
}
