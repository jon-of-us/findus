package main

import (
	"bufio"
	"findus/backend"
	"fmt"
	"os"
	"strings"
)

func main() {
	bState := backend.InitState()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Printf("Current path: %s\n", bState.Path.String())
		fmt.Printf("Forward path: %s\n", bState.ForwardPath.String())

		// Use FindMatches to get folder suggestions
		matches, _ := bState.FindMatches("", 100) // Search for empty string, get up to 10 matches
		folderSuggestions := []string{}
		for _, match := range matches {
			if match.IsFolder { // Assuming File struct has an IsFolder method or field
				// We need to get the base name of the folder, not the full relative path from FindMatches
				if len(match.Path) > 0 {
					folderSuggestions = append(folderSuggestions, match.Path[len(match.Path)-1])
				}
			}
			if len(folderSuggestions) == 3 {
				break
			}
		}

		if len(folderSuggestions) > 0 {
			fmt.Printf("Suggested subfolders: %s\n", strings.Join(folderSuggestions, ", "))
		} else {
			fmt.Println("No subfolder suggestions found.")
		}

		fmt.Print("Enter folder name to navigate, 'exit' to go up, or press Enter to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Exiting interactive navigation.")
			break
		}

		if strings.ToLower(input) == "exit" {
			bState.PopPath()
		} else {
			bState.AddToPath(backend.Path{input})
		}
	}

	bState.QuitAndSetPath()
}
