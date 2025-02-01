// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.team.deleteMember

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// TeamDeleteMember_Input is the input argument to a tools.ozone.team.deleteMember call.
type TeamDeleteMember_Input struct {
	Did string `json:"did" cborgen:"did"`
}

// TeamDeleteMember calls the XRPC method "tools.ozone.team.deleteMember".
func TeamDeleteMember(ctx context.Context, c *xrpc.Client, input *TeamDeleteMember_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "tools.ozone.team.deleteMember", nil, input, nil); err != nil {
		return err
	}

	return nil
}
