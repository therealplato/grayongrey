package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/therealplato/grayongrey"
)

func main() {
	var n uint
	var input = &bytes.Buffer{}
	flag.UintVar(&n, "n", 2, "number of aliens")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

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
			log.Fatal("Pipe the starting city topography in or 'gg filename'")
		} else if flag.NArg() > 1 {
			log.Println("ignoring arguments after " + flag.Arg(0))
		}
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalf("issue opening file %q: %v", os.Args[1], err)
		}
		_, err = io.Copy(input, f)
		if err != nil {
			log.Fatalf("issue reading input file: %v", err)
		}
	}
	world, err := grayongrey.New(input, n)
	if err != nil {
		log.Fatalf("issue instantiating world: %v", err)
	}
	for !world.GameOver() {
		world.Iterate()
	}
	fmt.Println("\nFinal Topography:")
	world.Log()
}
