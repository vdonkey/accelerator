package inbound

import (
	"context"

	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/features"
)

// Handler is the interface for handlers that process inbound connections.
//
// accelerator:api:stable
type Handler interface {
	common.Runnable
	// The tag of this handler.
	Tag() string

	// Deprecated: Do not use in new code.
	GetRandomInboundProxy() (interface{}, net.Port, int)
}

// Manager is a feature that manages InboundHandlers.
//
// accelerator:api:stable
type Manager interface {
	features.Feature
	// GetHandlers returns an InboundHandler for the given tag.
	GetHandler(ctx context.Context, tag string) (Handler, error)
	// AddHandler adds the given handler into this Manager.
	AddHandler(ctx context.Context, handler Handler) error

	// RemoveHandler removes a handler from Manager.
	RemoveHandler(ctx context.Context, tag string) error
}

// ManagerType returns the type of Manager interface. Can be used for implementing common.HasType.
//
// accelerator:api:stable
func ManagerType() interface{} {
	return (*Manager)(nil)
}
