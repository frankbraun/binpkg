// Package def defines default values used in binpkg.
package def

import (
	"path/filepath"
)

// ConfigFile is the default config file.
const ConfigFile = "config.binpkg"

// ArchiveDir is the default archive directory.
var ArchiveDir string

func init() {
	ArchiveDir = filepath.Join(".codechain", "binpkg", "archives")
}
