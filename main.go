package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"imageprocessor/extractor"
	"imageprocessor/reader"
	"imageprocessor/utils"
	"imageprocessor/writer"
)

func main() {
	// main.go is a command-line application that scans a specified directory for image files,
	// extracts their EXIF metadata, and writes the metadata to an output file in CSV or HTML format.
	//
	// Flags:
	// -dir: Directory to scan for images. Defaults to the current directory.
	// -output: Output file name. Can be either "output.csv" or "output.html". Defaults to "output.csv".
	//
	// The application uses the following components:
	// - DirectoryReader: Reads the specified directory and retrieves a list of image files.
	// - Writer: Writes the extracted metadata to the specified output file format.
	// - FileExifExtractor: Extracts EXIF metadata from image files.
	//
	// The main workflow includes:
	// 1. Parsing command-line flags to get the directory to scan and the output file format.
	// 2. Reading the directory to get a list of image files.
	// 3. Creating the output file.
	// 4. Iterating over the list of image files, extracting their EXIF metadata, and appending it to a metadata slice.
	// 5. Writing the collected metadata to the output file.
	//
	// If an error occurs while creating the output file or extracting EXIF metadata, the application logs the error and continues processing the remaining files.
	directory := flag.String("dir", ".", "Directory to scan for images")
	outputFile := flag.String("output", "output.csv or output.html", "Output file")
	flag.Parse()

	fileFormat := filepath.Ext(*outputFile)

	file, err := utils.CreateFile(*outputFile)
	defer func() {
		file.Close()
	}()
	if err != nil {
		fmt.Println("Error creating file to write metadata: ", err.Error())
		return
	}

	directoryReader := reader.NewDirectoryReader()
	writer, writerErr := writer.NewWriter(fileFormat)
	if writerErr != nil {
		fmt.Println("Error: ", writerErr.Error())
		return
	}

	directoryReader.Read(*directory)
	files := reader.GetFiles(directoryReader)

	exifExtractor := extractor.NewFileExifExtractor()
	var metadata []map[string]string
	for i := range files {
		meta, err := exifExtractor.Extract(files[i])
		if err != nil {
			fmt.Println("Error extracting exif metadata:", err.Error())
			continue
		}
		metadata = append(metadata, meta)
	}

	err = writer.Write(file, metadata)
	if err != nil {
		os.Remove(file.Name())
		fmt.Println("Error writing exif data to file: ", err.Error())
	}
}
