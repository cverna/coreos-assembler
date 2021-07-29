// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
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

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// CreateNamespaceCreatedCode is the HTTP code returned for type CreateNamespaceCreated
const CreateNamespaceCreatedCode int = 201

/*CreateNamespaceCreated A successful response.

swagger:response createNamespaceCreated
*/
type CreateNamespaceCreated struct {
}

// NewCreateNamespaceCreated creates CreateNamespaceCreated with default headers values
func NewCreateNamespaceCreated() *CreateNamespaceCreated {

	return &CreateNamespaceCreated{}
}

// WriteResponse to the client
func (o *CreateNamespaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*CreateNamespaceDefault Generic error response.

swagger:response createNamespaceDefault
*/
type CreateNamespaceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateNamespaceDefault creates CreateNamespaceDefault with default headers values
func NewCreateNamespaceDefault(code int) *CreateNamespaceDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateNamespaceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create namespace default response
func (o *CreateNamespaceDefault) WithStatusCode(code int) *CreateNamespaceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create namespace default response
func (o *CreateNamespaceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create namespace default response
func (o *CreateNamespaceDefault) WithPayload(payload *models.Error) *CreateNamespaceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create namespace default response
func (o *CreateNamespaceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNamespaceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}