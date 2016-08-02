# semver-go

A simple tool for validating and incrementing semantic version strings.

# Usage
#### Semver validation
```
err := semver.Validate("12.1.1")

if (err == nil) {
    fmt.Printf("semver is valid")
} else {
    fmt.Printf("semver is invalid")
}
```
#### Bumping versions
```
vA, _ := semver.BumpMajor("12.1.1")
// vA === "13.1.1"
vB, _ := semver.BumpMinor("12.1.1")
// vB === "12.2.1"
vC, _ := semver.BumpPatch("12.1.1")
// vC === "12.1.2"
```
