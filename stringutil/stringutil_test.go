package stringutil

import (
	"fmt"
	"testing"
)

func TestReverseStr(t *testing.T) {
	str := []byte("abcde")
	ReverseStr(str)
	fmt.Println(string(str))
}
