package writer

import (
	"encoding/csv"
	"errors"
	"fmt"
)

type CSVExifWriterStrategy struct{}

// Write writes the provided metadata to the given CSV file.
//
// Parameters:
//   - file: A pointer to the os.File where the CSV data will be written.
//   - metadata: A slice of maps, where each map represents a row of metadata with string keys and values.
//
// Returns:
//   - error: An error if any issue occurs during the writing process, otherwise nil.
func (csvWriter *CSVExifWriterStrategy) Write(file CustomIOWriter, metadata []map[string]string) error {
	if len(metadata) == 0 {
		return errors.New("no data to write to csv file")
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var keys []string
	for key := range metadata[0] {
		keys = append(keys, key)
	}
	if err := writer.Write(keys); err != nil {
		return fmt.Errorf("writing headers to csv file, error: %s", err.Error())
	}
	for _, data := range metadata {
		var row []string
		for _, datum := range data {
			row = append(row, datum)
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing rows to csv file, error: %s", err.Error())
		}
	}
	return nil
}
