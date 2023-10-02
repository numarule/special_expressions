package main

import (
	"fmt"
	"regexp"
)

// Example of wanted query string - 'Player.Position.X'

// Test data to get us working with regexps
const identRe_fmt = "(?P<%s>[[:alpha:]]+[[:alnum:]]*)"
const numRe_fmt = "(?P<%s>[[:digit:]]+)"

var nameRe = fmt.Sprintf(identRe_fmt, "Name")
var xRe = fmt.Sprintf(numRe_fmt, "X")
var yRe = fmt.Sprintf(numRe_fmt, "Y")
var zRe = fmt.Sprintf(numRe_fmt, "Z")
var positionRe_fmt = fmt.Sprintf("(?P<Position>\\(\\s*%s\\s*,\\s*%s\\s*,\\s*%s\\s*\\))", xRe, yRe, zRe)
var playerRe_fmt = fmt.Sprintf("(?P<Player>Player:\\s*%s\\s*-\\s*Position:\\s*%s)", nameRe, positionRe_fmt)

type MatchMap[V string | *MatchMap] map[string]V
type MatchMapString MatchMap[string]

func CreateMatchMap(re *regexp.Regexp, matches []string) MatchMapString {
	matchMap := make(MatchMapString, re.NumSubexp())
	for i, matchName := range re.SubexpNames() {
		matchMap[matchName] = matches[i]
	}
	return matchMap
}

func main() {
	fmt.Println("Hello World")

	input := "Player: vognod - Position: (56,45,11)"
	re := regexp.MustCompile(playerRe_fmt)

	matches := re.FindStringSubmatch(input)
	if len(matches) <= 0 {
		panic(fmt.Errorf("ERROR! No matches!!!\n"))
	}

	fmt.Printf("Matches: %#v\n", matches)
	fmt.Printf("SubexpNames: %#v\n", re.SubexpNames())
	fmt.Printf("SubexpNameIndexes: %#v\n", re.NumSubexp())

	matchMap := CreateMatchMap(re, matches)

	fmt.Printf("matchMap: %#v\n", matchMap)
}
