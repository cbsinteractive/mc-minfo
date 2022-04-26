package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/as/minfo"
)

func main() {
	flag.Parse()
	for _, a := range flag.Args() {
		var (
			file minfo.File
			err  error
		)
		if a == "-" {
			file, err = minfo.Read(os.Stdin)
		} else {
			file, err = minfo.ReadURL(context.Background(), a)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file %q: %v", a, err)
		} else {
			fmt.Println(file.String())
		}
	}

}
