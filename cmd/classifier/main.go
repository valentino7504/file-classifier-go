package main

import (
	"log"
	"os"

	"github.com/valentino7504/file-classifier-go/internal/classify"
	"github.com/valentino7504/file-classifier-go/internal/proc"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("no directory provided")
	}
	dirName := args[1]
	if dir, err := os.Stat(dirName); os.IsNotExist(err) {
		log.Fatal("directory provided does not exist")
	} else if err != nil {
		log.Fatal("problems checking directory")
	} else if !dir.IsDir() {
		log.Fatal("path provided is not a directory")
	}
	openFiles := proc.WalkProc(dirName)
	classify.Classify(dirName, openFiles)
}
