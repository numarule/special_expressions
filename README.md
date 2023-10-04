# Special Expressions

Addressed named captures in regexp.

Separator

## Simple flat example:
```go
	input := "Player: vognod - Position: (56,45,11)"
  // Simplest 1d version is just a `map[captureName]Value`
	matchMap := CreateMatchMap(re, matches)
  //...
	fmt.Printf("matchMap[\"Player\"]: %v\n", matchMap["Player"])
	fmt.Printf("matchMap[\"Position\"]: %v\n", matchMap["Position"])
  //...
```
## Nested example:
```go
	input := "Player: vognod - Position: (56,45,11)"
  // Simplest 1d version is just a `map[captureName]Value`
	matchMap := CreateMatchMap(re, matches)
  //...
	fmt.Printf("matchMap[\"Player\"]: %v\n", matchMap["Player"])
	fmt.Printf("matchMap[\"Position\"]: %v\n", matchMap["Position"])
  //...
```

Primarily this is a small scope project while I get streaming setup, tested and operating well.

### Feature Ideas
- Builder pattern for constructing regular expressions
- Structured Captures
  - [X] Start with function that returns named captures as `map[string]string`
    - (?<Player>[[:alnum:]]+: )
  - Indexed captures
  - Add sub-captures
    - Query sytnax similar to js objects 'Player.Position.X'
  - Output should support different dialects
