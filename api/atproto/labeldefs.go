// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.label.defs

import (
	"github.com/davhofer/indigo/lex/util"
)

// LabelDefs_Label is a "label" in the com.atproto.label.defs schema.
//
// Metadata tag on an atproto resource (eg, repo or record).
type LabelDefs_Label struct {
	// cid: Optionally, CID specifying the specific version of 'uri' resource this label applies to.
	Cid *string `json:"cid,omitempty" cborgen:"cid,omitempty"`
	// cts: Timestamp when this label was created.
	Cts string `json:"cts" cborgen:"cts"`
	// exp: Timestamp at which this label expires (no longer applies).
	Exp *string `json:"exp,omitempty" cborgen:"exp,omitempty"`
	// neg: If true, this is a negation label, overwriting a previous label.
	Neg *bool `json:"neg,omitempty" cborgen:"neg,omitempty"`
	// sig: Signature of dag-cbor encoded label.
	Sig util.LexBytes `json:"sig,omitempty" cborgen:"sig,omitempty"`
	// src: DID of the actor who created this label.
	Src string `json:"src" cborgen:"src"`
	// uri: AT URI of the record, repository (account), or other resource that this label applies to.
	Uri string `json:"uri" cborgen:"uri"`
	// val: The short string name of the value or type of this label.
	Val string `json:"val" cborgen:"val"`
	// ver: The AT Protocol version of the label object.
	Ver *int64 `json:"ver,omitempty" cborgen:"ver,omitempty"`
}

// LabelDefs_LabelValueDefinition is a "labelValueDefinition" in the com.atproto.label.defs schema.
//
// Declares a label value and its expected interpretations and behaviors.
type LabelDefs_LabelValueDefinition struct {
	// adultOnly: Does the user need to have adult content enabled in order to configure this label?
	AdultOnly *bool `json:"adultOnly,omitempty" cborgen:"adultOnly,omitempty"`
	// blurs: What should this label hide in the UI, if applied? 'content' hides all of the target; 'media' hides the images/video/audio; 'none' hides nothing.
	Blurs string `json:"blurs" cborgen:"blurs"`
	// defaultSetting: The default setting for this label.
	DefaultSetting *string `json:"defaultSetting,omitempty" cborgen:"defaultSetting,omitempty"`
	// identifier: The value of the label being defined. Must only include lowercase ascii and the '-' character ([a-z-]+).
	Identifier string                                   `json:"identifier" cborgen:"identifier"`
	Locales    []*LabelDefs_LabelValueDefinitionStrings `json:"locales" cborgen:"locales"`
	// severity: How should a client visually convey this label? 'inform' means neutral and informational; 'alert' means negative and warning; 'none' means show nothing.
	Severity string `json:"severity" cborgen:"severity"`
}

// LabelDefs_LabelValueDefinitionStrings is a "labelValueDefinitionStrings" in the com.atproto.label.defs schema.
//
// Strings which describe the label in the UI, localized into a specific language.
type LabelDefs_LabelValueDefinitionStrings struct {
	// description: A longer description of what the label means and why it might be applied.
	Description string `json:"description" cborgen:"description"`
	// lang: The code of the language these strings are written in.
	Lang string `json:"lang" cborgen:"lang"`
	// name: A short human-readable name for the label.
	Name string `json:"name" cborgen:"name"`
}

// LabelDefs_SelfLabel is a "selfLabel" in the com.atproto.label.defs schema.
//
// Metadata tag on an atproto record, published by the author within the record. Note that schemas should use #selfLabels, not #selfLabel.
type LabelDefs_SelfLabel struct {
	// val: The short string name of the value or type of this label.
	Val string `json:"val" cborgen:"val"`
}

// LabelDefs_SelfLabels is a "selfLabels" in the com.atproto.label.defs schema.
//
// Metadata tags on an atproto record, published by the author within the record.
//
// RECORDTYPE: LabelDefs_SelfLabels
type LabelDefs_SelfLabels struct {
	LexiconTypeID string                 `json:"$type,const=com.atproto.label.defs#selfLabels" cborgen:"$type,const=com.atproto.label.defs#selfLabels"`
	Values        []*LabelDefs_SelfLabel `json:"values" cborgen:"values,preservenil"`
}
