package extension

import "github.com/vdonkey/accelerator/v5/features"

type SubscriptionManager interface {
	features.Feature
}

func SubscriptionManagerType() interface{} {
	return (*SubscriptionManager)(nil)
}
