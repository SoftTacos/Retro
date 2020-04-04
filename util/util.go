package util

import (
	"io/ioutil"
	"log"
	"os"
)

func LoadTextFile(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	filestring, err := ioutil.ReadAll(f)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return filestring
}
