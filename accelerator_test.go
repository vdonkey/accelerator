package core_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	. "github.com/vdonkey/accelerator/v5"
	"github.com/vdonkey/accelerator/v5/app/dispatcher"
	"github.com/vdonkey/accelerator/v5/app/proxyman"
	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/common/protocol"
	"github.com/vdonkey/accelerator/v5/common/serial"
	"github.com/vdonkey/accelerator/v5/common/uuid"
	"github.com/vdonkey/accelerator/v5/features/dns"
	"github.com/vdonkey/accelerator/v5/features/dns/localdns"
	_ "github.com/vdonkey/accelerator/v5/main/distro/all"
	"github.com/vdonkey/accelerator/v5/proxy/dokodemo"
	"github.com/vdonkey/accelerator/v5/proxy/vmess"
	"github.com/vdonkey/accelerator/v5/proxy/vmess/outbound"
	"github.com/vdonkey/accelerator/v5/testing/servers/tcp"
)

func TestAcceleratorDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestAcceleratorClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance(FormatProtobuf, cfgBytes)
	common.Must(err)
	server.Close()
}
