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

		// Prompt for fuzzy search string for folder suggestions
		fmt.Print("Enter search string for folder suggestions (or press Enter to skip): ")
		suggestionQuery, _ := reader.ReadString('\n')
		suggestionQuery = strings.TrimSpace(suggestionQuery)

		folderSuggestions := []string{}
		if suggestionQuery != "" {
			matches, _ := bState.FindMatches(suggestionQuery, 10) // Get up to 10 matches
			for _, match := range matches {
				if match.IsFolder { // Assuming File struct has an IsFolder method or field
					if len(match.Path) > 0 {
						// Display the full relative path for clarity in suggestions
						folderSuggestions = append(folderSuggestions, match.Path.String())
					}
				}
				// Limiting to a reasonable number of suggestions to display, e.g., 5
				if len(folderSuggestions) == 5 {
					break
				}
			}
		}

		if len(folderSuggestions) > 0 {
			fmt.Printf("Suggested folders based on '%s': %s\n", suggestionQuery, strings.Join(folderSuggestions, ", "))
		} else if suggestionQuery != "" {
			fmt.Printf("No folder suggestions found for '%s'.\n", suggestionQuery)
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
