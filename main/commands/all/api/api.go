package api

import (
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

// CmdAPI calls an API in an Accelerator process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "call Accelerator API",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate Accelerator via its API.
`,
	Commands: []*base.Command{
		cmdLog,
		cmdStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
	},
}
