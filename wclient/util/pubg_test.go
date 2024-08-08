package util

import (
	"fmt"
	"testing"
)

func TestGetkey(t *testing.T) {
	for i := 0; i < 300; i++ {
		index := APP_KEY.GetNext()
		fmt.Println(index)
	}
}
func TestGetkey2(t *testing.T) {
	var TimeMap = make(map[string]int)

	for i := 0; i < 300; i++ {
		TimeMap["num"]++
		fmt.Printf("输出:%d", TimeMap["num"])
	}
}
func TestAres(t *testing.T) {
	a := Plugin{}
	str := a.CheckAstro("sss", "白羊")
	fmt.Printf(str)
}
