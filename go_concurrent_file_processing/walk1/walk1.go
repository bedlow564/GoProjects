package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type pair struct {
	hash string
	path string
}

type fileList []string

type results map[string]fileList

func hashFile(path string) pair {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths { //this is blocked until paths is written to?
		pairs <- hashFile(path) //get hash of file // write to pair channel //allows collect hashes to begin processing?
	}

	done <- true
}

func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path) //append files to slice containing same hash
	}

	result <- hashes //wirte hashes to results channel
}

func searchTree(dir string, paths chan<- string) error {
	//put closure in a variable
	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			paths <- p //write path to paths channel
		}

		return nil
	}

	return filepath.Walk(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0) //how many threads should go create //double number because all threads wont be busy all the time
	paths := make(chan string)           //paths channel
	pairs := make(chan pair)             //pairs channel (hash, path)
	done := make(chan bool)              //done channel
	result := make(chan results)         //results channel (map of hashes to file paths)

	for range workers { //create go routine for how many workers there are
		go processFiles(paths, pairs, done)
	}

	//we need another go routine so we dont block here
	go collectHashes(pairs, result)

	if err := searchTree(dir, paths); err != nil {
		return nil
	}

	close(paths) //close paths channel to signal that no more paths will be sent
	//workers dont have to block and wait for more paths

	//wait for all workers to be done
	for range workers {
		<-done
	}

	//by closing pairs we signal tht all the hashes
	//have been collected; we have to do it here AFTER
	//all the workers are done
	close(pairs)

	return <-result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide dir name!")
	}

	if hashes := run(os.Args[1]); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				// we will use just 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("  ", file)
				}
			}
		}
	}
}
