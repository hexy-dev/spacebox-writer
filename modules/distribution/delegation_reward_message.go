package distribution

import (
	"context"

	"github.com/bro-n-bro/spacebox-writer/internal/rep"
	"github.com/bro-n-bro/spacebox-writer/modules/utils"
	"github.com/bro-n-bro/spacebox/broker/model"
)

// DelegationRewardMessageHandler is a handler for delegation reward message event
func DelegationRewardMessageHandler(ctx context.Context, msgs [][]byte, ch rep.Storage) error {
	vals, err := utils.ConvertMessages[model.DelegationRewardMessage](msgs)
	if err != nil {
		return err
	}
	return ch.DelegationRewardMessage(vals)
}
