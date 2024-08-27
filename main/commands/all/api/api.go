package api

import (
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

// CmdAPI calls an API in an Vdonkey process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "call Vdonkey API",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate Vdonkey via its API.
`,
	Commands: []*base.Command{
		cmdLog,
		cmdStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
	},
}
