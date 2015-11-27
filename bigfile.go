package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	walk(os.Args[1])
}

func walk(dirname string) {
	f, err := os.Open(dirname)
	defer f.Close()

	if err != nil {
		log.Print(err)
		return
	}

	infos, err := f.Readdir(0)

	if err != nil {
		log.Print(err)
		return
	}

	max := struct {
		size int64
		name string
	}{}

	for _, info := range infos {
		pathname := path.Join(dirname, info.Name())
		if info.IsDir() {
			walk(pathname)
			continue
		}
		if max.size < info.Size() {
			max.name = pathname
			max.size = info.Size()
		}
	}

	if max.name == "" {
		return
	}

	fmt.Printf("%q %d\n", max.name, max.size)
}
