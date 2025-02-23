package writer

import (
	"errors"
	"io"
)

type CustomIOWriter interface {
	io.StringWriter
	io.Writer
}

type Writer interface {
	Write(file CustomIOWriter, metadata []map[string]string) error
}

func newCSVWriter() Writer {
	return new(CSVExifWriterStrategy)
}

func newHTMLWriter() Writer {
	return new(HTMLExifWriterStrategy)
}

func NewWriter(format string) (Writer, error) {
	var w Writer
	switch format {
	case ".csv":
		w = newCSVWriter()
	case ".html":
		w = newHTMLWriter()
	default:
		return nil, errors.New("incorrect file format provided")
	}
	return w, nil
}
