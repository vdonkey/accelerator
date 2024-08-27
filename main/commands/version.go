package commands

import (
	"fmt"

	core "github.com/vdonkey/accelerator/v5"
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

// CmdVersion prints Vdonkey Versions
var CmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "print Vdonkey version",
	Long: `Prints the build information for Vdonkey.
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
