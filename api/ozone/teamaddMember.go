// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.team.addMember

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// TeamAddMember_Input is the input argument to a tools.ozone.team.addMember call.
type TeamAddMember_Input struct {
	Did  string `json:"did" cborgen:"did"`
	Role string `json:"role" cborgen:"role"`
}

// TeamAddMember calls the XRPC method "tools.ozone.team.addMember".
func TeamAddMember(ctx context.Context, c *xrpc.Client, input *TeamAddMember_Input) (*TeamDefs_Member, error) {
	var out TeamDefs_Member
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "tools.ozone.team.addMember", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
