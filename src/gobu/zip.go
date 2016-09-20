package main

import (
"archive/zip"
"io"
"log"
"os"
"path/filepath"
"strings"
)

func zip(source, target string) {

	r, err := zip.OpenReader(source)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	for _, f := range r.File {

		fmt.Printf("%v", f)
	}


}
