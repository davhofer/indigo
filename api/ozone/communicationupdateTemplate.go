// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.communication.updateTemplate

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// CommunicationUpdateTemplate_Input is the input argument to a tools.ozone.communication.updateTemplate call.
type CommunicationUpdateTemplate_Input struct {
	// contentMarkdown: Content of the template, markdown supported, can contain variable placeholders.
	ContentMarkdown *string `json:"contentMarkdown,omitempty" cborgen:"contentMarkdown,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty" cborgen:"disabled,omitempty"`
	// id: ID of the template to be updated.
	Id string `json:"id" cborgen:"id"`
	// lang: Message language.
	Lang *string `json:"lang,omitempty" cborgen:"lang,omitempty"`
	// name: Name of the template.
	Name *string `json:"name,omitempty" cborgen:"name,omitempty"`
	// subject: Subject of the message, used in emails.
	Subject *string `json:"subject,omitempty" cborgen:"subject,omitempty"`
	// updatedBy: DID of the user who is updating the template.
	UpdatedBy *string `json:"updatedBy,omitempty" cborgen:"updatedBy,omitempty"`
}

// CommunicationUpdateTemplate calls the XRPC method "tools.ozone.communication.updateTemplate".
func CommunicationUpdateTemplate(ctx context.Context, c *xrpc.Client, input *CommunicationUpdateTemplate_Input) (*CommunicationDefs_TemplateView, error) {
	var out CommunicationDefs_TemplateView
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "tools.ozone.communication.updateTemplate", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
