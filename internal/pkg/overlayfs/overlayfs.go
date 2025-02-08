package overlayfs

import (
	"errors"
	"io/fs"
	"os"
)

// ensure it implements ReadDirFs
var _ fs.ReadDirFS = (*overlayFS)(nil)

// overlayFS combines filesystems. It represents a union of the filesystems with preference being given to the files from the first one in the slice
type overlayFS struct {
	filesystems []fs.ReadDirFS
}

// Open opens the named file.
func (ofs overlayFS) Open(name string) (fs.File, error) {
	for _, fs := range ofs.filesystems {
		file, err := fs.Open(name)
		if err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

func (ofs overlayFS) ReadDir(name string) ([]fs.DirEntry, error) {
	directories := make(map[string]fs.DirEntry)
	var neCnt int
	// reads from all filesystems
	// hard stops on any error other then ErrNotExist
	for _, filesystem := range ofs.filesystems {
		dirEntries, err := fs.ReadDir(filesystem, name)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				neCnt++
				continue
			}
			return nil, err
		}
		for _, e := range dirEntries {
			// if an entry already exists do not override it
			if _, ok := directories[e.Name()]; !ok {
				directories[e.Name()] = e
			}
		}
	}
	if len(ofs.filesystems) == neCnt {
		return nil, fs.ErrNotExist
	}

	// transform back to slice
	dirs := make([]fs.DirEntry, 0, len(directories))
	for _, e := range directories {
		dirs = append(dirs, e)
	}
	return dirs, nil
}

func reverse(fs []fs.ReadDirFS) {
	for i, j := 0, len(fs)-1; i < j; i, j = i+1, j-1 {
		fs[i], fs[j] = fs[j], fs[i]
	}
}

// New creates a overlayfilesystem the last added one will have precedence
func New(f ...fs.ReadDirFS) *overlayFS {
	reverse(f)
	o := overlayFS{
		filesystems: f,
	}
	return &o
}
