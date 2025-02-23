package reader

type FileReader interface {
	Read(path string) error
}
