package goutils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

type ForeachFunc func(filename string) error

func ForeachDir(pathname string, f ForeachFunc) error {
	rd, outErr := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			if pathname[len(pathname)-1] != '/' {
				pathname = pathname + "/"
			}
			err := ForeachDir(pathname+fi.Name()+"/", f)
			if err != nil {
				outErr = err
			}

		} else {
			name := pathname + fi.Name()
			err := f(name)
			if err != nil {
				outErr = err
			}
		}
	}

	return outErr
}

func ReadLinesFromFile(fn string) ([]string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)

	lines := []string{}
	for {
		byteline, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(byteline)+"\n")
	}
	return lines, nil
}

func WriteLinesToFile(fn string, lines []string) error {
	err := os.Remove(fn)
	if err != nil {
		return err
	}

	newf, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer newf.Close()

	for _, v := range lines {
		newf.Write([]byte(v))
	}

	return nil
}
