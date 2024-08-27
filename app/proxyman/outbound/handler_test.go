package outbound_test

import (
	"context"
	"testing"
	_ "unsafe"

	"github.com/vdonkey/accelerator/v5/common/environment/systemnetworkimpl"

	"github.com/vdonkey/accelerator/v5/common/environment"
	"github.com/vdonkey/accelerator/v5/common/environment/envctx"
	"github.com/vdonkey/accelerator/v5/common/environment/transientstorageimpl"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/vdonkey/accelerator/v5"
	"github.com/vdonkey/accelerator/v5/app/policy"
	. "github.com/vdonkey/accelerator/v5/app/proxyman/outbound"
	"github.com/vdonkey/accelerator/v5/app/stats"
	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/common/serial"
	"github.com/vdonkey/accelerator/v5/features/outbound"
	"github.com/vdonkey/accelerator/v5/proxy/freedom"
	"github.com/vdonkey/accelerator/v5/transport/internet"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}

//go:linkname toContext github.com/vdonkey/accelerator/v5.toContext
func toContext(ctx context.Context, v *core.Instance) context.Context

func TestOutboundWithoutStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						InboundUplink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	defaultNetworkImpl := systemnetworkimpl.NewSystemNetworkDefault()
	rootEnv := environment.NewRootEnvImpl(ctx, transientstorageimpl.NewScopedTransientStorageImpl(), defaultNetworkImpl.Dialer(), defaultNetworkImpl.Listener())
	proxyEnvironment := rootEnv.ProxyEnvironment("o")
	ctx = envctx.ContextWithEnvironment(ctx, proxyEnvironment)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if ok {
		t.Errorf("Expected conn to not be StatCouterConnection")
	}
}

func TestOutboundWithStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						OutboundUplink:   true,
						OutboundDownlink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	defaultNetworkImpl := systemnetworkimpl.NewSystemNetworkDefault()
	rootEnv := environment.NewRootEnvImpl(ctx, transientstorageimpl.NewScopedTransientStorageImpl(), defaultNetworkImpl.Dialer(), defaultNetworkImpl.Listener())
	proxyEnvironment := rootEnv.ProxyEnvironment("o")
	ctx = envctx.ContextWithEnvironment(ctx, proxyEnvironment)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if !ok {
		t.Errorf("Expected conn to be StatCouterConnection")
	}
}
