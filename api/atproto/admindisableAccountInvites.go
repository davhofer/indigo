// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.disableAccountInvites

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// AdminDisableAccountInvites_Input is the input argument to a com.atproto.admin.disableAccountInvites call.
type AdminDisableAccountInvites_Input struct {
	Account string `json:"account" cborgen:"account"`
	// note: Optional reason for disabled invites.
	Note *string `json:"note,omitempty" cborgen:"note,omitempty"`
}

// AdminDisableAccountInvites calls the XRPC method "com.atproto.admin.disableAccountInvites".
func AdminDisableAccountInvites(ctx context.Context, c *xrpc.Client, input *AdminDisableAccountInvites_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.disableAccountInvites", nil, input, nil); err != nil {
		return err
	}

	return nil
}
