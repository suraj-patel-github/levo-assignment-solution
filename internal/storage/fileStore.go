package storage

import (
    "io"
    "os"
    "path/filepath"
)

type FileStore struct {
    BasePath string
}

func NewFileStore(path string) *FileStore {
    os.MkdirAll(path, 0755)
    return &FileStore{BasePath: path}
}

func (f *FileStore) Save(app, service, filename string, r io.Reader) (string, error) {
    dir := filepath.Join(f.BasePath, app)
    if service != "" {
        dir = filepath.Join(dir, service)
    }
    if err := os.MkdirAll(dir, 0755); err != nil {
        return "", err
    }
    fullPath := filepath.Join(dir, filename)
    out, err := os.Create(fullPath)
    if err != nil {
        return "", err
    }
    defer out.Close()
    _, err = io.Copy(out, r)
    return fullPath, err
}
