package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
)

func readZipFile() *zip.ReadCloser {
	read, err := zip.OpenReader(zipFileName)
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	return read
}

func readCSVFile(file *zip.File) (*csv.Reader, error) {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return nil, fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	lines, err := ioutil.ReadAll(fileread)
	r := csv.NewReader(bytes.NewBuffer(lines))
	return r, nil
}
