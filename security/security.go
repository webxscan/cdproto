// Package security provides the Chrome DevTools Protocol
// commands, types, and events for the Security domain.
//
// Security.
//
// Generated by the cdproto-gen command.
package security

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// DisableParams disables tracking security state changes.
type DisableParams struct{}

// Disable disables tracking security state changes.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Security.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables tracking security state changes.
type EnableParams struct{}

// Enable enables tracking security state changes.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Security.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// SetIgnoreCertificateErrorsParams enable/disable whether all certificate
// errors should be ignored.
type SetIgnoreCertificateErrorsParams struct {
	Ignore bool `json:"ignore"` // If true, all certificate errors will be ignored.
}

// SetIgnoreCertificateErrors enable/disable whether all certificate errors
// should be ignored.
//
// parameters:
//   ignore - If true, all certificate errors will be ignored.
func SetIgnoreCertificateErrors(ignore bool) *SetIgnoreCertificateErrorsParams {
	return &SetIgnoreCertificateErrorsParams{
		Ignore: ignore,
	}
}

// Do executes Security.setIgnoreCertificateErrors against the provided context.
func (p *SetIgnoreCertificateErrorsParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandSetIgnoreCertificateErrors, p, nil)
}

// Command names.
const (
	CommandDisable                    = "Security.disable"
	CommandEnable                     = "Security.enable"
	CommandSetIgnoreCertificateErrors = "Security.setIgnoreCertificateErrors"
)
