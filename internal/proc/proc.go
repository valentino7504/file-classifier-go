// Package proc provides functionality to check if a file is currently being written to.
//
// This helps to not perform operations on files that are still in use.
package proc

import (
	"log"
	"os"
	"strings"
)

func WalkProc(checkingPath string) map[string]bool {
	const procDir string = "/proc"
	filesInUse := make(map[string]bool)
	processes, err := os.ReadDir(procDir)
	if err != nil {
		return filesInUse
	}
	for _, entry := range processes {
		// if the entry is not a directory then there is obviously not a /fd directory to be read
		if !entry.IsDir() {
			continue
		}
		fdDirName := filepath.Join(procDir, entry.Name(), "fd")
		symlinks, err := os.ReadDir(fdDirName)
		if err != nil {
			continue
		}
		for _, link := range symlinks {
			truePath, err := os.Readlink(filepath.Join(fdDirName, link.Name()))
			if err != nil || !strings.HasPrefix(truePath, checkingPath) {
				continue
			}
			filesInUse[truePath] = true
		}
	}
	return filesInUse
}

func IsAvailable(filePath string, openFiles map[string]bool) bool {
	if _, ok := openFiles[filePath]; ok {
		log.Printf("ERROR: %s - file in use by another process", filePath)
		return false
	}
	return true
}
