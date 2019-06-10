package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/binpkg/pkg"
)

// Install implements the 'install' command.
func Install(argv0 string, args ...string) error {
	fs := flag.NewFlagSet(argv0, flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", argv0)
		fmt.Fprintf(os.Stderr, "Install downloaded binary package for current platform.\n")
		fs.PrintDefaults()
	}
	prefix := fs.String("p", defaultPrefix, "Set prefix")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 0 {
		fs.Usage()
		return flag.ErrHelp
	}
	return pkg.Install(*prefix)
}
