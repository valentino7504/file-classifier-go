package classify

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/valentino7504/file-classifier-go/internal/proc"
)

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

func Classify(basePath string, openFiles map[string]bool) {
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
		if entry.IsDir() || strings.HasSuffix(entryName, ".part") {
			continue
		}
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
