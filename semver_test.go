package semver

import (
    "testing"
    "strings"
    "fmt"
)

func printValidateResults(testName string, input string, expected bool, actual bool) {
    fmt.Printf("Test Name: %s\n", testName)
    fmt.Printf("input: %s\n", input)
    fmt.Printf("expected : %v\n", expected)
    fmt.Printf("valid: %v\n\n", actual)
}

func printBumpResults(testName string, input string, expected string, actual string) {
    fmt.Printf("Test Name: %s\n", testName)
    fmt.Printf("input: %s\n", input)
    fmt.Printf("expected : %s\n", expected)
    fmt.Printf("actual: %s\n\n", actual)
}

func TestValidSemver(t *testing.T) {
    input := "12.1.1"
    expected := true
    valid := Validate(input)
    printValidateResults("Test Valid Semver", input, expected, valid)
    if (!valid) {
        t.Fail()
    }
}


func TestInvalidSemver(t *testing.T) {
    input := "12.1.1.1"
    expected := false
    valid := Validate(input)
    printValidateResults("Test Invalid Semver", input, expected, valid)
    if (valid) {
        t.Fail()
    }
}

func TestAnotherInvalidSemver(t *testing.T) {
    input := ".12.1.1"
    expected := false
    valid := Validate(input)
    printValidateResults("Test Invalid Semver", input, expected, valid)
    if (valid) {
        t.Fail()
    }
}

func TestBumpMajor(t *testing.T) {
    input := "12.1.1"
    expected := "13.0.0"
    actual, error := BumpMajor(input)
    printBumpResults("TestBumpMajor", input, expected, actual)
    if (error != nil || strings.Compare(expected, actual) != 0) {
        t.Fail()
    }
}

func TestBumpMinor(t *testing.T) {
    input := "12.1.1"
    expected := "12.2.0"
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
