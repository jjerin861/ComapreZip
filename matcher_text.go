package matcher

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

// NewTextMatcher
func NewTextMatcher(text1Name, text2Name string) (*Texts, error) {
	textFile1, err := os.Open(text1Name)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"can not open file:%s",
			text1Name,
		)
	}
	text := &Texts{
		text1: &Text{
			Name:   text1Name,
			Reader: textFile1,
		},
	}
	textFile2, err := os.Open(text2Name)
	if err != nil {
		return text, errors.Wrapf(
			err,
			"can not open file:%s",
			text2Name,
		)
	}
	text.text2 = &Text{
		Name:   text2Name,
		Reader: textFile2,
	}
	return text, nil
}

// Compare
func (text *Texts) Compare() (bool, error) {
	readerErr, err := Compare(text.text1.Reader, text.text2.Reader, true)
	if err == io.EOF {
		return false, nil
	}
	if err != nil {
		if readerErr {
			return true, errors.Wrapf(
				err,
				"file:%s",
				text.text2.Name,
			)
		}
		return false, errors.Wrapf(
			err,
			"file:%s",
			text.text1.Name,
		)
	}
	return false, nil
}
