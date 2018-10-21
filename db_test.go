package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbConnect(t *testing.T) {
	db, err := connectDb("test.db")
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
