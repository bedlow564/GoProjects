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
	//create clousre for WalkFunc
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

	if len(os.Args) < 2 { //provide dir name
		log.Fatal("Missing parameter, provide dir name!")
	}
	if hashes, err := searchTree(os.Args[1]); err == nil { //search given directory
		for hash, files := range hashes { //loop through map of files to hashes
			if len(files) > 1 {
				//we will use just 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files)) //print first 7 chars of hash and number of files

				for _, file := range files { //print all files that the have same hash
					fmt.Println(" ", file)
				}
			}
		}
	}
}
