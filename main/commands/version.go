package commands

import (
	"fmt"

	core "github.com/vdonkey/accelerator/v5"
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

// CmdVersion prints Accelerator Versions
var CmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "print Accelerator version",
	Long: `Prints the build information for Accelerator.
`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
