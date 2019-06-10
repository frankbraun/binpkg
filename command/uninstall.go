package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/binpkg/pkg"
)

// Unnstall implements the 'uninstall' command.
func Uninstall(argv0 string, args ...string) error {
	fs := flag.NewFlagSet(argv0, flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", argv0)
		fmt.Fprintf(os.Stderr, "Uninstall installed binary package for current platform.\n")
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
	return pkg.Uninstall(*prefix)
}
