// Binary to actually play games using the games library.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Eventually, this could contain the whole game specification.
// For now the name will just refer to a game implementation.
type GameSpec struct {
	Name string
}

// A game is fully specified by
// - State: A struct representing the game state.
// - Moves: A function that given a game state produces all
//          possible legal moves and the corresponding game states.
// - Won: A function that given a game state determines if the game is
//        is over and if so who won.

// A player is fully specified by a function that given a game determines
// the next move to play.

func (g GameSpec) String() string {
	return g.Name
}

var gameSpecPath = flag.String("game_spec", "",
	"Path of file containing game specification in JSON encoding")

func gameSpecFromJsonFile(path string) (*GameSpec, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var gameSpec GameSpec
	json.Unmarshal(raw, &gameSpec)
	return &gameSpec, nil
}

func main() {
	flag.Parse()
	if *gameSpecPath == "" {
		fmt.Println("Must pass flag 'game_spec'")
		os.Exit(1)
	}
	gameSpec, err := gameSpecFromJsonFile(*gameSpecPath)
	if err != nil {
		fmt.Println("Could not read game specification from location " +
			"provided in 'game_spec'")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Read the following game_spec: '%s'\n", gameSpec)
}
