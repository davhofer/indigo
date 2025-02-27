// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.unspecced.applyLabels

import (
	"context"

	comatprototypes "github.com/davhofer/indigo/api/atproto"
	"github.com/davhofer/indigo/xrpc"
)

// UnspeccedApplyLabels_Input is the input argument to a app.bsky.unspecced.applyLabels call.
type UnspeccedApplyLabels_Input struct {
	Labels []*comatprototypes.LabelDefs_Label `json:"labels" cborgen:"labels"`
}

// UnspeccedApplyLabels calls the XRPC method "app.bsky.unspecced.applyLabels".
func UnspeccedApplyLabels(ctx context.Context, c *xrpc.Client, input *UnspeccedApplyLabels_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "app.bsky.unspecced.applyLabels", nil, input, nil); err != nil {
		return err
	}

	return nil
}
