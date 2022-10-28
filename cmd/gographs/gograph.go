package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gographs "github.com/leandrotocalini/gographs/pkg"
)

func main() {

	root := gographs.CreateNode("", "/", false)
	search := false
	onlyActive := false
	searchString := ""
	if len(os.Args) >= 2 {
		for _, val := range os.Args {
			if val == "-a" {
				onlyActive = true
			} else {
				search = true
				searchString = val
			}
		}

	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		active := false
		if !info.IsDir() {
			if search && strings.Contains(path, searchString) {
				active = true
			}
			root.Insert(path, active)
		}

		return nil
	})

	if err != nil {
		return
	}

	fmt.Println(root.ToString(0, onlyActive))
}
