// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operator_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// GetMPIntegrationOKCode is the HTTP code returned for type GetMPIntegrationOK
const GetMPIntegrationOKCode int = 200

/*GetMPIntegrationOK A successful response.

swagger:response getMPIntegrationOK
*/
type GetMPIntegrationOK struct {

	/*
	  In: Body
	*/
	Payload *GetMPIntegrationOKBody `json:"body,omitempty"`
}

// NewGetMPIntegrationOK creates GetMPIntegrationOK with default headers values
func NewGetMPIntegrationOK() *GetMPIntegrationOK {

	return &GetMPIntegrationOK{}
}

// WithPayload adds the payload to the get m p integration o k response
func (o *GetMPIntegrationOK) WithPayload(payload *GetMPIntegrationOKBody) *GetMPIntegrationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get m p integration o k response
func (o *GetMPIntegrationOK) SetPayload(payload *GetMPIntegrationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMPIntegrationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetMPIntegrationDefault Generic error response.

swagger:response getMPIntegrationDefault
*/
type GetMPIntegrationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMPIntegrationDefault creates GetMPIntegrationDefault with default headers values
func NewGetMPIntegrationDefault(code int) *GetMPIntegrationDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMPIntegrationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get m p integration default response
func (o *GetMPIntegrationDefault) WithStatusCode(code int) *GetMPIntegrationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get m p integration default response
func (o *GetMPIntegrationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get m p integration default response
func (o *GetMPIntegrationDefault) WithPayload(payload *models.Error) *GetMPIntegrationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get m p integration default response
func (o *GetMPIntegrationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMPIntegrationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
