package data_importers

type importer interface {
	Init(dirname string) bool
	Import() (int, error)
}
