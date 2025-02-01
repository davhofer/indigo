// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.moderation.getEvent

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// ModerationGetEvent calls the XRPC method "tools.ozone.moderation.getEvent".
func ModerationGetEvent(ctx context.Context, c *xrpc.Client, id int64) (*ModerationDefs_ModEventViewDetail, error) {
	var out ModerationDefs_ModEventViewDetail

	params := map[string]interface{}{
		"id": id,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.moderation.getEvent", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
