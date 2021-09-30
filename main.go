package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main() {
	fmt.Println(errors.New("new error created by "), os.Args[1])
}
