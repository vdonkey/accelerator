//go:build !linux || (linux && !(amd64 || arm64))
// +build !linux linux,!amd64,!arm64

package gvisor

import "github.com/vdonkey/accelerator/v5/app/tun/device"

func New(options device.Options) (device.Device, error) {
	return nil, newError("not supported").AtError()
}
