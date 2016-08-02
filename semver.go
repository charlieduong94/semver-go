package semver

import (
    "strings"
    "regexp"
    "errors"
    "strconv"
    "fmt"
)

func buildSemver(numbers []string) (string) {
    return fmt.Sprintf("%s.%s.%s", numbers[0], numbers[1], numbers[2])
}

func incrementVersion(semver string, pos int) (string, error){
    numberRegex, _ := regexp.Compile("[1-9]*")
    err := Validate(semver)
    if (err != nil) {
        return "", err
    }
    matches := numberRegex.FindAllString(semver, -1)
    num, _ := strconv.Atoi(matches[pos])
    num++;
    matches[pos] = strconv.Itoa(num)

    res := buildSemver(matches)
    return res, nil
}

func Validate(semver string) (error) {
    regex, _ := regexp.Compile("^[1-9]*\\.[1-9]*\\.[1-9]*$")
    semver = strings.TrimSpace(semver)
    match := regex.FindString(semver)
    if (len(match) == 0) {
        return errors.New("Invalid semver: " + semver);
    } else {
        return nil
    }
}

func BumpMajor(semver string) (string, error) {
    return incrementVersion(semver, 0)
}

func BumpMinor(semver string) (string, error) {
    return incrementVersion(semver, 1)
}

func BumpPatch(semver string) (string, error) {
    return incrementVersion(semver, 2)
}
