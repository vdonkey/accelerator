package entries

import "github.com/vdonkey/accelerator/v5/app/subscription/specs"

//go:generate go run github.com/vdonkey/accelerator/v5/common/errors/errorgen

type Converter interface {
	ConvertToAbstractServerConfig(rawConfig []byte, kindHint string) (*specs.SubscriptionServerConfig, error)
}
