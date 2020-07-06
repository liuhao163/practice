package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "abc"
	assert.Equal(t, ReverseString(s), "cba")
	s = "abcd"
	assert.Equal(t, ReverseString(s), "dcba")
}

func TestCreateTriangle(t *testing.T) {
	ret := CreateTriangle(5)
	Print(ret)
}
