package main

import (
	//"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	templateName string
	inputName    string
	xeroxFlg     bool
	versionFlg   bool

	commit      string
	version     string
	TEMPLATEDIR string
)

type file struct {
	name string
	dir  string
	ext  string
}

func init() {
	flag.BoolVarP(&versionFlg, "version", "v", false, "Print version and exit")
	flag.BoolVarP(&xeroxFlg, "xerox", "x", false, "Xerox the output")
	flag.StringVarP(&templateName, "template", "t", "", "Template name")

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	} else {
		numarg := flag.NArg()
		inputName = flag.Arg(numarg - 1)
	}

	TEMPLATEDIR = fmt.Sprintf("%s/.config/markup", os.Getenv("HOME"))
}


func main() {
	if versionFlg {
		fmt.Printf("Version: %s (%s)\n", version, commit)
		os.Exit(0)
	}

	mdFile, err := findFile(inputName)
	if err != nil {
		log.Fatal(err)
	}

	template, err := findTemplate(templateName)
	if err != nil {
		log.Fatal(err)
	}

	err = markup(mdFile, template)
}


func markup(input *file, template *file) error {
	err := checkBinary(`pandoc`)
	if err != nil {
		return err
	}

	out, err := exec.Command(
		`pandoc`,
		fmt.Sprintf("%s/%s.%s", input.dir, input.name, input.ext),
		fmt.Sprintf("--template=%s/%s.tex", template.dir, template.name),
		fmt.Sprintf("-o %s/%s.pdf", input.dir, input.name),
		"--pdf-engine=xelatex",
		"--listings",
	).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func xerox(input *file) error {
	err := checkBinary(`convert`)
	if err != nil {
		return err
	}

	out, err := exec.Command(
		`convert`,
		`-density 150`,
		fmt.Sprintf("%s/%s.pdf", input.dir, input.name),
		`-rotate 0.5`,
		`-attenuate 0.7`,
		`+noise Multiplicative`,
		`-colorspace Gray`,
		fmt.Sprintf("%s/%s.pdf", input.dir, input.name),
	).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil

}


func checkBinary(name string) error {
	_, err := exec.LookPath(name)
	return err
}


func findTemplate(name string) (*file, error) {
	/*var templates []string

	err := filepath.Walk(
		TEMPLATEDIR,
    func(path string, info os.FileInfo, err error) error {
      if err != nil {
          return err
      }

      if filepath.Ext(path) == `.tex` {
        templates = append(templates, path)
      }
      return nil
  })
	if err != nil {
		return nil, err
	} else if len(templates) == 0 {
		return nil, errors.New("no templates found")
	}*/

	return findFile(fmt.Sprintf("%s/%s.tex", TEMPLATEDIR, name))
}


func findFile(input string) (*file, error) {
	if ! checkFile(input) {
		return nil, fmt.Errorf("file does not exist: %s", input)
	}

	path, err := filepath.Dir(input)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(input)

	return &file{
		name: strings.TrimSuffix(input, ext),
		dir: path,
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
