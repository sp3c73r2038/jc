package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ugorji/go/codec"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	var h codec.Handle = new(codec.JsonHandle)
	var dec *codec.Decoder = codec.NewDecoderBytes(b, h)
	err = dec.Decode(&v)
	if err != nil {
		var read int = dec.NumBytesRead()
		var char int = 1
		var line int = 1 // human readable line
		var newline = true
		for i := 0; i < read; i++ {
			if b[i] == '\n' {
				line++
				newline = true
			} else {
				if newline {
					char = 0
					newline = false
				} else {
					char++
				}
			}
		}
		if len(b) > 0 && read-1 >= 0 && b[read-1] == '\n' {
			// in case last char is a newline
			line--
		}
		// log.Println(b[read-1])
		log.Printf("error at line %d, char %d", line, char)
		// log.Print(reflect.TypeOf(err))
		log.Fatal(err)
	}
}
