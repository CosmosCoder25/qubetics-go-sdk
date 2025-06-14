package core

import (
	"context"
	"fmt"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/types"
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"

	"github.com/qubetics/qubetics-go-sdk/utils"
)

// NodeStartSession initiates a new session on a specified node. On success, it returns the session ID.
func (c *Client) NodeStartSession(ctx context.Context, nodeAddr types.NodeAddress, gigabytes, hours int64, denom string) (uint64, error) {
	// Retrieve the message from address.
	fromAddr, err := c.MsgFromAddr()
	if err != nil {
		return 0, fmt.Errorf("failed to get message from addr: %w", err)
	}

	// Construct the session start request message for a node session.
	msgs := []cosmossdk.Msg{
		v3.NewMsgStartSessionRequest(fromAddr, nodeAddr, gigabytes, hours, denom),
	}

	// Broadcast the transaction and wait for its inclusion in a block.
	_, res, err := c.BroadcastTxBlock(ctx, msgs...)
	if err != nil {
		return 0, fmt.Errorf("node start session tx failed: %w", err)
	}

	// Extract and return the session ID from the transaction events.
	id, err := utils.IDFromEvents(res.TxResult.GetEvents(), &v3.EventCreateSession{})
	if err != nil {
		return 0, fmt.Errorf("failed to get id from events: %w", err)
	}

	return id, nil
}
