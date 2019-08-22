package matcher

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func Compare(reader1, reader2 io.Reader, byLine bool) (bool, error) {
	br1 := bufio.NewReader(reader1)
	br2 := bufio.NewReader(reader2)
	r := 1
	for {
		if byLine {
			lineString1, err := br1.ReadString('\n')
			if err == io.EOF {
				lineString2, err := br2.ReadString('\n')
				if err == io.EOF {
					return false, err
				}
				if err != nil {
					return true, errors.Wrapf(
						err,
						"data in line:%d error while reading",
						r,
					)
				}
				return true, fmt.Errorf(
					"data in line:%d not matching for\nf1: EOF\n f2:%s",
					r,
					lineString2,
				)
			}
			if err != nil {
				return false, errors.Wrapf(
					err,
					"data in line:%d error while reading",
					r,
				)
			}
			lineString2, err := br2.ReadString('\n')
			if err == io.EOF {
				return true, fmt.Errorf(
					"data in line:%d not matching for\nf1: %s\n f2:EOF",
					r,
					lineString1,
				)
			}
			if err != nil {
				return true, errors.Wrapf(
					err,
					"data in line:%d error while reading",
					r,
				)
			}
			err = CompareByLine(lineString1, lineString2)
			if err != nil {
				return true, errors.Wrapf(
					err,
					" data in line:%d",
					r,
				)
			}
		} else {
			lineByte1, err := br1.ReadBytes('\n')
			if err == io.EOF {
				lineByte2, err := br2.ReadBytes('\n')
				if err == io.EOF {
					return false, err
				}
				if err != nil {
					return true, errors.Wrapf(
						err,
						"data in row:%d error while reading",
						r,
					)
				}
				return true, fmt.Errorf(
					"data in row:%d not matching for\nf1: EOF\n f2:%s",
					r,
					lineByte2,
				)
			}
			if err != nil {
				return false, errors.Wrapf(
					err,
					"data in row:%d error while reading",
					r,
				)
			}
			lineByte2, err := br2.ReadBytes('\n')
			if err == io.EOF {
				return true, fmt.Errorf(
					"data in line:%d not matching for\nf1: %s\n f2:EOF",
					r,
					lineByte1,
				)
			}
			if err != nil {
				return true, errors.Wrapf(
					err,
					"data in line:%d error while reading",
					r,
				)
			}
			lineString1 := string(lineByte1)
			lineString2 := string(lineByte2)
			err = CompareByCell(
				strings.Split(lineString1, ","),
				strings.Split(lineString2, ","),
			)
			if err != nil {
				return true, errors.Wrapf(
					err,
					" data in row:%d",
					r,
				)
			}
		}

	}
}
