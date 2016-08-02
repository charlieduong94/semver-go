package semver

import (
    "testing"
    "strings"
    "fmt"
)

func printValidateResults(testName string, input string, expected bool, actual error) {
    fmt.Printf("Test Name: %s\n", testName)
    fmt.Printf("input: %s\n", input)
    fmt.Printf("expected : %v\n", expected)
    fmt.Printf("valid: %v\n\n", actual == nil)
}

func printBumpResults(testName string, input string, expected string, actual string) {
    fmt.Printf("Test Name: %s\n", testName)
    fmt.Printf("input: %s\n", input)
    fmt.Printf("expected : %s\n", expected)
    fmt.Printf("valid: %s\n\n", actual)
}

func TestValidateSemver(t *testing.T) {
    input := "12.1.1"
    expected := true
    actual := Validate(input)
    printValidateResults("Test Validate Semver", input, expected, actual)
    if (actual != nil) {
        t.Fail()
    }
}


func TestValidateSemverFail(t *testing.T) {
    input := "12.1.1.1"
    expected := false
    actual := Validate(input)
    printValidateResults("Test Validate Semver", input, expected, actual)
    if (actual == nil) {
        t.Fail()
    }
}

func TestBumpMajor(t *testing.T) {
    input := "12.1.1"
    expected := "13.1.1"
    actual, error := BumpMajor(input)
    printBumpResults("TestBumpMajor", input, expected, actual)
    if (error != nil || strings.Compare(expected, actual) != 0) {
        t.Fail()
    }
}

func TestBumpMinor(t *testing.T) {
    input := "12.1.1"
    expected := "12.2.1"
    actual, error := BumpMinor(input)
    printBumpResults("TestBumpMinor", input, expected, actual)
    if (error != nil || strings.Compare(expected, actual) != 0) {
        t.Fail()
    }
}

func TestBumpPatch(t *testing.T) {
    input := "12.1.1"
    expected := "12.1.2"
    actual, error := BumpPatch(input)
    printBumpResults("TestBumpPatch", input, expected, actual)
    if (error != nil || strings.Compare(expected, actual) != 0) {
        t.Fail()
    }
}
