package matcher

import (
	"archive/zip"
)

type ArchiveType interface {
	Compare() (bool, error)
}
type Zips struct {
	zip1, zip2 *Zip
	Matcher    FileType
}

// Zip represents zip.
type Zip struct {
	Name       string
	ReadCloser *zip.ReadCloser
	Files      map[string]*zip.File
}
