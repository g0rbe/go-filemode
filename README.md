# go-filemode

[![Go Reference](https://pkg.go.dev/badge/github.com/g0rbe/go-filemode.svg)](https://pkg.go.dev/github.com/g0rbe/go-filemode)
[![Go Report Card](https://goreportcard.com/badge/github.com/g0rbe/go-filemode)](https://goreportcard.com/report/github.com/g0rbe/go-filemode)

Golang module to maniulate file mode bits on Linux systems.

## Install

```bash
go get github.com/g0rbe/go-filemode
```

## Usage

- `Set()`, `Unset()` and other functions are works with `Mode` type.
- `SetFile()`, `UnsetFile()` and other `...File()` functions are works with `*os.File` type.
- `SetPath()`, `UnsetPath()` and other `...Path()` functions are works with `string` that specify the file's path. **NOTE**: These functions does not follows symlinks!

## Example

```go
package main

import (
	"fmt"

	"github.com/g0rbe/go-filemode"
)

func main() {

	if err := filemode.SetPath("/path/to/file", filemode.ExecOther); err != nil {
		// Handle error
	}

	isSet, err := filemode.IsSetPath("/path/to/file", filemode.ExecOther)
	if err != nil {
		// Handle error
	}

	if !isSet {
		fmt.Printf("%s in not executable\n", "/path/to/file")
	}
}
```