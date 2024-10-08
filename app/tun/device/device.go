package device

import (
	"gvisor.dev/gvisor/pkg/tcpip/stack"

	"github.com/vdonkey/accelerator/v5/common"
)

//go:generate go run github.com/vdonkey/accelerator/v5/common/errors/errorgen

type Device interface {
	stack.LinkEndpoint

	common.Closable
}

type Options struct {
	Name string
	MTU  uint32
}

type DeviceConstructor func(Options) (Device, error)
