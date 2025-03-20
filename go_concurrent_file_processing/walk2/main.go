package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
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
		log.Fatal()
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths { //read from paths channel and write hash of file into pairs channel
		pairs <- hashFile(path)
	}
	//NOTE: This channel will block until something is written to paths channel
	done <- true
}

func collectHashes(pair <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pair { //read from pairs channel and path of file to hashes map
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	//when all pairs are read write hashes map to result channel
	result <- hashes
}

// WaitGroup has to be passed as a pointer
func searchTree(dir string, paths chan<- string, wg *sync.WaitGroup) error {
	defer wg.Done() //defer calling done on the wait group

	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		//ignore dir itself to avoid an infinte loop!
		if fi.Mode().IsDir() && p != dir { //. is the dir being searched
			wg.Add(1)                   //increment WaitGroup
			go searchTree(p, paths, wg) //launch go routine that searches dir
			return filepath.SkipDir     //return from directory once go routine for it is launched
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 { //check if file is not empty and no mode type bits are set
			paths <- p //write path name to paths channel
		}

		return nil
	}

	return filepath.Walk(dir, visit)

}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0) //Get number of CPU cores for parallel processing. 0 mean dont change current value just return it
	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	result := make(chan results)
	wg := new(sync.WaitGroup) //wait group measures unit of work

	for range workers { // launch go routine for amount of workers
		go processFiles(paths, pairs, done)
	}

	//we need another go routine so we dont block here
	go collectHashes(pairs, result)

	//multi-threaded walk of the directory tree; we need a
	//waitGroup because we dont know how many to wait for
	wg.Add(1)

	err := searchTree(dir, paths, wg)

	if err != nil {
		log.Fatal(err)
	}

	//we must close the paths channel so the workers stop
	wg.Wait() // blocks until wait group is 0 (all work is complete)
	close(paths)

	//wait for all the workers to be done
	for range workers {
		<-done
	}

	//by closing pairs we signal that all the hashes
	//have been collected; we have to do it here AFTER
	//all the workers are done
	close(pairs)

	return <-result
}

func main() {
	if len(os.Args) < 2 { //check if dir was provided to command line
		log.Fatal("Missing parameter, provide dir name!")
	}

	if hashes := run(os.Args[1]); hashes != nil { //get result from run func. Print if hashes isnt nil
		for hash, files := range hashes { //loop through hashes (hash and files)
			if len(files) > 1 { //if files is found is more than 1
				//we will use the last 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files { //loop through files and print file names
					fmt.Println(" ", file)
				}
			}
		}
	}
}
