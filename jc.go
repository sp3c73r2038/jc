package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// https: //adrianhesketh.com/2017/03/18/getting-line-and-character-positions-from-gos-json-unmarshal-errors/
func reportJSONError(err error, b []byte) (rv error) {
	var line int64
	var char int64
	if err == nil {
		return
	}
	switch err.(type) {
	case *json.SyntaxError:
		jerr := err.(*json.SyntaxError)
		line, char, rv = getLineAndChar(string(b), jerr.Offset)
	case *json.UnmarshalTypeError:
		jerr := err.(*json.UnmarshalTypeError)
		line, char, rv = getLineAndChar(string(b), jerr.Offset)
	}

	if rv != nil {
		return
	}

	fmt.Printf("found error at %d, char %d, %s", line, char, err)
	return
}

func getLineAndChar(s string, offset int64) (
	line int64, char int64, err error) {
	l := len(s)
	if offset > int64(l) || offset < 0 {
		err = fmt.Errorf("invalid offset %d, with content length: %d", offset, l)
		return
	}
	lf := rune(0x0A)
	line = 1
	for i, b := range s {
		if b == lf {
			line++
			char = 0
		}
		char++
		if int64(i) == offset {
			break
		}
	}
	return
}

func main() {

	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Printf("usage: %s <file> \n", os.Args[0])
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(args[0])
	must(err)

	var v interface{}
	err = json.Unmarshal(b, &v)
	reportJSONError(err, b)

}
