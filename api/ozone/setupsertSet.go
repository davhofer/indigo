// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.set.upsertSet

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// SetUpsertSet calls the XRPC method "tools.ozone.set.upsertSet".
func SetUpsertSet(ctx context.Context, c *xrpc.Client, input *SetDefs_Set) (*SetDefs_SetView, error) {
	var out SetDefs_SetView
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "tools.ozone.set.upsertSet", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
