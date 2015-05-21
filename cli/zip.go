package main

import (
	"archive/zip"
	"github.com/jhoonb/archivex"
	"io"
	"os"
	"path/filepath"
)

func Zip(filename string, files []string) error {
	zip := new(archivex.ZipFile)
	zip.Create(filename)
	for _, f := range files {
		zip.AddFile(f)
	}
	zip.Close()

	return nil
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
