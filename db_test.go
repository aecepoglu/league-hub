package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDbConnect(t *testing.T) {
	db, err := connectDb("test.db")
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

