package walkdirs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type AllInput struct {
	SL      bool
	D       bool
	F       bool
	EXT     string
	Dirname string
}

func (a *AllInput) ParseAllInput(sl, d, f bool, ext string, args []string) error {
	a.SL = sl
	a.D = d
	a.F = f

	if ext != "" {
		if a.F {
			a.EXT = ext
		} else {
			return errors.New("-f flag is not specified")
		}
	}

	if !a.SL && !a.D && !a.F {
		a.SL = true
		a.D = true
		a.F = true
	}

	if len(args) == 0 {
		return errors.New("path is not specified")
	}

	path := args[len(args)-1]

	a.Dirname = path

	return nil
}

func PrintDir(a AllInput) error {
	fsys := os.DirFS(a.Dirname)
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if p == "." {
			return nil
		}

		file := filepath.Join(a.Dirname, p)
		_, errs := os.Open(file)
		if errs != nil {
			error_string := fmt.Sprintf("%v", errs)
			contains := strings.Contains(error_string, "permission denied")
			if contains {
				return nil
			}
		}

		if a.D {
			if d.IsDir() {
				fmt.Println(file)
			}
		}
		if a.F {
			if d.Type().IsRegular() {
				if a.EXT != "" {
					if path.Ext(p) == "."+a.EXT {
						fmt.Println(file)
					}
				} else {
					fmt.Println(file)
				}
			}
		}
		if a.SL {
			if d.Type() == fs.ModeSymlink {
				path, errs := filepath.EvalSymlinks(file)
				file += " ->"
				if errs != nil {
					fmt.Println(file, "[broken]")
				} else {
					fmt.Println(file, path)
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
