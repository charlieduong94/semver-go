# semver-go

A simple tool for validating and incrementing semantic version strings.

# Usage
#### Semver validation
```
valid := semver.Validate("12.1.1")

if (valid) {
    fmt.Printf("semver is valid")
} else {
    fmt.Printf("semver is invalid")
}
```
#### Bumping versions
```
vA, _ := semver.BumpMajor("12.1.1")
// vA === "13.0.0"
vB, _ := semver.BumpMinor("12.1.1")
// vB === "12.2.0"
vC, _ := semver.BumpPatch("12.1.1")
// vC === "12.1.2"
```
