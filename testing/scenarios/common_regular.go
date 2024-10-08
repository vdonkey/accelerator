//go:build !coverage
// +build !coverage

package scenarios

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func BuildAccelerator() error {
	genTestBinaryPath()
	if _, err := os.Stat(testBinaryPath); err == nil {
		return nil
	}

	fmt.Printf("Building Accelerator into path (%s)\n", testBinaryPath)
	cmd := exec.Command("go", "build", "-o="+testBinaryPath, GetSourcePath())
	return cmd.Run()
}

func RunAcceleratorProtobuf(config []byte) *exec.Cmd {
	genTestBinaryPath()
	proc := exec.Command(testBinaryPath, "run", "-format=pb")
	proc.Stdin = bytes.NewBuffer(config)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout

	return proc
}
