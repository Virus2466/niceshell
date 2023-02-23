package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Variable for Taking Input From User.
	reader := bufio.NewReader(os.Stdin)

	for {


		fmt.Print("$~ ")

		fmt.Printf(os.Hostname())

		// Reading Keyboard Input .
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle The Execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

// Function For Executing Input
func execInput(input string) error {
	// Remove The newline Char.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to Separate the commands and Arguments
	args := strings.Split(input, " ")

	// Check For Built In Commands.
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet Supported
		if len(args) < 2 {
			return errors.New("path required")
		}

		// Change the Dir and return the error
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)

	}

	// Command to Execute
	cmd := exec.Command(args[0], args[1:]...)

	// Setting Correct Output Device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the Command the return error
	return cmd.Run()

}
