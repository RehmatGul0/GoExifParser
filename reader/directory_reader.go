package reader

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"imageprocessor/utils"
)

type DirectoryReaderStrategy struct {
	files       []string
	fileFilters []string
}

// Read traverses the directory tree rooted at the specified root path.
// For each file or directory encountered, it applies the provided filtering
// logic to determine if the file should be included in the DirectoryReaderStrategy's
// file list. If an error occurs during traversal, it returns a formatted error.
//
// Parameters:
// - root: The root directory path to start reading from.
//
// Returns:
// - error: An error if any occurs during the directory traversal.
func (dr *DirectoryReaderStrategy) Read(root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking on file path %s, error: %v", path, err)
		}
		shouldFilter := func(fileName string) bool {
			return fileName == d.Name()
		}
		if !d.IsDir() && !utils.Contains(dr.fileFilters, shouldFilter) {
			dr.files = append(dr.files, path)
		}
		return nil
	})
}

func (dr *DirectoryReaderStrategy) setFileFilters(fileFilters []string) {
	dr.fileFilters = fileFilters
}

func NewDirectoryReader() FileReader {
	dr := &DirectoryReaderStrategy{}
	dr.setFileFilters([]string{".DS_Store"})
	return dr
}

func GetFiles(dr FileReader) []string {
	if r, ok := dr.(*DirectoryReaderStrategy); ok {
		return r.files
	}
	return []string{}
}
