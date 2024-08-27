package wechat_test

import (
	"context"
	"testing"

	"github.com/vdonkey/accelerator/v5/common"
	"github.com/vdonkey/accelerator/v5/common/buf"
	. "github.com/vdonkey/accelerator/v5/transport/internet/headers/wechat"
)

func TestUTPWrite(t *testing.T) {
	videoRaw, err := NewVideoChat(context.Background(), &VideoConfig{})
	common.Must(err)

	video := videoRaw.(*VideoChat)

	payload := buf.New()
	video.Serialize(payload.Extend(video.Size()))

	if payload.Len() != video.Size() {
		t.Error("expected payload size ", video.Size(), " but got ", payload.Len())
	}
}
