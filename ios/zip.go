package main

import (
	"github.com/jhoonb/archivex"
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
