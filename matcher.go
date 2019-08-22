package matcher

import (
	"fmt"
	"regexp"
)

var typeRegex, _ = regexp.Compile(`^.*\.(zip|csv|txt)$`)

// Match comapres two files and returns nil if there is no mismatch.
func Match(fileName1, fileName2 string) error {
	matched1 := typeRegex.FindStringSubmatch(fileName1)
	if len(matched1) != 2 {
		return fmt.Errorf(
			"can not find file format for file:%s",
			fileName1,
		)
	}
	matched2 := typeRegex.FindStringSubmatch(fileName2)
	if len(matched2) != 2 {
		return fmt.Errorf(
			"can not find file format for file:%s",
			fileName2,
		)
	}
	if matched1[1] != matched2[1] {
		return fmt.Errorf(
			"file formats not matching f1:%s f2:%s",
			fileName1,
			fileName2,
		)
	}
	switch matched1[1] {
	case "zip":
		zip, err := NewZipMatcher(fileName1, fileName2)
		if err != nil {
			return err
		}
		return zip.Compare()
	case "csv":
		csv, err := NewCSVMatcher(fileName1, fileName2)
		if err != nil {
			return err
		}
		_, err = csv.Compare()
		return err
	case "txt":
		text, err := NewTextMatcher(fileName1, fileName2)
		if err != nil {
			return err
		}
		_, err = text.Compare()
		return err
	default:
		return fmt.Errorf(
			"unsupported file format f1:%s f2%s",
			fileName1,
			fileName2,
		)
	}
}
