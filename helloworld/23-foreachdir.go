package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

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

	// t := 0
	// n := 0
	// for s := range size {
	// 	n++
	// 	t += int(s)
	// }

	// 每隔一秒打印文件大小变化
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case s, ok := <-size:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += s

		case <-tick:
			printSize(nfiles, nbytes)
		}
	}

	printSize(nfiles, nbytes)
}

func printSize(fileCount int64, size int64) {
	fmt.Printf("%d files %.2f MB\n", fileCount, float64(size)/1024/1024)
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
