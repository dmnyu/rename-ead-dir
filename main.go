package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type EAD struct {
	XMLName xml.Name  `xml:"ead"`
	Header  EADHeader `xml:"eadheader"`
}

type EADHeader struct {
	EADID string `xml:"eadid"`
}

func main() {
	rootDir := os.Args[1]
	fi, err := os.Stat(rootDir)
	if err != nil {
		panic(err.Error())
	}

	if fi.IsDir() != true {
		panic(fmt.Errorf("%s is not a directory", rootDir))
	}

	eads, err := os.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}

	for _, ead := range eads {

		origFile := filepath.Join(rootDir, ead.Name())
		eadid, err := getEADID(origFile)
		if err != nil || eadid == nil {
			log.Println(err.Error())
			continue
		}

		newFile := filepath.Join(rootDir, (*eadid + ".xml"))
		fmt.Println(origFile, newFile)

		if err := os.Rename(origFile, newFile); err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func getEADID(path string) (*string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	ead := EAD{}

	if err := xml.Unmarshal(fileBytes, &ead); err != nil {
		return nil, err
	}

	return &ead.Header.EADID, nil
}
