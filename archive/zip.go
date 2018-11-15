package archive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Unzip unpacks .zip archive
func Unzip(source, target string) error {

	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer r.Close()

	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range r.File {

		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			log.Debugln("Creating directory :", file.Name)
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		defer targetFile.Close()

		log.Debugln("Unzipping :", file.Name)

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}

	}
	return nil
}
