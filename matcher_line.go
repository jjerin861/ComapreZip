package matcher

import "fmt"

// CompareByCell
func CompareByCell(line1, line2 []string) error {
	if len(line1) != len(line2) {
		return fmt.Errorf(
			"number of columns not matching\nf1: %s\nf2: %s",
			line1,
			line2,
		)
	}
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			return fmt.Errorf(
				"column:%d not matching for line\nf1: %s\nf2: %s",
				i,
				line1,
				line2,
			)
		}
	}
	return nil
}

// CompareByLine
func CompareByLine(line1, line2 string) error {
	if len(line1) != len(line2) {
		return fmt.Errorf(
			"length of line not matching\nf1: %s\nf2: %s",
			line1,
			line2,
		)
	}

	if line1 == line2 {
		return nil
	}
	return fmt.Errorf(
		"line not matching\nf1: %s\nf2: %s",
		line1,
		line2,
	)
}
