package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	minfo "github.com/cbsinteractive/mc-minfo"
)

var (
	onlydur = flag.Bool("dur", false, "only print duration of root container and exit")
)

func main() {
	flag.Parse()
	for _, a := range flag.Args() {
		var (
			file minfo.File
			err  error
		)
		file, err = minfo.ReadURL(context.Background(), a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file %q: %v", a, err)
		} else {
			if *onlydur {
				fmt.Println(file.Header.Duration)
			} else {
				fmt.Println(file.String())
			}
		}
	}

}
