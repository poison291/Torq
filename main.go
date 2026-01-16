package main

import (
	"fmt"
	"os"
	"torq/core"
)

func main() {
	data, err := os.ReadFile("test.torrent")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	pos := 0
	val, err := core.Parser(data, &pos)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	info := val.(map[string]any)["info"].(map[string]any)

	fmt.Println("Torrent name:", info["name"])
	fmt.Println("Total size:", info["length"])
	fmt.Println("Piece length:", info["piece length"])

	pieces := info["pieces"].(string)
	fmt.Println("Number of pieces:", len(pieces)/20)

}
