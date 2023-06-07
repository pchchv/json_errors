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
