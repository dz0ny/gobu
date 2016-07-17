package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func untar(source, target string) {

	file, err := os.Open(source)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(source, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			log.Fatal(err)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// get the individual filename and extract to the current directory
		filename := filepath.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			log.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				log.Fatal(err)
			}

		case tar.TypeReg:
			// handle normal file
			log.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				log.Fatal(err)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				log.Fatal(err)
			}

			writer.Close()
		default:
			log.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}

}
