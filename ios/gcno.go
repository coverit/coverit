package main

import (
	// "archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
