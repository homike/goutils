package goutils

import "strings"

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n+1:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func GetBetweenStr_FirstLast(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n+1:])
	m := strings.LastIndex(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

// delete last 'delStr' in str
func DelLastStr(str, delStr string) string {
	arrStr := []byte(str)
	index := 0
	for i := len(arrStr) - 1; i >= 0; i-- {
		if string(arrStr[i]) == delStr {
			index = i
			break
		}
	}
	return string(append(arrStr[:index], arrStr[index+1:]...))
}
