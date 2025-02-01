// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.confirmEmail

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// ServerConfirmEmail_Input is the input argument to a com.atproto.server.confirmEmail call.
type ServerConfirmEmail_Input struct {
	Email string `json:"email" cborgen:"email"`
	Token string `json:"token" cborgen:"token"`
}

// ServerConfirmEmail calls the XRPC method "com.atproto.server.confirmEmail".
func ServerConfirmEmail(ctx context.Context, c *xrpc.Client, input *ServerConfirmEmail_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.server.confirmEmail", nil, input, nil); err != nil {
		return err
	}

	return nil
}
