package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/johnstarich/godoc-static/gen"
	"github.com/johnstarich/godoc-static/static"
)

const out = "dist"

func main() {
	var srcPath string
	flag.StringVar(&srcPath, "path", "", "path to the `source`")
	var mod string
	flag.StringVar(&mod, "mod", "", "path of the `module`")
	keyValues := StringSlice("meta", nil, "Key=value pairs for customizing generated content")
	flag.Parse()

	if srcPath == "" || mod == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	meta := make(map[string]string)
	for _, pair := range *keyValues {
		tokens := strings.SplitN(pair, "=", 2)
		if len(tokens) < 2 {
			fmt.Printf("Invalid key-value pair: %s. Must use 'key=value' format.\n", pair)
			os.Exit(2)
		}
		meta[tokens[0]] = tokens[1]
	}

	render := gen.NewRenderer(srcPath, mod, out, meta)
	if err := render.GenerateAll("/", out); err != nil {
		log.Fatal(err)
	}

	if err := static.OutputResources(path.Join(out, "lib/godoc")); err != nil {
		log.Fatal(err)
	}

	if err := render.Make404("404.html"); err != nil {
		log.Fatal(err)
	}
}
