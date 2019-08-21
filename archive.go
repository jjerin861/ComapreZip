package matcher

import "archive/zip"

// Zip represents zip.
type Zip struct {
	Matcher    FileType
	Name       string
	ReadCloser *zip.ReadCloser
	Files      map[string]*zip.File
}
