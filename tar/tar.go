// Package tar reads and writes gzipped tar files.
package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// Create writes the files contained in root as a gzipped tar archive to w.
// The file strings in files can contain directory paths (starting from root).
func Create(w io.Writer, root string, files []string) error {
	zw := gzip.NewWriter(w)
	tw := tar.NewWriter(zw)

	for _, file := range files {
		f, err := os.Open(filepath.Join(root, file))
		if err != nil {
			return err
		}
		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return err
		}
		mode := fi.Mode() & os.ModePerm // only keep standard UNIX permission bits
		hdr := &tar.Header{
			Name:    file,
			Mode:    int64(mode),
			Size:    fi.Size(),
			ModTime: fi.ModTime(),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			f.Close()
			return err
		}
		if _, err := io.Copy(tw, f); err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}
	return zw.Close()
}

// Extract the files contained in the gzipped archive r to the directory
// hierarchy rooted at root.
func Extract(root string, r io.Reader) error {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	tr := tar.NewReader(zr)

	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break // end of archive
			}
			return err
		}
		outdir := filepath.Join(root, filepath.Dir(hdr.Name))
		if err := os.MkdirAll(outdir, 0755); err != nil {
			return err
		}
		fn := filepath.Join(root, hdr.Name)
		f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
		if err != nil {
			return err
		}
		if _, err := io.Copy(f, tr); err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}

	return zr.Close()
}

// CreateArchive writes the files contained in root as a gzipped tar archive
// to filename.
// The file strings in files can contain directory paths (starting from root).
func CreateArchive(filename, root string, files []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return Create(f, root, files)
}

// ExtractArchive extracts the files contained in the gzipped archive filename
// to the directory hierarchy rooted at root.
func ExtractArchive(root, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return Extract(root, f)
}
