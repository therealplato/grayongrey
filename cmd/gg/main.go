package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var n uint
	var input = &bytes.Buffer{}
	flag.UintVar(&n, "n", 2, "number of aliens")
	flag.Parse()
	_ = n

	// handle stdin pipes, via https://stackoverflow.com/a/43947435/1380669
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("issue inspecting stdin: %v", err)
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		_, err := io.Copy(input, os.Stdin)
		if err != nil {
			log.Fatalf("issue reading stdin: %v", err)
		}
	} else {
		// assumption: non-piped input is specified by filename
		if len(os.Args) < 2 {
			log.Fatal("Pipe the starting topography in or 'gg filename'")
		} else if len(os.Args) > 2 {
			log.Println("ignoring arguments after " + os.Args[1])
		}
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("issue opening file %q: %v", os.Args[1], err)
		}
		_, err = io.Copy(input, f)
		if err != nil {
			log.Fatalf("issue reading input file: %v", err)
		}
	}
	log.Println("input file:")
	log.Println(input.String())
}
