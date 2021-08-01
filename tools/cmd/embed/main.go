package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unsafe"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v2"
)

func main() {
	log.SetPrefix("idgen: ")
	log.SetFlags(0)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(wd); err != nil {
		log.Fatal(err)
	}
}

func run(wd string) error {
	input := flag.String("i", "", "input")
	output := flag.String("o", "", "output")
	name := flag.String("n", "", "const name")
	pkgname := flag.String("p", "", "package name")
	yamltojson := flag.Bool("yaml2json", false, "convert YAML to JSON")
	all := flag.Bool("all", false, "read all files")
	flag.Parse()

	if *input == "" {
		return errors.New("input option is required")
	}

	if *output == "" && !*all {
		return errors.New("output option is required")
	}

	if *name == "" {
		return errors.New("name option is required")
	}

	pkgs, err := packages.Load(&packages.Config{Dir: wd}, ".")
	if err != nil {
		return errors.Wrap(err, "failed to load package")
	}

	if *pkgname == "" {
		pkgname = &pkgs[0].Name
	}

	if *all {
		filename, ext := getFileNameExt(*input)
		filenames, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		for _, f := range filenames {
			if strings.HasPrefix(f.Name(), filename) && strings.HasSuffix(f.Name(), ext) {
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("failed to read file %s", f.Name()))
				}
				fstr, _ := getFileNameExt(f.Name())
				err = handleOneFile(f.Name(), fstr+"_gen.go", *pkgname, *name+fstr[len(filename):], yamltojson)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("failed to handle file %s", f.Name()))
				}
			}
		}
	} else {
		return handleOneFile(*input, *output, *pkgname, *name, yamltojson)
	}
	return nil
}

type templateData struct {
	PackageName string
	Name        string
	Content     string
}

var templ = template.Must(template.New("generated").Parse(`
// Code generated by github.com/reearth/reearth-backend/tools/cmd/embed, DO NOT EDIT.

package {{.PackageName}}

const {{.Name}} string = ` + "`{{.Content}}`" + ``))

// https://stackoverflow.com/questions/40737122/convert-yaml-to-json-without-struct
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func yaml2json(content []byte) ([]byte, error) {
	var y interface{}
	if err := yaml.Unmarshal([]byte(content), &y); err != nil {
		return nil, errors.Wrap(err, "failed to parse YAML")
	}
	y = convert(y)
	b, err := json.Marshal(&y)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marhsal JSON")
	}
	return b, nil
}

func processAndWriteOneFile(data templateData, output string) error {
	buf := &bytes.Buffer{}

	if err := templ.Execute(buf, data); err != nil {
		return errors.Wrap(err, "unable to generate code")
	}

	src, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return errors.Wrap(err, "unable to gofmt")
	}

	err = os.WriteFile(output, src, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}
	return nil
}

func handleOneFile(input, output, pkgname, name string, yamltojson *bool) error {
	content, err := os.ReadFile(input)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}

	if yamltojson != nil && *yamltojson {
		content, err = yaml2json(content)
		if err != nil {
			return errors.Wrap(err, "failed to read file")
		}
	}

	contentstr := *(*string)(unsafe.Pointer(&content))

	data := templateData{
		PackageName: pkgname,
		Name:        name,
		Content:     strings.ReplaceAll(contentstr, "`", "` + \"`\" + `"),
	}
	return processAndWriteOneFile(data, output)
}

func getFileNameExt(input string) (string, string) {
	ext := filepath.Ext(input)
	fname := input[0 : len(input)-len(ext)]
	return fname, ext
}