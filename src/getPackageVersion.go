package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// FindPackageJSON finds package.json with path.
func FindPackageJSON(path string) string {
	content, err := os.ReadFile(filepath.Join(path, "package.json"))
	if err != nil {
		panic(err)
	}
	return string(content)
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
