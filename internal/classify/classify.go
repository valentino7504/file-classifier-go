// Package classify provides functionality for actually running the movement
// of files within the target directory.
package classify

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/valentino7504/file-classifier-go/internal/proc"
)

// moveFile is a simple implementation of an overwrite-safe way to move the files as necessary.
func moveFile(src string, dest string) error {
	finalPath, ext := dest, filepath.Ext(dest)
	for i := 1; ; i++ {
		if _, err := os.Stat(finalPath); os.IsNotExist(err) {
			break
		} else if err != nil {
			return err
		}
		finalPath = fmt.Sprintf("%s(%d)%s", strings.TrimSuffix(dest, ext), i, ext)
	}
	return os.Rename(src, finalPath)
}

// Classify does the actual classification of files in the target directory.
//
// It creates the subdirectories if they do not exist, and then loops through each item in the target directory.
// If the entry is a subdirectory, is open for use by another process or has the .part extension it is skipped.
// Otherwise if it is a file with an extension that I have defined a destination for, then it is moved to the
// appropriate destination.
func Classify(basePath string, openFiles map[string]bool) {
	// create the subdirectories
	for _, folder := range folders {
		dirPath := filepath.Join(basePath, folder)
		err := os.MkdirAll(dirPath, 0o755)
		if err != nil {
			log.Printf("ERROR: creating directory failed - %s", dirPath)
			return
		}
	}
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return
	}
	for _, entry := range entries {
		entryName := entry.Name()
		// if the entry is a directory or has .part in its name
		if entry.IsDir() || strings.HasSuffix(entryName, ".part") {
			continue
		}
		// if the file is open for use by another process then skip it
		filePath := filepath.Join(basePath, entryName)
		if !proc.IsAvailable(filePath, openFiles) {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entryName))
		if dir, ok := extensions[ext]; ok {
			destPath := filepath.Join(basePath, dir, entryName)
			err = moveFile(filePath, destPath)
			if err != nil {
				log.Printf("ERROR: moving file failed - %s to %s", filePath, destPath)
				continue
			}
			log.Printf("SUCCESS moved %s to %s", filePath, destPath)
		}
	}
}
