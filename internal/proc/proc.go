// Package proc provides functionality to check if a file is currently being written to.
//
// This helps to not perform operations on files that are still in use.
package proc

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// WalkProc walks through the /proc directory and builds a "set" of files in the target directory
// that are currently open for use as indicated by /proc/*/fd.
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
		// create the filepath for the fd entry to check whether the file is in use
		fdDirName := filepath.Join(procDir, entry.Name(), "fd")
		symlinks, err := os.ReadDir(fdDirName)
		if err != nil {
			continue
		}
		// start loops to read symbolic links and check if they belong to the target directory
		for _, link := range symlinks {
			truePath, err := os.Readlink(filepath.Join(fdDirName, link.Name()))
			// if there was an error reading the symbolic link or they do not belong to the target
			// directory then they can safely be skipped
			if err != nil || !strings.HasPrefix(truePath, checkingPath) {
				continue
			}
			filesInUse[truePath] = true
		}
	}
	return filesInUse
}

// IsAvailable checks if a specific file is within the files currently being used.
//
// openFiles is a "set" (actually a map) containing the results from WalkProc and filePath is the path
// to the file to be checked for usage.
func IsAvailable(filePath string, openFiles map[string]bool) bool {
	if _, ok := openFiles[filePath]; ok {
		log.Printf("ERROR: %s - file in use by another process", filePath)
		return false
	}
	return true
}
