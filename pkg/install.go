package pkg

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/frankbraun/codechain/tree"
	"github.com/frankbraun/codechain/util/file"
	"github.com/frankbraun/codechain/util/hex"
)

// Install binary package, see specification for details.
func Install(prefix string) error {
	// 1. If `$prefix` is not set, set `$prefix=/usr/local`.
	if prefix == "" {
		// `$prefix` should have been set on command line
		return errors.New("prefix not set")
	}

	// make sure `$prefix` exists and is a directory
	if err := dirExists(prefix); err != nil {
		return err
	}

	// 2. Determine the `$GOOS` and `$GOARCH` we are currently running on.
	platform := runtime.GOOS + "_" + runtime.GOARCH

	// 3. Read the corresponding `$GOOS_$GOARCH.binpkg` file and hash it to
	//    `treehash`. Abort, if `$GOOS_$GOARCH.binpkg` does not exist.
	binpkg := platform + ".binpkg"
	b, err := ioutil.ReadFile(binpkg)
	if err != nil {
		return err
	}
	h := sha256.Sum256(b)
	treehash := hex.Encode(h[:])

	// 4. Calculate tree hash of `.codechain/binpkg/$GOOS_GOARCH`. Abort, if
	//    the calculated tree hash does not equal `treehash`.
	root := filepath.Join(".codechain", "binpkg", platform)
	th, err := tree.Hash(root, nil)
	if err != nil {
		return err
	}
	if hex.Encode(th[:]) != treehash {
		return fmt.Errorf("treehashes do not match for '%s' and '%s'",
			binpkg, root)
	}

	// 5. Install: `cp -rf .codechain/binpkg/$GOOS_$GOARCH/* $prefix`
	entries, err := tree.List(root, nil)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		src := filepath.Join(root, entry.Filename)
		dst := filepath.Join(prefix, entry.Filename)
		if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
			return err
		}
		os.Remove(dst) // remove target first
		if err := file.Copy(src, dst); err != nil {
			return err
		}
	}
	return nil
}
