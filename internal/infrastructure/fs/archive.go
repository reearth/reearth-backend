package fs

import (
	"os"
	"path"
	"strings"

	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type Archive struct {
	p       string
	files   []string
	counter int
	name    string
	size    int64
}

func NewArchive(p string) (*Archive, error) {
	bp := strings.TrimSuffix(p, "/")
	files, size, err := dirwalk(bp, "", 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, rerror.ErrNotFound
		}
		return nil, rerror.ErrInternalBy(err)
	}

	return &Archive{
		p:       bp,
		files:   files,
		counter: 0,
		name:    path.Base(p),
		size:    size,
	}, nil
}

func (a *Archive) Next() (f *file.File, derr error) {
	if len(a.files) <= a.counter {
		return nil, nil
	}
	next := a.files[a.counter]
	a.counter++
	fi, err := os.Open(path.Join(a.p, next))
	if err != nil {
		derr = rerror.ErrInternalBy(err)
		return
	}
	stat, err := fi.Stat()
	if err != nil {
		derr = rerror.ErrInternalBy(err)
		return
	}

	f = &file.File{
		Content: fi,
		Path:    strings.TrimPrefix(next, a.p+"/"),
		Size:    stat.Size(),
	}
	return
}

func (a *Archive) Name() string {
	return a.name
}

func (a *Archive) Size() int64 {
	return a.size
}

func dirwalk(dir string, base string, size int64) ([]string, int64, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, 0, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			fname := file.Name()
			dfiles, dsize, err := dirwalk(path.Join(dir, fname), path.Join(base, fname), size)
			if err != nil {
				return []string{}, 0, err
			}
			paths = append(paths, dfiles...)
			size += dsize
			continue
		}
		paths = append(paths, path.Join(base, file.Name()))
		fileInfo, err := file.Info()
		if err != nil {
			return []string{}, 0, err
		}
		size += fileInfo.Size()
	}

	return paths, size, nil
}
