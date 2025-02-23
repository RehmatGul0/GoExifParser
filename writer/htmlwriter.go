package writer

import (
	"errors"
	"fmt"
)

type HTMLExifWriterStrategy struct {
}

// Write writes the provided metadata to the given file in HTML format.
//
// Parameters:
//   - file: The file to write the HTML content to.
//   - metadata: A slice of maps containing the metadata to be written.
//
// Returns:
//   - err: An error if something goes wrong during the writing process.
func (htmlWriter *HTMLExifWriterStrategy) Write(file CustomIOWriter, metadata []map[string]string) (err error) {
	if len(metadata) == 0 {
		return errors.New("no data to write to csv file")
	}
	_, err = file.WriteString("<html>\n<head>\n<title>Exif Metadata</title>\n</head>\n<body>\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("<table>\n<tr>\n")
	if err != nil {
		return err
	}
	for key := range metadata[0] {
		_, err = file.WriteString(fmt.Sprintf("<th>%s</th>", key))
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("</tr>\n")
	if err != nil {
		return err
	}
	for _, data := range metadata {
		_, err = file.WriteString("<tr>\n")
		if err != nil {
			return err
		}
		for _, datum := range data {
			_, err = file.WriteString(fmt.Sprintf("<td>%s</td>", datum))
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString("</tr>\n")
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("</table>\n</body>\n</html>")
	if err != nil {
		return err
	}

	return nil
}
