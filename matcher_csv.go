package matcher

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

// NewCSVMatcher
func NewCSVMatcher(csv1Name, csv2Name string) (*CSVs, error) {
	csvFile1, err := os.Open(csv1Name)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"can not open file:%s",
			csv1Name,
		)
	}
	csv := &CSVs{
		csv1: &CSV{
			Name:   csv1Name,
			Reader: csvFile1,
		},
	}
	csvFile2, err := os.Open(csv2Name)
	if err != nil {
		return csv, errors.Wrapf(
			err,
			"can not open file:%s",
			csv2Name,
		)
	}
	csv.csv2 = &CSV{
		Name:   csv2Name,
		Reader: csvFile2,
	}
	return csv, nil
}

// Compare
func (csv *CSVs) Compare() (bool, error) {
	readerErr, err := Compare(csv.csv1.Reader, csv.csv2.Reader, false)
	if err == io.EOF {
		return false, nil
	}
	if err != nil {
		if readerErr {
			return true, errors.Wrapf(
				err,
				"file:%s",
				csv.csv2.Name,
			)
		}
		return false, errors.Wrapf(
			err,
			"file:%s",
			csv.csv1.Name,
		)
	}
	return false, nil
}
