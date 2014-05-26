package main

import (
	"errors"
	"flag"
	"log"
	"math/rand"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

import (
	"github.com/skratchdot/open-golang/open"
)

var mimeFilter string

func init() {
	const mimeFilterUsage = "the mime-type to filter"
	flag.StringVar(&mimeFilter, "mime", "", mimeFilterUsage)
	flag.StringVar(&mimeFilter, "m", "", mimeFilterUsage+" (shorthand)")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "."
	}
	filenames, err := WalkByMime(root, mimeFilter)
	if err != nil {
		log.Fatal(err)
	}
	filename := filenames[rand.Intn(len(filenames))]
	log.Println("Opening " + filename)
	err = open.Run(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func WalkByMime(root, mimeFilter string) ([]string, error) {
	var paths []string
	err := filepath.Walk(root, func(currentPath string, f os.FileInfo, err error) error {
		if err != nil || f.IsDir() {
			return err
		}
		if mimeFilter == "" || strings.HasPrefix(mime.TypeByExtension(path.Ext(currentPath)), mimeFilter) {
			paths = append(paths, currentPath)
		}
		return nil
	})
	if err == nil && len(paths) < 1 {
		return paths, errors.New("No matching files")
	}
	return paths, err
}
