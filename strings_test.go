package goutils

import (
	"fmt"
	"testing"
)

func TestGetBetweenStr(t *testing.T) {
	str := "[2020]"
	fmt.Println(GetBetweenStr(str, "[", "]"))
}
