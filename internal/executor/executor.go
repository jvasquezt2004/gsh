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

	out, err := exec.Command(tokens, args...).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])
	fmt.Println(output)
}
