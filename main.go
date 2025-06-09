package main

import (
	"bufio"
	"fmt"
	"os"

	"findus/gui"
	"findus/gui/ansi"
)

func main() {
	restoreTerminal, err := gui.SetupTerminal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up terminal: %v\n", err)
		os.Exit(1)
	}
	defer restoreTerminal()

	// Create components
	headingComponent := gui.NewComponent("My Heading").Bold()
	bodyTextComponent := gui.NewComponent("Some text below.").Styled(ansi.Blue)
	chainedComponent := gui.NewComponent("Bold and Gray").Bold().Styled(ansi.Gray)
	// Prepare content to render
	output := fmt.Sprintf("%s\n%s\n%s\nPress any key to exit...",
		headingComponent.Content,
		bodyTextComponent.Content,
		chainedComponent.Content,
	)

	// Render the content (Render already clears and sets cursor)
	gui.Render(output)

	// Wait for user input to exit
	reader := bufio.NewReader(os.Stdin)
	_, _, err = reader.ReadRune()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
