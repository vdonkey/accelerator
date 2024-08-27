package core

import (
	"bytes"
	"context"

	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/common/environment/envctx"
	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/features/routing"
	"github.com/vdonkey/accelerator/v5/transport/internet/udp"
)

// CreateObject creates a new object based on the given Vdonkey instance and config. The Vdonkey instance may be nil.
func CreateObject(v *Instance, config interface{}) (interface{}, error) {
	return CreateObjectWithEnvironment(v, config, nil)
}

func CreateObjectWithEnvironment(v *Instance, config, environment interface{}) (interface{}, error) {
	var ctx context.Context
	if v != nil {
		ctx = toContext(v.ctx, v)
	}
	ctx = envctx.ContextWithEnvironment(ctx, environment)
	return common.CreateObject(ctx, config)
}

// StartInstance starts a new Vdonkey instance with given serialized config.
// By default Vdonkey only support config in protobuf format, i.e., configFormat = "protobuf". Caller need to load other packages to add JSON support.
//
// accelerator:api:stable
func StartInstance(configFormat string, configBytes []byte) (*Instance, error) {
	config, err := LoadConfig(configFormat, bytes.NewReader(configBytes))
	if err != nil {
		return nil, err
	}
	instance, err := New(config)
	if err != nil {
		return nil, err
	}
	if err := instance.Start(); err != nil {
		return nil, err
	}
	return instance, nil
}

// Dial provides an easy way for upstream caller to create net.Conn through Vdonkey.
// It dispatches the request to the given destination by the given Vdonkey instance.
// Since it is under a proxy context, the LocalAddr() and RemoteAddr() in returned net.Conn
// will not show real addresses being used for communication.
//
// accelerator:api:stable
func Dial(ctx context.Context, v *Instance, dest net.Destination) (net.Conn, error) {
	ctx = toContext(ctx, v)

	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in Vdonkey core")
	}

	r, err := dispatcher.(routing.Dispatcher).Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt net.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = net.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = net.ConnectionOutputMultiUDP(r.Reader)
	}
	return net.NewConnection(net.ConnectionInputMulti(r.Writer), readerOpt), nil
}

// DialUDP provides a way to exchange UDP packets through Vdonkey instance to remote servers.
// Since it is under a proxy context, the LocalAddr() in returned PacketConn will not show the real address.
//
// TODO: SetDeadline() / SetReadDeadline() / SetWriteDeadline() are not implemented.
//
// accelerator:api:beta
func DialUDP(ctx context.Context, v *Instance) (net.PacketConn, error) {
	ctx = toContext(ctx, v)

	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in Vdonkey core")
	}
	return udp.DialDispatcher(ctx, dispatcher.(routing.Dispatcher))
}
