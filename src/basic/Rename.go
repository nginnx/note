package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	getFilelist("/Users/songyuqiang/Desktop/_splash", "phone_gap.png")
}

func getFilelist(path string, targetName string) {

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".png") {
			outputPath := fmt.Sprint(filepath.Dir(path) + "/" + targetName)
			os.Rename(path, outputPath)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
