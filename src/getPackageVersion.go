package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"fmt"
	"log"
)

// FindPackageJSON finds package.json with path.
func FindPackageJSON(path string) string {
	fmt.Println("working directory: ",printWD())

	fmt.Println("files in working supplied dir (not wd): ", listDir(path))

	content, err := os.ReadFile(filepath.Join(path, "package.json"))
	if err != nil {
		panic(err)
	}
	return string(content)
}

func listDir(path string) []string {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
		files = append(files, e.Name())
	}
	return files
}

func printWD() string{
	mydir, err := os.Getwd()
	if err != nil { 
		fmt.Println(err) 
	} 
	return mydir 
}

// GetPackageVersion gets the version field within package.json.
func GetPackageVersion(path string) string {
	packageJSON := FindPackageJSON(path)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(packageJSON), &data); err != nil {
		panic(err) 
	}
	version, ok := data["version"].(string)
	if !ok {
		panic("Version not found in package.json") 
	}
	return version
}
