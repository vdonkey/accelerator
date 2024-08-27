package assembly

import (
	"context"

	"github.com/vdonkey/accelerator/v5/common"
)

//go:generate go run github.com/vdonkey/accelerator/v5/common/errors/errorgen

const protocolName = "request"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, newError("request is a transport protocol.")
	}))
}
