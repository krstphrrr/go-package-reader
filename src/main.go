package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]
	// result := GetPackageVersion(path)
	fmt.Println(GetPackageVersion(path))
}
