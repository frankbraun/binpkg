package pkg

import (
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/frankbraun/binpkg/config"
	"github.com/frankbraun/binpkg/internal/def"
	"github.com/frankbraun/codechain/util/file"
	"github.com/frankbraun/codechain/util/hex"
)

// Download binary package, see specification for details.
func Download() error {
	// 1. Parse the `config.binpkg` file.
	cfg, err := config.Load(def.ConfigFile)
	if err != nil {
		return err
	}

	// 2. Determine the `$GOOS` and `$GOARCH` we are currently running on.
	platform := runtime.GOOS + "_" + runtime.GOARCH

	// 3. Read the corresponding `$GOOS_$GOARCH.binpkg` file and hash it to
	//    `treehash`. Abort, if `$GOOS_$GOARCH.binpkg` does not exist.
	b, err := ioutil.ReadFile(platform + ".binpkg")
	if err != nil {
		return err
	}
	h := sha256.Sum256(b)
	treehash := hex.Encode(h[:])

	// 4. Pick a random URL from the `config.binpkg` file and try to download
	//    `URL/$GOOS_$GOARCH/treehash.tar.gz` to `.codechain/binpkg/archives`.
	url, err := cfg.RandomURL()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(def.ArchiveDir, 0755); err != nil {
		return err
	}
	archive := filepath.Join(def.ArchiveDir, treehash+".tar.gz")
	fullURL := url + "/" + platform + "/" + treehash + ".tar.gz"
	if err := file.Download(archive, fullURL); err != nil {
		return err
	}

	// 5. If download (or verification, see below) failed, try next URL.
	//    Abort, if all downloads fail permanently.

	// 6. Remove directory `.codechain/binpkg/$GOOS_$GOARCH`.

	// 7. Create directory `.codechain/binpkg/$GOOS_$GOARCH`.

	// 8. Extract `treehash.tar.gz` to `.codechain/binpkg/$GOOS_$GOARCH` and
	//    calculate tree hash. If the calculated tree hash does not equal
	//    `treehash` goto 5.

	// 9. The binary package to install for the current `$GOOS` and `$GOARCH`
	//    is now contained in the directory hierarchy under
	//    `.codechain/binpkg/$GOOS_$GOARCH`.

	return nil
}
