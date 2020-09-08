package markup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	templateDir = fmt.Sprintf("%s/.config/markup", os.Getenv("HOME"))
)

type file struct {
	name string
	ext  string
}

func Markup(input *file, template *file) error {
	err := checkBinary(`pandoc`)
	if err != nil {
		return err
	}

	out, err := exec.Command(
		`pandoc`,
		fmt.Sprintf("%s%s", input.name, input.ext),
		`-s`,
		fmt.Sprintf("--template=%s%s", template.name, template.ext),
		fmt.Sprintf("--output=%s.pdf", input.name),
		"--pdf-engine=xelatex",
		"--listings",
	).CombinedOutput()

	if len(out) > 1 {
		fmt.Println(string(out))
	}
	return err
}


func Xerox(input *file) error {
	err := checkBinary(`convert`)
	if err != nil {
		return err
	}

	out, err := exec.Command(
		`convert`,
		`-density`,`150`,
		fmt.Sprintf("%s.pdf", input.name),
		`-rotate`, `0.5`,
		`-attenuate`, `0.7`,
		`+noise`, `Multiplicative`,
		`-colorspace`, `Gray`,
		fmt.Sprintf("%s.pdf", input.name),
	).CombinedOutput()

	if len(out) > 1 {
		fmt.Println(string(out))
	}
	return err
}


func checkBinary(name string) error {
	_, err := exec.LookPath(name)
	return err
}

func FindAllTemplates() error {
	fmt.Printf("Template Directory: %s\n", templateDir)
	return filepath.Walk(templateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			if filepath.Ext(path) == `.tex` {
				fmt.Println(strings.TrimPrefix(path, filepath.Dir(path)))
			}
			return nil
	})
}


func FindTemplate(name string) (*file, error) {
	return FindFile(fmt.Sprintf("%s/%s.tex", templateDir, name))
}


func FindFile(input string) (*file, error) {
	if ! checkFile(input) {
		return nil, fmt.Errorf("file does not exist: %s", input)
	}

	ext := filepath.Ext(input)

	return &file{
		name: strings.TrimSuffix(input, ext),
		ext: ext,
	}, nil
}


func checkFile(path string) bool {
  info, err := os.Stat(path)
  if os.IsNotExist(err) {
      return false
  }
  return !info.IsDir()
}
