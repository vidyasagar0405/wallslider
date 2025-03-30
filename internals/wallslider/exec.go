package wallslider

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func execCommand(cmd string) error {

	args := strings.Fields(cmd)
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	// Create the command. In this example, we use "echo" to print a message.
	_cmd := exec.Command(args[0], args[1:]...)

	// Create buffers to capture standard output and standard error.
	var stdout, stderr bytes.Buffer
	_cmd.Stdout = &stdout
	_cmd.Stderr = &stderr

	// Run the command.
	err := _cmd.Run()
	if err != nil {
		// Include the standard error output in the returned error message.
		return fmt.Errorf("command execution failed: %v\nstderr: %s", err, stderr.String())
	}

	return nil
}

func  executeWithPath(path string) error {
	return execCommand("nitrogen --set-zoom-fill " + path + " --save")
}
