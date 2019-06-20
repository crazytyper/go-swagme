package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/crazytyper/swagme"
	"github.com/crazytyper/swagme/gengo"
)

var (
	packageNameOpt = flag.String("p", "myapi", "Package name")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

	specfile := args[0]

	var err error
	var spec *swagme.Spec

	if spec, err = readSpecFile(specfile); err != nil {
		log.Fatal(err)
	}

	w := gengo.NewWriter(os.Stdout)

	if err = gengo.GenerateModels(w, *packageNameOpt, spec); err != nil {
		log.Fatal(err)
	}
}

func readSpecFile(filename string) (*swagme.Spec, error) {
	specdata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var spec swagme.Spec
	if err := json.Unmarshal(specdata, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
