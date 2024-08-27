package all

import (
	"fmt"

	"github.com/vdonkey/accelerator/v5/common/uuid"
	"github.com/vdonkey/accelerator/v5/main/commands/base"
)

var cmdUUID = &base.Command{
	UsageLine: "{{.Exec}} uuid",
	Short:     "generate new UUID",
	Long: `Generate new UUID.
`,
	Run: executeUUID,
}

func executeUUID(cmd *base.Command, args []string) {
	u := uuid.New()
	fmt.Println(u.String())
}
