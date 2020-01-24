package main

import (
	"fmt"
	"os"

	"github.com/TaylorOno/go-sample/cmd/go-sample/service"
)

func main() {
	if err := service.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
