package dns

import (
	"github.com/vdonkey/accelerator/v5/common/net"
	"github.com/vdonkey/accelerator/v5/features"
)

// FakeDNSEngine is a Accelerator feature for converting between domain and fake IPs.
//
// accelerator:api:beta
type FakeDNSEngine interface {
	features.Feature

	// GetFakeIPForDomain returns fake IP addresses for the given domain, and registers the domain with the returned IPs.
	GetFakeIPForDomain(domain string) []net.Address

	// GetDomainFromFakeDNS returns the bound domain name for the given fake IP.
	GetDomainFromFakeDNS(ip net.Address) string
}

// FakeDNSEngineRev0 adds additional APIs for FakeDNSEngine.
//
// accelerator:api:beta
type FakeDNSEngineRev0 interface {
	FakeDNSEngine

	// IsIPInIPPool tests whether the given IP address resides in managed fake IP pools.
	IsIPInIPPool(ip net.Address) bool

	// GetFakeIPForDomain3 registers and returns fake IP addresses for the given domain in IPv4 and/or IPv6.
	GetFakeIPForDomain3(domain string, IPv4 bool, IPv6 bool) []net.Address
}

// ClientWithFakeDNS is an optional feature for utilizing FakeDNS feature of DNS client.
//
// accelerator:api:beta
type ClientWithFakeDNS interface {
	// AsFakeDNSClient converts the client to dns.Client that enables FakeDNS querying option.
	AsFakeDNSClient() Client

	// AsFakeDNSEngine converts the client to dns.FakeDNSEngine that could serve FakeDNSEngine feature.
	AsFakeDNSEngine() FakeDNSEngine
}

// FakeDNSEngineType returns the type of FakeDNSEngine interface. Can be used for implementing common.HasType.
//
// accelerator:api:beta
func FakeDNSEngineType() interface{} {
	return (*FakeDNSEngine)(nil)
}
