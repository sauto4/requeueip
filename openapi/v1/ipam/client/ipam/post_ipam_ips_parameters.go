// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 The RequeueIP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/sauto4/requeueip/openapi/v1/ipam/models"
)

// NewPostIpamIpsParams creates a new PostIpamIpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostIpamIpsParams() *PostIpamIpsParams {
	return &PostIpamIpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpamIpsParamsWithTimeout creates a new PostIpamIpsParams object
// with the ability to set a timeout on a request.
func NewPostIpamIpsParamsWithTimeout(timeout time.Duration) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		timeout: timeout,
	}
}

// NewPostIpamIpsParamsWithContext creates a new PostIpamIpsParams object
// with the ability to set a context for a request.
func NewPostIpamIpsParamsWithContext(ctx context.Context) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		Context: ctx,
	}
}

// NewPostIpamIpsParamsWithHTTPClient creates a new PostIpamIpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostIpamIpsParamsWithHTTPClient(client *http.Client) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		HTTPClient: client,
	}
}

/*
PostIpamIpsParams contains all the parameters to send to the API endpoint

	for the post ipam ips operation.

	Typically these are written to a http.Request.
*/
type PostIpamIpsParams struct {

	// Args.
	Args *models.CmdAddArgs

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post ipam ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamIpsParams) WithDefaults() *PostIpamIpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ipam ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamIpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post ipam ips params
func (o *PostIpamIpsParams) WithTimeout(timeout time.Duration) *PostIpamIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipam ips params
func (o *PostIpamIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipam ips params
func (o *PostIpamIpsParams) WithContext(ctx context.Context) *PostIpamIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipam ips params
func (o *PostIpamIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipam ips params
func (o *PostIpamIpsParams) WithHTTPClient(client *http.Client) *PostIpamIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipam ips params
func (o *PostIpamIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the post ipam ips params
func (o *PostIpamIpsParams) WithArgs(args *models.CmdAddArgs) *PostIpamIpsParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the post ipam ips params
func (o *PostIpamIpsParams) SetArgs(args *models.CmdAddArgs) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpamIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
