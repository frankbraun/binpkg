package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/frankbraun/binpkg/config"
	"github.com/frankbraun/binpkg/internal/def"
	"github.com/frankbraun/binpkg/tar"
	"github.com/frankbraun/codechain/tree"
	"github.com/frankbraun/codechain/util/file"
	"github.com/frankbraun/codechain/util/hex"
)

// Generate binary package, see specification for details.
func Generate(bindir string) error {
	// 1. Parse the `config.binpkg` file.
	cfg, err := config.Load(def.ConfigFile)
	if err != nil {
		return err
	}

	// 2. Ensure that the base name of `$bindir` is a valid `$GOOS_$GOARCH`
	//    combination.
	platform := filepath.Base(bindir)
	if err := platformIsValid(platform); err != nil {
		return err
	}

	// 3. Ensure that the file `$GOOS_$GOARCH.binpkg` does not exist.
	filename := platform + ".binpkg"
	exists, err := file.Exists(filename)
	if err != nil {
		return err
	}
	if exists {
		fmt.Errorf("file '%s' exists already", filename)
	}

	// 4. Calculate the tree list of `$bindir`, write it to
	//    `$GOOS_$GOARCH.binpkg`, and hash it to `treehash`.
	entries, err := tree.List(bindir, nil)
	if err != nil {
		return err
	}
	h := tree.HashList(entries)
	treehash := hex.Encode(h[:])
	if err := ioutil.WriteFile(filename, tree.PrintList(entries), 0644); err != nil {
		return err
	}
	fmt.Printf("'%s' written \n", filename)

	// 5. Write the directory hierarchy below `$bindir` as as archive to
	//    `.codechain/binpkg/archives/treehash.tar.gz`.
	if err := os.MkdirAll(def.ArchiveDir, 0755); err != nil {
		return err
	}
	archive := filepath.Join(def.ArchiveDir, treehash+".tar.gz")
	var files []string
	for _, entry := range entries {
		files = append(files, entry.Filename)
	}
	if err := tar.CreateArchive(archive, bindir, files); err != nil {
		return err
	}
	fmt.Printf("'%s' written \n", archive)

	// 6. Display all paths `URL/$GOOS_$GOARCH/treehash.tar.gz` where
	//    `.codechain/binpkg/archives/treehash.tar.gz` needs to be uploaded to.
	fmt.Printf("upload '%s' to the following URLs:\n", archive)
	fmt.Println()
	for _, url := range cfg.URLs {
		fmt.Println(url + "/" + platform + "/" + treehash + ".tar.gz")
	}

	return nil
}
