package matcher

import (
	"archive/zip"
	"fmt"

	"github.com/pkg/errors"
)

// NewZipMatcher
func NewZipMatcher(zip1Name, zip2Name string) (*Zips, error) {
	zip1Reader, err := zip.OpenReader(zip1Name)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"can not open:%s",
			zip1Name,
		)
	}
	zips := &Zips{
		zip1: &Zip{
			Name:       zip1Name,
			ReadCloser: zip1Reader,
			Files:      map[string]*zip.File{},
		},
	}
	for _, f := range zip1Reader.File {
		zips.zip1.Files[f.Name] = f
	}
	zip2Reader, err := zip.OpenReader(zip2Name)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"can not open:%s",
			zip2Name,
		)
	}

	zips.zip2 = &Zip{
		Name:       zip2Name,
		ReadCloser: zip2Reader,
		Files:      map[string]*zip.File{},
	}
	for _, f := range zip2Reader.File {
		zips.zip2.Files[f.Name] = f
	}
	return zips, nil
}

// Compare compares two zips and returns the first error or diff in the
// zips. If two of them are matching then nil will be returned.
func (zips *Zips) Compare() error {

	// Iterate through the files in the archive, returning diff of their
	// contents.
	for f1Name, f1 := range zips.zip1.Files {
		if zips.zip2.Files[f1Name] == nil {
			return fmt.Errorf(
				"file not found in zip:%s file:%s",
				zips.zip2.Name,
				f1Name,
			)
		}
		f2 := zips.zip2.Files[f1Name]
		matched := fileTypeRegex.FindStringSubmatch(f1Name)
		if len(matched) != 2 {
			return fmt.Errorf(
				"can not find file format zip:%s file:%s",
				zips.zip1.Name,
				f1Name,
			)
		}
		r1, err := f1.Open()
		if err != nil {
			return errors.Wrapf(
				err,
				"can not open file zip:%s file:%s",
				zips.zip1.Name,
				f1Name,
			)
		}
		r2, err := f2.Open()
		if err != nil {
			return errors.Wrapf(
				err,
				"can not open file zip:%s file:%s",
				zips.zip2.Name,
				f1Name,
			)
		}
		if matched[1] == "csv" {
			csv := &CSVs{
				csv1: &CSV{
					Name:   f1Name,
					Reader: r1,
				},
				csv2: &CSV{
					Name:   f1Name,
					Reader: r2,
				},
			}

			zips.Matcher = csv
		} else if matched[1] == "txt" {

			text := &Texts{
				text1: &Text{
					Name:   f1Name,
					Reader: r1,
				},
				text2: &Text{
					Name:   f1Name,
					Reader: r2,
				},
			}

			zips.Matcher = text
		}
		readerErr, err := zips.Matcher.Compare()
		if err != nil {
			if readerErr {
				return errors.Wrapf(
					err,
					"zip:%s",
					zips.zip2.Name,
				)
			}
			return errors.Wrapf(
				err,
				"zip:%s",
				zips.zip1.Name,
			)
		}

	}
	return nil
}
