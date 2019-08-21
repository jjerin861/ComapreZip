package matcher

import (
	"io"
	"regexp"
)

var fileTypeRegex, _ = regexp.Compile(`^.*\.(csv|txt)$`)

type FileType interface {
	Compare() (bool, error)
}
type CSVs struct {
	csv1, csv2 *CSV
}
type CSV struct {
	Name   string
	Reader io.Reader
}
type Texts struct {
	text1, text2 *Text
}
type Text struct {
	Name   string
	Reader io.Reader
}
