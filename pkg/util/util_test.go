package util

import (
	"fmt"
	"testing"
)

func TestRandNum(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandNum(5)
		fmt.Println(n)
	}
}
