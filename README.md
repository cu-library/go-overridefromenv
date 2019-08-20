# overridefromenv
[![Go Report Card](https://goreportcard.com/badge/cu-library/overridefromenv)](https://goreportcard.com/report/github.com/cu-library/overridefromenv)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-100%25-brightgreen.svg?longCache=true&style=flat)</a>
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/cu-library/overridefromenv?status.svg)](https://godoc.org/github.com/cu-library/overridefromenv)

A Go library for setting unset flags from environment variables.

## Usage

Here's an example of a small command line tool called 'scanner' with a flag which can be set
on the command line or from the environment. Set flags are not overwritten.

```go
package main

import (
        "flag"
        "fmt"
        "github.com/cu-library/overridefromenv"
        "os"
)

const (
        PREFIX = "SCANNER_"
)

func main() {
        v := flag.Int("powerlevel", 0, "power level")
        flag.Parse()
        err := overridefromenv.Override(flag.CommandLine, PREFIX)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Printf("Power level: %v\n", *v)
}
```

Then, from the command line:

```bash
$ scanner
Power level: 0
$ SCANNER_POWERLEVEL=1000
$ scanner
Power level: 1000
$ scanner -powerlevel 9000
Power level: 9000
$ SCANNER_POWERLEVEL="One hundred puppies."
$ scanner
Unable to set flag powerlevel from environment variable SCANNER_POWERLEVEL, which has a value of "One hundred puppies.": parse error
```
