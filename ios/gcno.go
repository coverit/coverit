package main

import (
	"fmt"
	"github.com/jhoonb/archivex"
	"os"
	"path/filepath"
	"strings"
)

func ZipGcnos(filename string, files []string) error {
	zip := new(archivex.ZipFile)
	zip.Create(filename)
	for _, f := range files {
		zip.AddFile(f)
	}
	zip.Close()

	return nil
}

func CollectGcno(root string) ([]string, error) {

	gcnos := make([]string, 0, 10)

	visit := func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() && strings.HasSuffix(path, ".gcno") {
			gcnos = append(gcnos, path)
			fmt.Println(path)
		}
		return nil
	}

	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Errorf("Walk faild")
		return nil, err
	}

	return gcnos, nil
}
