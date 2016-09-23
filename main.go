// CRC is a simple command-line utility to compute CRC values for one or more files.
package main

import (
	"flag"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"os"
)

func main() {
	mode := flag.String("mode", "crc64-ecma",
		"CRC method to use.  Valid values are 'crc32' (IEEE), 'crc64-iso', and 'crc64-ecma'")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-mode=<MODE>] file [file ...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	hasher, err := NewHasher(*mode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Specify one or more filenames to checksum.\n")
		os.Exit(2)
	}

	exitCode := 0
	for _, filename := range flag.Args() {
		hasher, _ = NewHasher(*mode)
		f, err := os.Open(filename)
		if err != nil {
			exitCode = 3
			fmt.Fprintf(os.Stderr, "Cannot open %q: %v\n", filename, err)
			continue
		}
		_, err = io.Copy(hasher, f)
		f.Close()
		if err != nil {
			exitCode = 3
			fmt.Fprintf(os.Stderr, "Error reading %q: %v\n", filename, err)
			continue
		}

		if flag.NArg() == 1 {
			fmt.Printf("%0*x\n", hasher.Size(), hasher.Sum(nil))
		} else {
			fmt.Printf("%0*x\t%s\n", hasher.Size(), hasher.Sum(nil), filename)
		}
	}
	os.Exit(exitCode)
}

func NewHasher(mode string) (hash.Hash, error) {
	switch mode {
	case "crc32", "crc32-ieee":
		return crc32.NewIEEE(), nil
	case "crc64-iso":
		return crc64.New(crc64.MakeTable(crc64.ISO)), nil
	case "crc64-ecma":
		return crc64.New(crc64.MakeTable(crc64.ECMA)), nil
	default:
		return nil, fmt.Errorf("ERROR: Invalid mode %q", mode)
	}
}
