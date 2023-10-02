package main

import(
	"fmt"
	"regexp"
)

//Test data to get us working with regexps
const nameRe = "(?P<Name>[[:alnum:]]+)"
const numRe_fmt = "(?P<%s>[[:digit:]]+)"
var xRe = fmt.Sprintf(numRe_fmt, "X")
var yRe = fmt.Sprintf(numRe_fmt, "Y")
var zRe = fmt.Sprintf(numRe_fmt, "Z")
var positionRe = fmt.Sprintf("(?P<Position>\\(\\s*%s\\s*,\\s*%s\\s*,\\s*%s\\s*\\))", xRe, yRe, zRe)
var playerRe = fmt.Sprintf("(?P<Player>Player:\\s*%s\\s*-\\s*Position:\\s*%s)", nameRe, positionRe)

// Example Query string - 'Player.Position.X'

type MatchMap map[string]string

func main() {
	fmt.Println("Hello World")

	input := "Player: vognod - Position: (56,45,11)"
	re := regexp.MustCompile(playerRe)

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

func CreateMatchMap(re *regexp.Regexp, matches []string) MatchMap {
	matchMap := make(map[string]string, re.NumSubexp())
	for i, matchName := range re.SubexpNames() {
		matchMap[matchName] = matches[i]
	}
	return matchMap
}
