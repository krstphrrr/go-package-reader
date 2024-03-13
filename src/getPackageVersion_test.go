package main

import (
    // "encoding/json"
    "os"
    "path/filepath"
    "testing"
)

func TestGetPackageVersion(t *testing.T) {
    // Create a temporary directory and package.json file for testing
    tempDir := t.TempDir()
    packageJSONPath := filepath.Join(tempDir, "package.json")
    version := "7.7.7"
    packageJSONContent := []byte(`{"version":"` + version + `"}`)
    if err := os.WriteFile(packageJSONPath, packageJSONContent, 0644); err != nil {
        t.Fatal(err)
    }

    // Test GetPackageVersion function
    expectedVersion := version
    actualVersion := GetPackageVersion(tempDir)
    if actualVersion != expectedVersion {
        t.Errorf("Expected version %s, got %s", expectedVersion, actualVersion)
    }
}

func TestGetPackageVersion_InvalidJSON(t *testing.T) {
    // Create a temporary directory and package.json file with invalid JSON for testing
    tempDir := t.TempDir()
    packageJSONPath := filepath.Join(tempDir, "package.json")
    packageJSONContent := []byte(`invalid JSON`)
    if err := os.WriteFile(packageJSONPath, packageJSONContent, 0644); err != nil {
        t.Fatal(err)
    }

    // Test GetPackageVersion function with invalid JSON
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected GetPackageVersion to panic with invalid JSON")
        }
    }()
    GetPackageVersion(tempDir)
}