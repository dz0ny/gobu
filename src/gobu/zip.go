package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unzip(source, target string) {

	r, err := zip.OpenReader(source)

	defer r.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(target, 0755); err != nil {

		log.Fatal(err)
	}

	for _, file := range r.File {

		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			log.Println("Creating directory :", file.Name)
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Fatal(err)
		}

		defer targetFile.Close()

		log.Println("Unzipping :", file.Name)

		if _, err := io.Copy(targetFile, fileReader); err != nil {

			log.Fatal(err)
		}

	}

}
