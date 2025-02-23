package utils

import (
	"errors"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type ExifWalker struct {
	metadata map[string]string
}

// Walk processes a given EXIF field name and its corresponding TIFF tag.
// It stores the string representation of the tag in the metadata map of the Walker.
// If the provided tag is nil, it returns an error indicating an invalid tag.
//
// Parameters:
//   - name: The EXIF field name to be processed.
//   - tag: A pointer to the TIFF tag associated with the EXIF field name.
//
// Returns:
//   - error: An error if the provided tag is nil, otherwise nil.
func (w ExifWalker) Walk(name exif.FieldName, tag *tiff.Tag) error {
	if tag == nil {
		return errors.New("invalid tag provided value is nil")
	}
	w.metadata[string(name)] = tag.String()
	return nil
}

// GetMetadata returns the metadata associated with the Walker instance.
// The metadata is represented as a map where the keys and values are strings.
func (w ExifWalker) GetMetadata() map[string]string {
	return w.metadata
}

func NewExifWalker() ExifWalker {
	w := ExifWalker{}
	w.metadata = make(map[string]string)
	return w
}
