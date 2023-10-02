package main

import(
	"fmt"
	"regexp"
)

//Test data to get us working with regexps
const nameRe = "(?P<Name>[[:alnum:]]+)"
const positionRe = "(?P<Position>\\(\\d+,\\d+,\\d+\\))"
var playerRe = fmt.Sprintf("(?P<Player>Player:\\s*%s\\s*-\\s*Position:\\s*%s)", nameRe, positionRe)


func main() {
	fmt.Println("Hello World")

	input := "Player: vognod - Position: (56,45,11)"
	re_str := fmt.Sprintf(
		"%s",
		playerRe,
	)
	re := regexp.MustCompile(re_str)

	matches := re.FindStringSubmatch(input)
	fmt.Printf("Matches: %#v\n", matches)
	fmt.Printf("SubexpNames: %#v\n", re.SubexpNames())
	fmt.Printf("SubexpNameIndexes: %#v\n", re.NumSubexp())

	matchMap := CreateMatchMap(re, matches)

	fmt.Printf("matchMap: %#v\n", matchMap)
}

func CreateMatchMap(re *regexp.Regexp, matches []string) map[string]string {
	matchMap := make(map[string]string, re.NumSubexp())
	for i, matchName := range re.SubexpNames() {
		matchMap[matchName] = matches[i]
	}
	return matchMap
}
