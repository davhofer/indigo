// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.moderation.getRecords

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/davhofer/indigo/lex/util"
	"github.com/davhofer/indigo/xrpc"
)

// ModerationGetRecords_Output is the output of a tools.ozone.moderation.getRecords call.
type ModerationGetRecords_Output struct {
	Records []*ModerationGetRecords_Output_Records_Elem `json:"records" cborgen:"records"`
}

type ModerationGetRecords_Output_Records_Elem struct {
	ModerationDefs_RecordViewDetail   *ModerationDefs_RecordViewDetail
	ModerationDefs_RecordViewNotFound *ModerationDefs_RecordViewNotFound
}

func (t *ModerationGetRecords_Output_Records_Elem) MarshalJSON() ([]byte, error) {
	if t.ModerationDefs_RecordViewDetail != nil {
		t.ModerationDefs_RecordViewDetail.LexiconTypeID = "tools.ozone.moderation.defs#recordViewDetail"
		return json.Marshal(t.ModerationDefs_RecordViewDetail)
	}
	if t.ModerationDefs_RecordViewNotFound != nil {
		t.ModerationDefs_RecordViewNotFound.LexiconTypeID = "tools.ozone.moderation.defs#recordViewNotFound"
		return json.Marshal(t.ModerationDefs_RecordViewNotFound)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ModerationGetRecords_Output_Records_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "tools.ozone.moderation.defs#recordViewDetail":
		t.ModerationDefs_RecordViewDetail = new(ModerationDefs_RecordViewDetail)
		return json.Unmarshal(b, t.ModerationDefs_RecordViewDetail)
	case "tools.ozone.moderation.defs#recordViewNotFound":
		t.ModerationDefs_RecordViewNotFound = new(ModerationDefs_RecordViewNotFound)
		return json.Unmarshal(b, t.ModerationDefs_RecordViewNotFound)

	default:
		return nil
	}
}

// ModerationGetRecords calls the XRPC method "tools.ozone.moderation.getRecords".
func ModerationGetRecords(ctx context.Context, c *xrpc.Client, uris []string) (*ModerationGetRecords_Output, error) {
	var out ModerationGetRecords_Output

	params := map[string]interface{}{
		"uris": uris,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.moderation.getRecords", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
