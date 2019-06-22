binpkg — a binary installer for Codechain secure packages
---------------------------------------------------------

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/frankbraun/binpkg)
[![Build
Status](https://img.shields.io/travis/frankbraun/binpkg.svg?style=flat-square)](https://travis-ci.org/frankbraun/binpkg)
[![Go Report
Card](https://goreportcard.com/badge/github.com/frankbraun/binpkg?style=flat-square)](https://goreportcard.com/report/github.com/frankbraun/binpkg)

`binpkg` allows to install binaries as secure packages with
[Codechain](https://github.com/frankbraun/codechain).

Adding binaries directly to a Codechain secure package would increase
its size too much and binaries cannot be reviewed in a meaningful way
anyway.

Under normal circumstances building from source is preferable, but in
the rare cases where you want to distributed binaries in a secure and
multiparty reviewed way, you can use `binpkg` as follows:

1.  Add the `.secpkg` file of this `binpkg` repository to the `.secdep`
    directory of your package.
2.  Add a Makefile that calls `binpkg download` for `make`,
    `binpkg install` for `make install`, and `binpkg uninstall` for
    `make uninstall`. Also make sure to pass through the `$prefix`
    variable from `make` to the `-p` option for the `binpkg install` and
    `binpkg uninstall` commands (see example `Makefile`).
3.  Add the configuration file `config.binpkg` as described below.
4.  Add a distribution file `$GOOS_$GOARCH.binpkg` for every platform
    you want to support (with the help of `binpkg generate`, see below).
5.  Upload the distribution archives to the configured web server paths
    (as displayed by `binpkg generate`).
6.  Add all `*.binpkg` files to Codechain, review them, and publish the
    secure package.

This ensures multiparty signatures of the _hashes_ of all installed
binaries. Without [Reproducible
Builds](https://reproducible-builds.org/) this just records in an
unmodifiable way which binaries have been pushed by the developers. With
Reproducible Builds these binaries could be audited with the
corresponding source code, but the specifics of such a procedure are
outside of the scope of `binpkg`.

Using Codechain secure dependencies allows to extend Codechain with
binary packages without blowing up Codechain itself unnecessarily.

### Commands

#### `binpkg download`

Download binary package for current platform, see
[specification](https://godoc.org/github.com/frankbraun/binpkg/pkg#hdr-Download_specification)
for details.

#### `binpkg install`

Install downloaded binary package for current platform, see
[specification](https://godoc.org/github.com/frankbraun/binpkg/pkg#hdr-Install_specification)
for details.

#### `binpkg uninstall`

Uninstall installed binary package for current platform, see
[specification](https://godoc.org/github.com/frankbraun/binpkg/pkg#hdr-Uninstall_specification)
for details.

#### `binpkg generate $bindir`

Generate binary package for `$bindir` directory, see
[specification](https://godoc.org/github.com/frankbraun/binpkg/pkg#hdr-Generate_specification)
for details.

### Files and directories

#### `config.binpkg`

A binary package configuration file (`config.binpkg`) contains a JSON
object with the following keys:

    {
      "URLs": [
        "list of binary package download URLs"
      ]
    }

Example `config.binpkg` file:

    {
      "URLs": [
        "http://example.com/binpkg/testpackage",
        "http://example.net/binpkg/testpackage",
        "http://example.org/binpkg/testpackage"
      ]
    }

#### `$GOOS_$GOARCH.binpkg` files

A `$GOOS_$GOARCH.binpkg` file (e.g., `linux_amd64.binpkg`) contains a
[tree list](https://godoc.org/github.com/frankbraun/codechain/tree) of
all files in their relative directories and their hashes that are
installed by `binpkg install` for this platform.

Example `linux_amd64.binpkg` file:

    x 1c9d23c245ef06a87f178c5d82221b702084540fe072b329c6a992d6036e6649 bin/testbin
    x e39447e1a9d87131b62ee4f5fcfe0bd11aa5a8c545b706424d38ca7a23d24f9c bin/testbin2

#### `.codechain/binpkg` directory tree

`binpkg` uses the directory tree under `.codechain/binpkg` for temporary
data. By being under the `.codechain` hierarchy the temporary data is
excluded from Codechain's hash chain.

`.codechain/binpkg/archives` is used for storing package archive files.

`.codechain/binpkg/$GOOS_$GOARCH` directories are used to extract
package archives for the corresponding platform in order to check the
contents and prepare for the installation.

#### path on web server

A common path on a web server would like this:

    URL/binpkg/package_name/$GOOS_$GOARCH/treehash.tar.gz

where:

-   `binpkg` is optional.
-   `package_name` is the name of the package (optional).
-   `$GOOS_$GOARCH` is the platform string (mandatory, but not part of
    URL in `config.binpkg`).
-   `treehash` is the [tree
    hash](https://godoc.org/github.com/frankbraun/codechain/tree) in hex
    notation (lowercase) of all installed files for this platform and
    `treehash.tar.gz` contains the corresponding directory tree as a
    `.tar.gz` archive.

### ToDo

-   Make `hello-bin` example project.
-   Add example `Makefile` to README.md.
-   Remove `github.com/mholt/archiver` dependency, can be written with
    stdlib.
-   https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
