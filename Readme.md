# Project Title

Go Cli Tool for reading EXIF data from the images.

## Description

This project is a command-line utility written in Go that scans a specified directory (and its sub-directories) for image files, extracts their EXIF metadata, and writes the following attributes to a CSV or HTML file:

- Image file path
- GPS position latitude
- GPS position longitude


### Dependencies

- [github.com/rwcarlsen/goexif](https://github.com/rwcarlsen/goexif)
- Go v1.20

### Installing

* go mod tidy

### Executing program
```
go run main.go -dir=../images -output=output.csv
go run main.go -dir=../images -output=output.html
```
