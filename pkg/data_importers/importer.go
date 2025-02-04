package data_importers

type Importer interface {
	Init(dirname string) bool
	Import() (int, error)
}
