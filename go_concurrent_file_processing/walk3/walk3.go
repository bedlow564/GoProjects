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

func processFile(path string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()

	limits <- true //use limits channel to control how many goroutines can interact with file disk

	//defer can only be used on a func so wrap reading from limits in an anonymous func (closure)
	defer func() {
		<-limits
	}()

	pairs <- hashFile(path)

}

func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

func searchTree(dir string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) error {
	defer wg.Done()

	visit := func(p string, fi os.FileInfo, err error) error {

		if err != nil && err != os.ErrNotExist {
			return err
		}

		if fi.Mode().IsDir() && p != dir {
			wg.Add(1)
			go searchTree(p, pairs, wg, limits)
			return filepath.SkipDir
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			wg.Add(1)
			go processFile(p, pairs, wg, limits)
		}

		return nil

	}

	limits <- true //user limits to control how many goroutines can interact with file system

	defer func() {
		<-limits
	}()

	return filepath.Walk(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)
	limits := make(chan bool, workers) //buffer channels (i/o and network bound work fights for contention of shared resoures so)
	//so limit goroutines can help performance
	pairs := make(chan pair, workers) //buffer channels
	result := make(chan results)
	wg := new(sync.WaitGroup)

	//we need another goroutine so we dont block here
	go collectHashes(pairs, result)

	//multi-threaded walk of the directory tree; we need a
	//waitGroup because we dont know how many to wait for
	wg.Add(1)

	err := searchTree(dir, pairs, wg, limits)

	if err != nil {
		log.Fatal(err)
	}

	//we must close the paths channel so the workers stop
	wg.Wait()

	//by closing pair we signal that all the hashes
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
				//we will use last 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))
				for _, file := range files {
					fmt.Println(" ", file)
				}
			}
		}
	}

}
