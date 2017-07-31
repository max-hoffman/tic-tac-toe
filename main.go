package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Let's play Tic-Tac-Toe")
	fmt.Println("----------------------")
	fmt.Println("New Game!")

	for {

		var text, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("User input was invalid")
			continue
		}

		text = strings.Replace(text, "\n", " ", -1)
		fmt.Println("user input", text)
	}
}
