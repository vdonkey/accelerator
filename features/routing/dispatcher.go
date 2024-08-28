package routing

import (
	"context"

	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/features"
	"github.com/vdonkey/accelerator/v5/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a Accelerator instance to make Accelerator function properly.
//
// accelerator:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// accelerator:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
