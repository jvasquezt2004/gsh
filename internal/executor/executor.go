package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteInput(tokens string, args []string) {
	if tokens == "exit" {
		os.Exit(0)
	}

	if tokens == "cd" {
		if len(args) < 1 {
			fmt.Println("Expected a directory")
			return
		}
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Printf("Error changing directory: %s\n", err)
		}
		return
	}

	out, err := exec.Command(tokens, args...).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])
	fmt.Println(output)
}
