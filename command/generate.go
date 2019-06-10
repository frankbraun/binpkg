package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/binpkg/pkg"
)

// Generate implements the 'generate' command.
func Generate(argv0 string, args ...string) error {
	fs := flag.NewFlagSet(argv0, flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s bindir\n", argv0)
		fmt.Fprintf(os.Stderr, "Generate binary package for bindir directory.\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		fs.Usage()
		return flag.ErrHelp
	}
	return pkg.Generate(fs.Arg(0))
}
