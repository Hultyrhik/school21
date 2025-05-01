package main

import (
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func main() {
	dest := flag.String("a", "", "use -a to specife location where archive will be saved")
	flag.Parse()

	var files []string
	if (*dest) == "" {
		files = os.Args[1:]
	} else {
		files = os.Args[3:]
	}

	if len(files) == 0 {
		log.Fatal("File is not specified")
	}

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			info, err := os.Stat(file)
			if err != nil {
				log.Fatal(err)
			}
			mtime := info.ModTime().Unix()

			basename := path.Base(file)
			filename := strings.TrimSuffix(basename, filepath.Ext(basename))
			// fmt.Println(filename, mtime)

			archiveName := filename + "_" + strconv.Itoa(int(mtime)) + ".tar.gz"
			// fmt.Println(archiveName)
			inputFile, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer inputFile.Close()

			var path string
			if *dest != "" {
				path = *dest + "/"
			}
			gzipWriter, err := os.Create(path + archiveName)
			if err != nil {
				log.Fatal(err)
			}
			defer gzipWriter.Close()

			zipWriter := gzip.NewWriter(gzipWriter)
			defer zipWriter.Close()

			_, err = io.Copy(zipWriter, inputFile)
			if err != nil {
				log.Fatal(err)
			}

			zipWriter.Close()
		}()
	}
	wg.Wait()
}
