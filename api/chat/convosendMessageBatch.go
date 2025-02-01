// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package chat

// schema: chat.bsky.convo.sendMessageBatch

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// ConvoSendMessageBatch_BatchItem is a "batchItem" in the chat.bsky.convo.sendMessageBatch schema.
type ConvoSendMessageBatch_BatchItem struct {
	ConvoId string                  `json:"convoId" cborgen:"convoId"`
	Message *ConvoDefs_MessageInput `json:"message" cborgen:"message"`
}

// ConvoSendMessageBatch_Input is the input argument to a chat.bsky.convo.sendMessageBatch call.
type ConvoSendMessageBatch_Input struct {
	Items []*ConvoSendMessageBatch_BatchItem `json:"items" cborgen:"items"`
}

// ConvoSendMessageBatch_Output is the output of a chat.bsky.convo.sendMessageBatch call.
type ConvoSendMessageBatch_Output struct {
	Items []*ConvoDefs_MessageView `json:"items" cborgen:"items"`
}

// ConvoSendMessageBatch calls the XRPC method "chat.bsky.convo.sendMessageBatch".
func ConvoSendMessageBatch(ctx context.Context, c *xrpc.Client, input *ConvoSendMessageBatch_Input) (*ConvoSendMessageBatch_Output, error) {
	var out ConvoSendMessageBatch_Output
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "chat.bsky.convo.sendMessageBatch", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
