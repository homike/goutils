package goutils

import (
	"fmt"
	"testing"
)

func TestReadDir(t *testing.T) {
	ForeachDir("./cmd", func(fn string) error {
		lines, _ := ReadLinesFromFile(fn)
		for k, v := range lines {
			lines[k] = v + " 20210921 "
		}

		WriteLinesToFile(fn, lines)

		lines, _ = ReadLinesFromFile(fn)
		for _, v := range lines {
			fmt.Println("[filename] ", fn, "[lines] ", v)
		}
		return nil
	})
}
