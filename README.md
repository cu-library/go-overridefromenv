# overridefromenv
[![Go Report Card](https://goreportcard.com/badge/github.com/benjohns1/scheduled-tasks/services)](https://goreportcard.com/report/github.com/benjohns1/scheduled-tasks/services)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-100%25-brightgreen.svg?longCache=true&style=flat)</a>
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/culibrary/overridefromenv?status.svg)](https://godoc.org/github.com/culibrary/overridefromenv)

A tiny golang library. If any flags are not set, use environment variables to set them.

## Usage

Here's an example of a small command line tool with a flag which can be set
on the command line or from the environment. Set flags are not overwritten.

```go
package main

import (
    "flag"
    "fmt"
    "https://github.com/cu-library/overridefromenv"
)

const (
    PREFIX = "SCANNER_"
)

func main() {
    v := flag.Int("powerlevel", 0, "power level")
    flag.Parse()
    overridefromenv.Override(flags.CommandLine, PREFIX)
    fmt.Printf("Power level: %v", *v)
}
```

Then, from the command line:

```bash
$ scanner
Power level: 0
$ export SCANNER_POWERLEVEL=1000
$ scanner
Power level: 1000
$ scanner -powerlevel 9000
Power level: 9000
```
