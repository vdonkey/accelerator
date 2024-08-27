//go:build coverage
// +build coverage

package scenarios

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/vdonkey/accelerator/v5/common/uuid"
)

func BuildAccelerator() error {
	genTestBinaryPath()
	if _, err := os.Stat(testBinaryPath); err == nil {
		return nil
	}

	cmd := exec.Command("go", "test", "-tags", "coverage coveragemain", "-coverpkg", "github.com/vdonkey/accelerator/v5/...", "-c", "-o", testBinaryPath, GetSourcePath())
	return cmd.Run()
}

func RunAcceleratorProtobuf(config []byte) *exec.Cmd {
	genTestBinaryPath()

	covDir := os.Getenv("ACCELERATOR_COV")
	os.MkdirAll(covDir, os.ModeDir)
	randomID := uuid.New()
	profile := randomID.String() + ".out"
	proc := exec.Command(testBinaryPath, "run", "-format=pb", "-test.run", "TestRunMainForCoverage", "-test.coverprofile", profile, "-test.outputdir", covDir)
	proc.Stdin = bytes.NewBuffer(config)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout

	return proc
}
