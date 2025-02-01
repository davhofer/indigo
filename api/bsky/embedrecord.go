// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.embed.record

import (
	"encoding/json"
	"fmt"

	comatprototypes "github.com/davhofer/indigo/api/atproto"
	"github.com/davhofer/indigo/lex/util"
)

func init() {
	util.RegisterType("app.bsky.embed.record#main", &EmbedRecord{})
} // EmbedRecord is a "main" in the app.bsky.embed.record schema.
// RECORDTYPE: EmbedRecord
type EmbedRecord struct {
	LexiconTypeID string                         `json:"$type,const=app.bsky.embed.record" cborgen:"$type,const=app.bsky.embed.record"`
	Record        *comatprototypes.RepoStrongRef `json:"record" cborgen:"record"`
}

// EmbedRecord_View is a "view" in the app.bsky.embed.record schema.
//
// RECORDTYPE: EmbedRecord_View
type EmbedRecord_View struct {
	LexiconTypeID string                   `json:"$type,const=app.bsky.embed.record#view" cborgen:"$type,const=app.bsky.embed.record#view"`
	Record        *EmbedRecord_View_Record `json:"record" cborgen:"record"`
}

// EmbedRecord_ViewBlocked is a "viewBlocked" in the app.bsky.embed.record schema.
//
// RECORDTYPE: EmbedRecord_ViewBlocked
type EmbedRecord_ViewBlocked struct {
	LexiconTypeID string                  `json:"$type,const=app.bsky.embed.record#viewBlocked" cborgen:"$type,const=app.bsky.embed.record#viewBlocked"`
	Author        *FeedDefs_BlockedAuthor `json:"author" cborgen:"author"`
	Blocked       bool                    `json:"blocked" cborgen:"blocked"`
	Uri           string                  `json:"uri" cborgen:"uri"`
}

// EmbedRecord_ViewDetached is a "viewDetached" in the app.bsky.embed.record schema.
//
// RECORDTYPE: EmbedRecord_ViewDetached
type EmbedRecord_ViewDetached struct {
	LexiconTypeID string `json:"$type,const=app.bsky.embed.record#viewDetached" cborgen:"$type,const=app.bsky.embed.record#viewDetached"`
	Detached      bool   `json:"detached" cborgen:"detached"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// EmbedRecord_ViewNotFound is a "viewNotFound" in the app.bsky.embed.record schema.
//
// RECORDTYPE: EmbedRecord_ViewNotFound
type EmbedRecord_ViewNotFound struct {
	LexiconTypeID string `json:"$type,const=app.bsky.embed.record#viewNotFound" cborgen:"$type,const=app.bsky.embed.record#viewNotFound"`
	NotFound      bool   `json:"notFound" cborgen:"notFound"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// EmbedRecord_ViewRecord is a "viewRecord" in the app.bsky.embed.record schema.
//
// RECORDTYPE: EmbedRecord_ViewRecord
type EmbedRecord_ViewRecord struct {
	LexiconTypeID string                                `json:"$type,const=app.bsky.embed.record#viewRecord" cborgen:"$type,const=app.bsky.embed.record#viewRecord"`
	Author        *ActorDefs_ProfileViewBasic           `json:"author" cborgen:"author"`
	Cid           string                                `json:"cid" cborgen:"cid"`
	Embeds        []*EmbedRecord_ViewRecord_Embeds_Elem `json:"embeds,omitempty" cborgen:"embeds,omitempty"`
	IndexedAt     string                                `json:"indexedAt" cborgen:"indexedAt"`
	Labels        []*comatprototypes.LabelDefs_Label    `json:"labels,omitempty" cborgen:"labels,omitempty"`
	LikeCount     *int64                                `json:"likeCount,omitempty" cborgen:"likeCount,omitempty"`
	QuoteCount    *int64                                `json:"quoteCount,omitempty" cborgen:"quoteCount,omitempty"`
	ReplyCount    *int64                                `json:"replyCount,omitempty" cborgen:"replyCount,omitempty"`
	RepostCount   *int64                                `json:"repostCount,omitempty" cborgen:"repostCount,omitempty"`
	Uri           string                                `json:"uri" cborgen:"uri"`
	// value: The record data itself.
	Value *util.LexiconTypeDecoder `json:"value" cborgen:"value"`
}

type EmbedRecord_ViewRecord_Embeds_Elem struct {
	EmbedImages_View          *EmbedImages_View
	EmbedVideo_View           *EmbedVideo_View
	EmbedExternal_View        *EmbedExternal_View
	EmbedRecord_View          *EmbedRecord_View
	EmbedRecordWithMedia_View *EmbedRecordWithMedia_View
}

func (t *EmbedRecord_ViewRecord_Embeds_Elem) MarshalJSON() ([]byte, error) {
	if t.EmbedImages_View != nil {
		t.EmbedImages_View.LexiconTypeID = "app.bsky.embed.images#view"
		return json.Marshal(t.EmbedImages_View)
	}
	if t.EmbedVideo_View != nil {
		t.EmbedVideo_View.LexiconTypeID = "app.bsky.embed.video#view"
		return json.Marshal(t.EmbedVideo_View)
	}
	if t.EmbedExternal_View != nil {
		t.EmbedExternal_View.LexiconTypeID = "app.bsky.embed.external#view"
		return json.Marshal(t.EmbedExternal_View)
	}
	if t.EmbedRecord_View != nil {
		t.EmbedRecord_View.LexiconTypeID = "app.bsky.embed.record#view"
		return json.Marshal(t.EmbedRecord_View)
	}
	if t.EmbedRecordWithMedia_View != nil {
		t.EmbedRecordWithMedia_View.LexiconTypeID = "app.bsky.embed.recordWithMedia#view"
		return json.Marshal(t.EmbedRecordWithMedia_View)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecord_ViewRecord_Embeds_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.images#view":
		t.EmbedImages_View = new(EmbedImages_View)
		return json.Unmarshal(b, t.EmbedImages_View)
	case "app.bsky.embed.video#view":
		t.EmbedVideo_View = new(EmbedVideo_View)
		return json.Unmarshal(b, t.EmbedVideo_View)
	case "app.bsky.embed.external#view":
		t.EmbedExternal_View = new(EmbedExternal_View)
		return json.Unmarshal(b, t.EmbedExternal_View)
	case "app.bsky.embed.record#view":
		t.EmbedRecord_View = new(EmbedRecord_View)
		return json.Unmarshal(b, t.EmbedRecord_View)
	case "app.bsky.embed.recordWithMedia#view":
		t.EmbedRecordWithMedia_View = new(EmbedRecordWithMedia_View)
		return json.Unmarshal(b, t.EmbedRecordWithMedia_View)

	default:
		return nil
	}
}

type EmbedRecord_View_Record struct {
	EmbedRecord_ViewRecord         *EmbedRecord_ViewRecord
	EmbedRecord_ViewNotFound       *EmbedRecord_ViewNotFound
	EmbedRecord_ViewBlocked        *EmbedRecord_ViewBlocked
	EmbedRecord_ViewDetached       *EmbedRecord_ViewDetached
	FeedDefs_GeneratorView         *FeedDefs_GeneratorView
	GraphDefs_ListView             *GraphDefs_ListView
	LabelerDefs_LabelerView        *LabelerDefs_LabelerView
	GraphDefs_StarterPackViewBasic *GraphDefs_StarterPackViewBasic
}

func (t *EmbedRecord_View_Record) MarshalJSON() ([]byte, error) {
	if t.EmbedRecord_ViewRecord != nil {
		t.EmbedRecord_ViewRecord.LexiconTypeID = "app.bsky.embed.record#viewRecord"
		return json.Marshal(t.EmbedRecord_ViewRecord)
	}
	if t.EmbedRecord_ViewNotFound != nil {
		t.EmbedRecord_ViewNotFound.LexiconTypeID = "app.bsky.embed.record#viewNotFound"
		return json.Marshal(t.EmbedRecord_ViewNotFound)
	}
	if t.EmbedRecord_ViewBlocked != nil {
		t.EmbedRecord_ViewBlocked.LexiconTypeID = "app.bsky.embed.record#viewBlocked"
		return json.Marshal(t.EmbedRecord_ViewBlocked)
	}
	if t.EmbedRecord_ViewDetached != nil {
		t.EmbedRecord_ViewDetached.LexiconTypeID = "app.bsky.embed.record#viewDetached"
		return json.Marshal(t.EmbedRecord_ViewDetached)
	}
	if t.FeedDefs_GeneratorView != nil {
		t.FeedDefs_GeneratorView.LexiconTypeID = "app.bsky.feed.defs#generatorView"
		return json.Marshal(t.FeedDefs_GeneratorView)
	}
	if t.GraphDefs_ListView != nil {
		t.GraphDefs_ListView.LexiconTypeID = "app.bsky.graph.defs#listView"
		return json.Marshal(t.GraphDefs_ListView)
	}
	if t.LabelerDefs_LabelerView != nil {
		t.LabelerDefs_LabelerView.LexiconTypeID = "app.bsky.labeler.defs#labelerView"
		return json.Marshal(t.LabelerDefs_LabelerView)
	}
	if t.GraphDefs_StarterPackViewBasic != nil {
		t.GraphDefs_StarterPackViewBasic.LexiconTypeID = "app.bsky.graph.defs#starterPackViewBasic"
		return json.Marshal(t.GraphDefs_StarterPackViewBasic)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecord_View_Record) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.record#viewRecord":
		t.EmbedRecord_ViewRecord = new(EmbedRecord_ViewRecord)
		return json.Unmarshal(b, t.EmbedRecord_ViewRecord)
	case "app.bsky.embed.record#viewNotFound":
		t.EmbedRecord_ViewNotFound = new(EmbedRecord_ViewNotFound)
		return json.Unmarshal(b, t.EmbedRecord_ViewNotFound)
	case "app.bsky.embed.record#viewBlocked":
		t.EmbedRecord_ViewBlocked = new(EmbedRecord_ViewBlocked)
		return json.Unmarshal(b, t.EmbedRecord_ViewBlocked)
	case "app.bsky.embed.record#viewDetached":
		t.EmbedRecord_ViewDetached = new(EmbedRecord_ViewDetached)
		return json.Unmarshal(b, t.EmbedRecord_ViewDetached)
	case "app.bsky.feed.defs#generatorView":
		t.FeedDefs_GeneratorView = new(FeedDefs_GeneratorView)
		return json.Unmarshal(b, t.FeedDefs_GeneratorView)
	case "app.bsky.graph.defs#listView":
		t.GraphDefs_ListView = new(GraphDefs_ListView)
		return json.Unmarshal(b, t.GraphDefs_ListView)
	case "app.bsky.labeler.defs#labelerView":
		t.LabelerDefs_LabelerView = new(LabelerDefs_LabelerView)
		return json.Unmarshal(b, t.LabelerDefs_LabelerView)
	case "app.bsky.graph.defs#starterPackViewBasic":
		t.GraphDefs_StarterPackViewBasic = new(GraphDefs_StarterPackViewBasic)
		return json.Unmarshal(b, t.GraphDefs_StarterPackViewBasic)

	default:
		return nil
	}
}
