package semver

import (
    "strings"
    "regexp"
    "errors"
    "strconv"
    "fmt"
)
const finalSemverPos = 2
func buildSemver(numbers []string) (string) {
    return fmt.Sprintf("%s.%s.%s", numbers[0], numbers[1], numbers[2])
}

func incrementVersion(semver string, pos int) (string, error){
    numberRegex, _ := regexp.Compile("[0-9]*")
    valid := Validate(semver)
    if (!valid) {
        return "", errors.New("Invalid semver: " + semver);
    }
    matches := numberRegex.FindAllString(semver, -1)
    num, _ := strconv.Atoi(matches[pos])
    num++;
    matches[pos] = strconv.Itoa(num)
    pos++;
    for pos <= finalSemverPos {
        matches[pos] = "0"
        pos++;
     }

    res := buildSemver(matches)
    return res, nil
}

func Validate(semver string) (bool) {
    regex, _ := regexp.Compile("^[0-9]*\\.[0-9]*\\.[0-9]*$")
    semver = strings.TrimSpace(semver)
    match := regex.FindString(semver)
    return strings.Compare(match, semver) == 0
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
