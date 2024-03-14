package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	// fmt.Printf("Load package.json at %s\n", path)

	result := GetPackageVersion(path)

	// fmt.Printf("::set-output name=version::%s", result)
	fmt.Printf("%s", result)
}
