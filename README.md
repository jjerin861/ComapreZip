# matcher
Got Archives? matcher's compare feature allows you to compare the contents of 
zip,csv,txt files

## Basic Usage
Here is a minimal example usage that will compare files given as arguments:

```go

package main

import (
    "fmt"
    "log"
    "github.com/jjerin861/matcher"
)

func main() {
	err := matcher.Match(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"%s and %s are matching/n",
		os.Args[1],
		os.Args[2],
	)

}

```
