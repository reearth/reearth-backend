// Package file provides convenient helpers for files and abstractions of files
package file

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"strings"
)

// File abstracts an abstract file
type File struct {
	Content io.ReadCloser
	Path    string
	Size    int64
	// If the content type is not explicitly specified, ContenType will be an empty string.
	ContentType string
}

// Iterator is an iterator of files
type Iterator interface {
	// Next returns the next File. If there is no next File, returns nil file and nil error
	Next() (*File, error)
}

// For debugging
type SimpleIterator struct {
	c     int
	files []File
}

func NewSimpleIterator(files []File) *SimpleIterator {
	files2 := make([]File, len(files))
	copy(files2, files)
	return &SimpleIterator{
		files: files2,
	}
}

func (s *SimpleIterator) Next() (*File, error) {
	if len(s.files) <= s.c {
		return nil, nil
	}
	n := s.files[s.c]
	s.c++
	return &n, nil
}

type PrefixIterator struct {
	a      Iterator
	prefix string
}

func NewPrefixIterator(a Iterator, prefix string) *PrefixIterator {
	return &PrefixIterator{
		a:      a,
		prefix: prefix,
	}
}

func (s *PrefixIterator) Next() (*File, error) {
	for {
		n, err := s.a.Next()
		if err != nil {
			return nil, err
		}
		if n == nil {
			return nil, nil
		}
		if s.prefix == "" {
			return n, nil
		}
		if strings.HasPrefix(n.Path, s.prefix+"/") {
			n2 := *n
			n2.Path = strings.TrimPrefix(n2.Path, s.prefix+"/")
			return &n2, nil
		}
	}
}

type FilteredIterator struct {
	a       Iterator
	skipper func(p string) bool
}

func NewFilteredIterator(a Iterator, skipper func(p string) bool) *FilteredIterator {
	return &FilteredIterator{
		a:       a,
		skipper: skipper,
	}
}

func (s *FilteredIterator) Next() (*File, error) {
	for {
		n, err := s.a.Next()
		if err != nil {
			return nil, err
		}
		if n == nil {
			return nil, nil
		}
		if !s.skipper(n.Path) {
			return n, nil
		}
	}
}

type ZipReader struct {
	zr *zip.Reader
	i  int
}

func NewZipReader(zr *zip.Reader) *ZipReader {
	return &ZipReader{zr: zr}
}

func ZipReaderFrom(r io.Reader, n int64) (*ZipReader, error) {
	b, err := io.ReadAll(io.LimitReader(r, n))
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	return NewZipReader(zr), nil
}

func (r *ZipReader) Next() (*File, error) {
	if r == nil || r.zr == nil {
		return nil, nil
	}

	if len(r.zr.File) <= r.i {
		return nil, nil
	}

	f := r.zr.File[r.i]
	r.i++

	fi := f.FileInfo()
	if fi.IsDir() {
		return r.Next()
	}

	c, err := f.Open()
	if err != nil {
		return nil, err
	}

	return &File{Content: c, Path: f.Name, Size: fi.Size()}, nil
}

type TarReader struct {
	tr *tar.Reader
}

func NewTarReader(tr *tar.Reader) *TarReader {
	return &TarReader{tr: tr}
}

func TarReaderFromTarGz(r io.Reader) (*TarReader, error) {
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &TarReader{tr: tar.NewReader(gzipReader)}, nil
}

func (r *TarReader) Next() (*File, error) {
	if r == nil || r.tr == nil {
		return nil, nil
	}

	h, err := r.tr.Next()
	if errors.Is(err, io.EOF) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	fi := h.FileInfo()
	if fi.IsDir() {
		return r.Next()
	}

	return &File{Content: io.NopCloser(r.tr), Path: h.Name, Size: fi.Size()}, nil
}
