package stringutil

import (
	"fmt"
	"testing"
)

func TestReverseStr(t *testing.T) {
	str := "abc"
	str = ReverseStr(str)
	fmt.Println(str)
}
