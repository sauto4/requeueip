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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteIpamIpsHandlerFunc turns a function with the right signature into a delete ipam ips handler
type DeleteIpamIpsHandlerFunc func(DeleteIpamIpsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteIpamIpsHandlerFunc) Handle(params DeleteIpamIpsParams) middleware.Responder {
	return fn(params)
}

// DeleteIpamIpsHandler interface for that can handle valid delete ipam ips params
type DeleteIpamIpsHandler interface {
	Handle(DeleteIpamIpsParams) middleware.Responder
}

// NewDeleteIpamIps creates a new http.Handler for the delete ipam ips operation
func NewDeleteIpamIps(ctx *middleware.Context, handler DeleteIpamIpsHandler) *DeleteIpamIps {
	return &DeleteIpamIps{Context: ctx, Handler: handler}
}

/*
	DeleteIpamIps swagger:route DELETE /ipam/ips ipam deleteIpamIps

DeleteIpamIps delete ipam ips API
*/
type DeleteIpamIps struct {
	Context *middleware.Context
	Handler DeleteIpamIpsHandler
}

func (o *DeleteIpamIps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteIpamIpsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
