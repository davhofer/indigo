// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.disableInviteCodes

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// AdminDisableInviteCodes_Input is the input argument to a com.atproto.admin.disableInviteCodes call.
type AdminDisableInviteCodes_Input struct {
	Accounts []string `json:"accounts,omitempty" cborgen:"accounts,omitempty"`
	Codes    []string `json:"codes,omitempty" cborgen:"codes,omitempty"`
}

// AdminDisableInviteCodes calls the XRPC method "com.atproto.admin.disableInviteCodes".
func AdminDisableInviteCodes(ctx context.Context, c *xrpc.Client, input *AdminDisableInviteCodes_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.disableInviteCodes", nil, input, nil); err != nil {
		return err
	}

	return nil
}
