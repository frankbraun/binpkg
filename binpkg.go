// binpkg is a binary installer for Codechain secure packages.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/binpkg/command"
)

func usage() {
	cmd := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage: %s download\n", cmd)
	fmt.Fprintf(os.Stderr, "       %s install [-p prefix]\n", cmd)
	fmt.Fprintf(os.Stderr, "       %s uninstall [-p prefix]\n", cmd)
	fmt.Fprintf(os.Stderr, "       %s generate bindir\n", cmd)
	os.Exit(2)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	argv0 := os.Args[0] + " " + os.Args[1]
	args := os.Args[2:]
	var err error
	switch os.Args[1] {
	case "download":
		err = command.Download(argv0, args...)
	case "install":
		err = command.Install(argv0, args...)
	case "uninstall":
		err = command.Uninstall(argv0, args...)
	case "generate":
		err = command.Generate(argv0, args...)
	default:
		usage()
	}
	if err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
			os.Exit(1)
		}
		os.Exit(2)
	}
}
