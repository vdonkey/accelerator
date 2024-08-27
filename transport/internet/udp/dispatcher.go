package udp

import (
	"context"
	"io"

	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/common/buf"
	"github.com/vdonkey/accelerator/v5/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}

var DispatcherConnectionTerminationSignalReceiverMark = "DispatcherConnectionTerminationSignalReceiverMark"

type DispatcherConnectionTerminationSignalReceiver interface {
	io.Closer
}
