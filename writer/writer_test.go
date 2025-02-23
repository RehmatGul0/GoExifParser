package writer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWriter(t *testing.T) {
	tests := []struct {
		format string
		wtr    Writer
	}{
		{".csv", &CSVExifWriterStrategy{}},
		{".html", &HTMLExifWriterStrategy{}},
		{".txt", nil},
	}

	for _, test := range tests {
		wtr, err := NewWriter(test.format)
		if err != nil {
			assert.ErrorContains(t, err, "incorrect file format provided")
			assert.Nil(t, wtr)
		} else {
			assert.IsType(t, wtr, test.wtr)
		}
	}
}

func TestCSVExifWriterStrategy_Write(t *testing.T) {
	tests := []struct {
		name     string
		metadata []map[string]string
		wantErr  bool
		data     string
	}{
		{
			name:     "Empty metadata",
			metadata: []map[string]string{},
			wantErr:  false,
			data:     "",
		},
		{
			name: "Valid metadata",
			metadata: []map[string]string{
				{"key1": "value1", "key2": "value2"},
			},
			wantErr: false,
			data:    "key1,key2\nvalue1,value2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			csvWriter := &CSVExifWriterStrategy{}
			err := csvWriter.Write(&buf, tt.metadata)
			if (err != nil) && tt.wantErr {
				t.Errorf("CSVExifWriterStrategy.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.data, buf.String())

		})
	}
}
