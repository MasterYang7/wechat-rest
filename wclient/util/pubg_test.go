package util

import (
	"fmt"
	"testing"
)

func TestGetkey(t *testing.T) {
	for i := 0; i < 300; i++ {
		index := getKey()
		fmt.Println(index)
	}
}
