package v4_test

import (
	"testing"

	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/common/protocol"
	"github.com/vdonkey/accelerator/v5/common/serial"
	"github.com/vdonkey/accelerator/v5/infra/conf/cfgcommon"
	"github.com/vdonkey/accelerator/v5/infra/conf/cfgcommon/testassist"
	v4 "github.com/vdonkey/accelerator/v5/infra/conf/v4"
	"github.com/vdonkey/accelerator/v5/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.ShadowsocksServerConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "accelerator-password"
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "accelerator-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
