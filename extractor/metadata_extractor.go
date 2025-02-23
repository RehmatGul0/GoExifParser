package extractor

import (
	"fmt"
	"os"

	"imageprocessor/utils"

	"github.com/rwcarlsen/goexif/exif"
)

type FileMetadataExtractor interface {
	Extract(filePath string) (map[string]string, error)
}

type fileExifExtractorStrategy struct {
	walker exif.Walker
}

func (exifExtractor *fileExifExtractorStrategy) Extract(filePath string) (map[string]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err.Error())
	}

	exifMetadata, err := exif.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("error decoding exif metadata: %s", err.Error())
	}
	exifMetadata.Walk(exifExtractor.walker)
	metadata := exifExtractor.walker.(utils.ExifWalker).GetMetadata()
	return exifExtractor.filterMetadata(metadata), nil
}

func (exifExtractor *fileExifExtractorStrategy) filterMetadata(metadata map[string]string) map[string]string {
	exifFilterTags := []exif.FieldName{exif.GPSLongitude, exif.GPSLatitude}
	for key := range metadata {
		filterFunc := func(filterTag exif.FieldName) bool {
			return string(filterTag) == key
		}
		if !utils.Contains(exifFilterTags, filterFunc) {
			delete(metadata, key)
		}
	}
	return metadata
}

// NewFileExifExtractor returns a new instance of FileMetadataExtractor using the fileExifExtractorStrategy.
func NewFileExifExtractor() FileMetadataExtractor {
	return &fileExifExtractorStrategy{walker: utils.NewExifWalker()}
}
