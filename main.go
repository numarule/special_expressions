package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Test data to get us working with regexps
const identRe_fmt = "(?P<%s>[[:alpha:]]+[[:alnum:]]*)"
const numRe_fmt = "(?P<%s>[[:digit:]]+)"

var nameRe = fmt.Sprintf(identRe_fmt, "Name")
var xRe = fmt.Sprintf(numRe_fmt, "X")
var yRe = fmt.Sprintf(numRe_fmt, "Y")
var zRe = fmt.Sprintf(numRe_fmt, "Z")
var positionRe_fmt = fmt.Sprintf("(?P<Position>\\(\\s*%s\\s*,\\s*%s\\s*,\\s*%s\\s*\\))", xRe, yRe, zRe)
var playerRe_fmt = fmt.Sprintf("(?P<Player>Player:\\s*%s\\s*-\\s*Position:\\s*%s)", nameRe, positionRe_fmt)

type MatchMap map[string]string


const ADDRESS_SEPARATOR = "."

// Returns value for matched capture @ `address`
// input := "Player: vognod - Position: (56,45,11)"
// re := regexp.MustCompile(?P<Player>Player:\\s*(?P<Name>[[:alpha:]]+[[:alnum:]]*)\\s*-\\s*Position:\\s*(?P<Position>\\(\\s*(?P<X>[[:digit:]]+)\\s*,\\s*(?P<Y>[[:digit:]]+)\\s*,\\s*(?P<Z>[[:digit:]]+)\\s*\\)))
//
// address := "Player.Position.X"
// fmt.Println(address) // "56"
// Requires:
// FindStringSubmatch - to parse out matches
// FindStringSubmatchIndex - to parse out match ranges (start = 2n, end = 2n+2)
// SubexpNames - Names of named expressions, in order of appearance in regex
func GetValue(reStr string, input string, address string) (string, error) {
	re := regexp.MustCompile(reStr)

	//TODO: Opt, store matches and indicies in object to prevent rematching
	matches := re.FindStringSubmatch(input)
	if len(matches) <= 0 {
		return "", fmt.Errorf("ERROR! No matches!\n")
	}

	indicies := re.FindStringSubmatchIndex(input)
	if len(indicies) <= 0 {
		return "", fmt.Errorf("ERROR! No indicies!\n")
	}

	captureNames := re.SubexpNames()
	if len(captureNames) <= 0 {
		return "", fmt.Errorf("ERROR! No named captures!\n")
	}

	//BUG Algorithm as is isn't working, and would match improper nesting
	//  NEED to use stack of capture indicies for determining current path
	//    Indicies are unique, unlike names
	//    Reconstruct current address with the stack portions indexed into captureNames
	currentAddressStack := make([]int, 0)
	for i := range captureNames {
		maybeAddress :=AddressStack_ToString(captureNames, append(currentAddressStack, i))
		if !strings.HasPrefix(address, maybeAddress) {
			//Not a match
			continue
		}
		currentAddressStack = append(currentAddressStack, i)
	}

	if len(currentAddressStack) <= 0 {
		return "", fmt.Errorf("ERROR! Address not found: '%s'\n", address)
	}

	currentIndex := currentAddressStack[len(currentAddressStack)-1]
	match := matches[currentIndex]

	return match, nil
}

func AddressStack_ToString(captureNames []string, stack []int) (r string) {
	for _, addressSegment := range stack {
		r += captureNames[addressSegment] + ADDRESS_SEPARATOR
	}
	r = strings.TrimSuffix(r, ADDRESS_SEPARATOR)
	return
}


func CreateMatchMap(re *regexp.Regexp, matches []string) MatchMap {
	matchMap := make(MatchMap, re.NumSubexp())
	for i, matchName := range re.SubexpNames() {
		matchMap[matchName] = matches[i]
	}
	return matchMap
}

func main() {
	fmt.Println("Hello World")

	input := "Player: vognod - Position: (56,45,11)"
	reStr := playerRe_fmt

	playerY, err := GetValue(reStr, input, "Player.Position.Y")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Player.Position.Y: '%s'\n", playerY)
}
