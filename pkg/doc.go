/*
Package pkg implements binary packages.

Download specification

Downloading binary packages works as follows:

  1. Parse the `config.binpkg` file.
  2. Determine the `$GOOS` and `$GOARCH` we are currently running on.
  3. Read the corresponding `$GOOS_$GOARCH.binpkg` file and hash it to
     `treehash`. Abort, if `$GOOS_$GOARCH.binpkg` does not exist.
  4. Pick a random URL from the `config.binpkg` file and try to download
     `URL/$GOOS_$GOARCH/treehash.tar.gz` to `.codechain/binpkg/archives`.
  5. If download (or verification, see below) failed, try next URL.
     Abort, if all downloads fail permanently.
  6. Remove directory `.codechain/binpkg/$GOOS_$GOARCH`.
  7. Create directory `.codechain/binpkg/$GOOS_$GOARCH`.
  8. Extract `treehash.tar.gz` to `.codechain/binpkg/$GOOS_$GOARCH` and
     calculate tree hash. If the calculated tree hash does not equal
     `treehash` goto 5.
  9. The binary package to install for the current `$GOOS` and `$GOARCH`
     is now contained in the directory hierarchy under
     `.codechain/binpkg/$GOOS_$GOARCH`.

Install specification

Installing binary packages works as follows:

  1. If `$prefix` is not set, set `$prefix=/usr/local`.
  2. Determine the `$GOOS` and `$GOARCH` we are currently running on.
  3. Read the corresponding `$GOOS_$GOARCH.binpkg` file and hash it to
     `treehash`. Abort, if `$GOOS_$GOARCH.binpkg` does not exist.
  4. Calculate tree hash of `.codechain/binpkg/$GOOS_GOARCH`. Abort, if
     the calculated tree hash does not equal `treehash`.
  5. Install: `cp -rf .codechain/binpkg/$GOOS_$GOARCH/* $prefix`

Uninstall specification

Uninstalling binary packages works as follows:

  1. If `$prefix` is not set, set `$prefix=/usr/local`.
  2. Determine the `$GOOS` and `$GOARCH` we are currently running on.
  3. Read corresponding `$GOOS_$GOARCH.binpkg` file and hash it to
     `treehash`. Abort, if `$GOOS_$GOARCH.binpkg` does not exist.
  4. Calculate tree hash of `.codechain/binpkg/$GOOS_GOARCH`. Abort, if
     the calculated tree hash does not equal `treehash`.
  5. Delete all files contained in `.codechain/binpkg/$GOOS_$GOARCH` from
     the directory hierarchy rooted at `$prefix`.

Generate specification

Generating binary packages works as follows:

  1. Parse the `config.binpkg` file.
  2. Ensure that the base name of `$bindir` is a valid `$GOOS_$GOARCH`
     combination.
  3. Ensure that the file `$GOOS_$GOARCH.binpkg` does not exist.
  4. Calculate the tree list of `$bindir`, write it to
     `$GOOS_$GOARCH.binpkg`, and hash it to `treehash`.
  5. Write the directory hierarchy below `$bindir` as as archive to
     `.codechain/binpkg/archives/treehash.tar.gz`.
  6. Display all paths `URL/$GOOS_$GOARCH/treehash.tar.gz` where
     `.codechain/binpkg/archives/treehash.tar.gz` needs to be uploaded to.
*/
package pkg
