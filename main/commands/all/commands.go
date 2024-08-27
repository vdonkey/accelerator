package all

import (
	"github.com/vdonkey/accelerator/v5/main/commands/all/api"
	"github.com/vdonkey/accelerator/v5/main/commands/all/tls"
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

//go:generate go run accelerator.com/core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
