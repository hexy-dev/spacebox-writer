package staking

import (
	"context"
	"encoding/json"

	"github.com/hexy-dev/spacebox-writer/internal/rep"
	"github.com/hexy-dev/spacebox/broker/model"
)

func RedelegationMessageHandler(ctx context.Context, msg []byte, ch rep.Storage) error {
	val := model.RedelegationMessage{}
	if err := json.Unmarshal(msg, &val); err != nil {
		return err
	}

	return ch.RedelegationMessage(val)
}
