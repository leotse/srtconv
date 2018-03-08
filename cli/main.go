package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"github.com/leotse/srtfix"
)

func main() {

	// single file mode
	if len(os.Args) == 2 {
		filepath := os.Args[1]

		converter := srtfix.NewFileConverter(
			NewFileReader(filepath),
			NewStdoutWriter(),
			NewParser(),
			NewResolver(),
		)

		err := converter.Convert()
		if err != nil {
			log.Panic(err)
		}

		os.Exit(0)
	}

	// dir mode
	if len(os.Args) == 3 {
		indir := os.Args[1]
		outdir := os.Args[2]

		// ensure outdir exists
		err := os.MkdirAll(outdir, 0755)
		if err != nil {
			log.Panic(err)
		}

		// read indir
		files, err := ioutil.ReadDir(indir)
		if err != nil {
			log.Panic(err)
		}

		// convert all srt files
		wg := &sync.WaitGroup{}
		for _, file := range files {
			name := file.Name()
			if !file.IsDir() && path.Ext(name) == ".srt" {
				inpath := path.Join(indir, name)
				outpath := path.Join(outdir, name)
				convertFile(inpath, outpath, wg)
				wg.Add(1)
			}
		}

		// wait for all conversions to complete
		wg.Wait()
		os.Exit(0)
	}

	usage()
}

func convertFile(inpath, outpath string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		err := convertFileSync(inpath, outpath)
		if err != nil {
			log.Println(err)
		}
	}()
}

func convertFileSync(inpath, outpath string) error {
	converter := srtfix.NewFileConverter(
		NewFileReader(inpath),
		NewFileWriter(outpath),
		NewParser(),
		NewResolver(),
	)
	return converter.Convert()
}

func usage() {
	fmt.Println()
	fmt.Println("\tUsage (1): srtfix [input_srt_path]")
	fmt.Println()
	fmt.Println("\tUsage (2): srtfix [input_dir] [output_dir]")
	fmt.Println()
}
