# `json_errors` provides a simple error handling for your REST applications.  [![PkgGoDev](https://pkg.go.dev/badge/github.com/pchchv/json_errors)](https://pkg.go.dev/github.com/pchchv/json_errors) [![Go Report Card](https://goreportcard.com/badge/github.com/pchchv/json_errors)](https://goreportcard.com/report/github.com/pchchv/json_errors)

## Features:

* Compatible with built-in `error` interface
* Error wrapping
* JSON escaping

#

## Installation:

```bash
go get github.com/pchchv/json_errors
```

#

## Usage:

```go
package main

import (
	"fmt"

	"github.com/pchchv/json_errors"
)

func someFunc() error {
	return json_errors.New("nope")
}

func main() {
	if err := someFunc(); err != nil {
		wrapped := json_errors.Wrap(err, "Message about error")
		fmt.Println(wrapped.Error())
	}
}
```

```bash
go run main.go
```

### Output:

```json
{"message":"Message about error","details":{"message":"nope"}}
```

### See [examples](examples) for more.

#