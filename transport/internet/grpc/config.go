package grpc

import (
	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/transport/internet"
)

const protocolName = "gun"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
