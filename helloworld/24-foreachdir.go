package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// 循环文件夹，优化一下

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	size := make(chan int64)

	go func() {
		for _, dir := range roots {
			walkDir(dir, size)
		}
		close(size)
	}()

	t := 0
	n := 0
	for s := range size {
		n++
		t += int(s)
	}

	totalSize := fmt.Sprintf("%.2f KB", float64(t)/1024)
	fmt.Println("totalFiles: ", n)
	fmt.Println("totalSize: ", totalSize)

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-size:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return []os.FileInfo{}
	}
	return entries
}

func walkDir(dir string, size chan<- int64) {
	fileinfo := dirents(dir)
	for _, entry := range fileinfo {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, size)
		} else {
			size <- entry.Size()
		}
	}
	return
}
