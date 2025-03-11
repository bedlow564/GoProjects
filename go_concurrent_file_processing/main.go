package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type pair struct {
	hash, path string
}

type fileList []string
type results map[string]fileList

func hashFile(path string) pair {
	file, err := os.Open(path) //open file

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //make sure file closes

	hash := md5.New() //not usually secure but good enough and fast

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path} //create string of hash using sprintf (%x) base 16 (hexadecimal)

}

func searchTree(dir string) (results, error) {
	hashes := make(results) //create a map of result types

	//walk does a walk of the given directory
	err := filepath.Walk(dir, func(p string, fi os.FileInfo, err error) error {

		//ignore files that are empty because they all have the same hash //check for regular files
		if fi.Mode().IsRegular() && fi.Size() > 0 {
			h := hashFile(p)
			hashes[h.hash] = append(hashes[h.hash], h.path) // add file to fileList slice //ok to append to a nil slice
		}

		return nil

	})

	return hashes, err

}

func main() {

}
