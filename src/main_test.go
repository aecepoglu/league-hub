package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFooGood(t *testing.T) {
	assert.Equal(t, foo(), 2)
}

func TestFooFail(t *testing.T) {
	assert.Equal(t, foo(), 2)
}
