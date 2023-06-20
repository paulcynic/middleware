// $GOPATH/src/github.com/paulcynic/middleware
// golang-standards/project-layout

package main

import (
    "log"

    "github.com/paulcynic/middleware/internal/pkg/app"
)

func main() {
    a, err := app.New()
    if err != nil {
        log.Fatal(err)
    }

    err = a.Run()
    if err != nil {
        log.Fatal(err)
    }
}

