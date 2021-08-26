package explorer

import (
	"io/fs"
	"path/filepath"
)

func find(root string, extensions []string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		for _, ext := range extensions {
			if filepath.Ext(d.Name()) == ext {
				a = append(a, s)
			}
		}
		return nil
	})
	return a
}