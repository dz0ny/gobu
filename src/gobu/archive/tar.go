package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Untar unpacks tar.gz archive
func Untar(source, target string) error {

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
			return err
		}

		// get the individual filename and extract to the current directory
		filename := filepath.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			log.Debugln("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				return err
			}

		case tar.TypeReg:
			// handle normal file
			log.Debugln("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				return err
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				return err
			}

			writer.Close()
		default:
			log.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}
	return nil
}
