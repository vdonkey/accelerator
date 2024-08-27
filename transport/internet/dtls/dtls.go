package dtls

import (
	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/transport/internet"
)

//go:generate go run github.com/vdonkey/accelerator/v5/common/errors/errorgen

const protocolName = "dtls"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
