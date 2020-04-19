package main

import (
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
	ext  string
}

func init() {
	flag.BoolVarP(&versionFlg, "version", "v", false, "Print version and exit")
	flag.BoolVarP(&xeroxFlg, "xerox", "x", false, "Xerox the output")
	flag.StringVarP(&templateName, "template", "t", "", "Template name")

	flag.Parse()

	if versionFlg {
		fmt.Printf("Version: %s (%s)\n", version, commit)
		os.Exit(0)
	}

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
	mdFile, err := findFile(inputName)
	if err != nil {
		log.Fatal(err)
	}

	template, err := findTemplate(templateName)
	if err != nil {
		log.Fatal(err)
	}

	err = markup(mdFile, template)
	if err != nil {
		log.Fatal(err)
	}

	if xeroxFlg {
		err = xerox(mdFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}


func markup(input *file, template *file) error {
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


func xerox(input *file) error {
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


func findTemplate(name string) (*file, error) {
	return findFile(fmt.Sprintf("%s/%s.tex", TEMPLATEDIR, name))
}


func findFile(input string) (*file, error) {
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
