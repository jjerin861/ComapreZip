# ComapreZip
Got Archives? ComapreZip's archive compare feature allows you to compare the contents of .zip files

** Basic Usage
Here is a minimal example usage that will compare zip files given as arguments:

#+BEGIN_SRC go

package main

import (
    "fmt"
    "log"
    "github.com/jjerin861/ComapreZip"
)

func main() {
	err := compareZip.Compare(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"%s and %s are matching/n",
		os.Args[1],
		os.Args[2],
	)

}

#+END_SRC