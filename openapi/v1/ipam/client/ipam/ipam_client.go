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
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ipam API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteIpamIps(params *DeleteIpamIpsParams, opts ...ClientOption) (*DeleteIpamIpsOK, error)

	GetIpamHealthz(params *GetIpamHealthzParams, opts ...ClientOption) (*GetIpamHealthzOK, error)

	PostIpamIps(params *PostIpamIpsParams, opts ...ClientOption) (*PostIpamIpsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteIpamIps delete ipam ips API
*/
func (a *Client) DeleteIpamIps(params *DeleteIpamIpsParams, opts ...ClientOption) (*DeleteIpamIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIpamIpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteIpamIps",
		Method:             "DELETE",
		PathPattern:        "/ipam/ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"unix"},
		Params:             params,
		Reader:             &DeleteIpamIpsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIpamIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteIpamIps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIpamHealthz get ipam healthz API
*/
func (a *Client) GetIpamHealthz(params *GetIpamHealthzParams, opts ...ClientOption) (*GetIpamHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIpamHealthzParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIpamHealthz",
		Method:             "GET",
		PathPattern:        "/ipam/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"unix"},
		Params:             params,
		Reader:             &GetIpamHealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIpamHealthzOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIpamHealthz: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostIpamIps post ipam ips API
*/
func (a *Client) PostIpamIps(params *PostIpamIpsParams, opts ...ClientOption) (*PostIpamIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIpamIpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostIpamIps",
		Method:             "POST",
		PathPattern:        "/ipam/ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"unix"},
		Params:             params,
		Reader:             &PostIpamIpsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIpamIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostIpamIps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
