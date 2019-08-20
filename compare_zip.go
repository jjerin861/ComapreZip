package matcher

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Compare compares two zips and returns the first error or diff in the
// zips. If two of them are matching then nil will be returned.
func Compare(zip1Name, zip2Name string) error {
	// Open zip archive for reading.
	zip1, err := zip.OpenReader(zip1Name)
	if err != nil {
		return errors.Wrapf(
			err,
			"can not open:%s",
			zip1Name,
		)
	}
	defer zip1.Close()

	zip2, err := zip.OpenReader(zip2Name)
	if err != nil {
		return errors.Wrapf(
			err,
			"can not open:%s",
			zip2Name,
		)
	}
	defer zip2.Close()

	// Iterate through the files in the archive, returning diff of their
	// contents.
	for _, f1 := range zip1.File {
		rc1, err := f1.Open()
		if err != nil {
			return errors.Wrapf(
				err,
				"zip name:%s can not open file:%s",
				zip1Name,
				f1.Name,
			)
		}

		defer rc1.Close()

		csvReader1 := csv.NewReader(rc1)
		csvReader1.ReuseRecord = true

		var csvReader2 *csv.Reader
		for _, f2 := range zip2.File {
			if f2.Name == f1.Name {
				rc2, err := f2.Open()
				if err != nil {
					return errors.Wrapf(
						err,
						"zip name:%s can not open file:%s",
						zip2Name,
						f2.Name,
					)
				}

				defer rc2.Close()

				csvReader2 = csv.NewReader(rc2)
				csvReader2.ReuseRecord = true
			}
		}
		if csvReader2 == nil {
			return fmt.Errorf(
				"file not found in second zip:%s",
				f1.Name,
			)
		}
		for {
			r := 1
			line1, err := csvReader1.Read()
			if err == io.EOF {
				line2, err := csvReader2.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return errors.Wrapf(
						err,
						"for zip:%s file:%s data in row:%d error while reading",
						zip2Name,
						f1.Name,
						r,
					)
				}
				return fmt.Errorf(
					"for file:%s data in row:%d not matching for\nf1: EOF \nf2: %s",
					f1.Name,
					r,
					line2,
				)
			}
			if err != nil {
				return errors.Wrapf(
					err,
					"for zip:%s file:%s data in row:%d error while reading",
					zip1Name,
					f1.Name,
					r,
				)
			}
			line2, err := csvReader2.Read()
			if err == io.EOF {
				return fmt.Errorf(
					"for file:%s data in row:%d not matching for\nf1: %s\n f2:EOF",
					f1.Name,
					r,
					line1,
				)
			}
			if err != nil {
				return errors.Wrapf(
					err,
					"for zip:%s file:%s data in row:%d error while reading",
					zip2Name,
					f1.Name,
					r,
				)
			}
			if len(line1) != len(line2) {
				return fmt.Errorf(
					"for file:%s data in row:%d number of columns not matching\nf1: %s\nf2: %s",
					f1.Name,
					r,
					line1,
					line2,
				)
			}
			for i := 0; i < len(line1); i++ {
				if line1[i] != line2[i] {
					return fmt.Errorf(
						"for file:%s data in row:%d column:%d not matching for line\nf1: %s\nf2: %s",
						f1.Name,
						r,
						i,
						line1,
						line2,
					)
				}
			}
		}
	}
	return nil
}
