package distribution

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/hexy-dev/spacebox-writer/adapter/clickhouse"
	"github.com/hexy-dev/spacebox/broker/model"
)

func CommunityPoolHandler(ctx context.Context, msg []byte, ch rep.Storage) error {
	val := model.CommunityPool{}
	if err := jsoniter.Unmarshal(msg, &val); err != nil {
		return errors.Wrap(err, "unmarshall error")
	}

	return ch.CommunityPool(val)
}
