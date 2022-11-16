package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	res := fizzBuzz(3)
	assert.Equal(t, "Fizz", res)

	res = fizzBuzz(15)
	assert.Equal(t, "FizzBuzz", res)

	res = fizzBuzz(2)
	assert.Equal(t, "2", res)
}
