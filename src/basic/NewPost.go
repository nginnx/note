package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io"
)

func main() {
	walkPath("../quiz/")
}

const copyTo string = "/Users/songyuqiang/projects/blog/source/_posts/"

func walkPath(src string) []string {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path,".go"){
			return nil
		}

		pureName := strings.TrimSuffix(filepath.Base(path), ".go")
		copyToPath := fmt.Sprint(copyTo + pureName+".md")
		copyToFile, e := os.Create(copyToPath)
		copyToFile.WriteString("title: " + pureName + "\n"+"tag: [leetcode]\n"+"---\n" + "```" + "\n")
		srcFile, e1 := os.Open(path)
		if e != nil {
			return e
		}
		if e1 != nil {
			return e1
		}
		fmt.Printf("%v\n", pureName)

		io.Copy(copyToFile, srcFile)
		copyToFile.WriteString("\n```")
		return nil
	})

	if err != nil {
		fmt.Print(err)
	}
	return nil
}
