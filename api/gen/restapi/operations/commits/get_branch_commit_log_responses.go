// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/treeverse/lakefs/api/gen/models"
)

// GetBranchCommitLogOKCode is the HTTP code returned for type GetBranchCommitLogOK
const GetBranchCommitLogOKCode int = 200

/*GetBranchCommitLogOK commit log

swagger:response getBranchCommitLogOK
*/
type GetBranchCommitLogOK struct {

	/*
	  In: Body
	*/
	Payload *GetBranchCommitLogOKBody `json:"body,omitempty"`
}

// NewGetBranchCommitLogOK creates GetBranchCommitLogOK with default headers values
func NewGetBranchCommitLogOK() *GetBranchCommitLogOK {

	return &GetBranchCommitLogOK{}
}

// WithPayload adds the payload to the get branch commit log o k response
func (o *GetBranchCommitLogOK) WithPayload(payload *GetBranchCommitLogOKBody) *GetBranchCommitLogOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get branch commit log o k response
func (o *GetBranchCommitLogOK) SetPayload(payload *GetBranchCommitLogOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBranchCommitLogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBranchCommitLogUnauthorizedCode is the HTTP code returned for type GetBranchCommitLogUnauthorized
const GetBranchCommitLogUnauthorizedCode int = 401

/*GetBranchCommitLogUnauthorized Unauthorized

swagger:response getBranchCommitLogUnauthorized
*/
type GetBranchCommitLogUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetBranchCommitLogUnauthorized creates GetBranchCommitLogUnauthorized with default headers values
func NewGetBranchCommitLogUnauthorized() *GetBranchCommitLogUnauthorized {

	return &GetBranchCommitLogUnauthorized{}
}

// WithPayload adds the payload to the get branch commit log unauthorized response
func (o *GetBranchCommitLogUnauthorized) WithPayload(payload interface{}) *GetBranchCommitLogUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get branch commit log unauthorized response
func (o *GetBranchCommitLogUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBranchCommitLogUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBranchCommitLogNotFoundCode is the HTTP code returned for type GetBranchCommitLogNotFound
const GetBranchCommitLogNotFoundCode int = 404

/*GetBranchCommitLogNotFound branch not found

swagger:response getBranchCommitLogNotFound
*/
type GetBranchCommitLogNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBranchCommitLogNotFound creates GetBranchCommitLogNotFound with default headers values
func NewGetBranchCommitLogNotFound() *GetBranchCommitLogNotFound {

	return &GetBranchCommitLogNotFound{}
}

// WithPayload adds the payload to the get branch commit log not found response
func (o *GetBranchCommitLogNotFound) WithPayload(payload *models.Error) *GetBranchCommitLogNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get branch commit log not found response
func (o *GetBranchCommitLogNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBranchCommitLogNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBranchCommitLogDefault generic error response

swagger:response getBranchCommitLogDefault
*/
type GetBranchCommitLogDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBranchCommitLogDefault creates GetBranchCommitLogDefault with default headers values
func NewGetBranchCommitLogDefault(code int) *GetBranchCommitLogDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBranchCommitLogDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get branch commit log default response
func (o *GetBranchCommitLogDefault) WithStatusCode(code int) *GetBranchCommitLogDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get branch commit log default response
func (o *GetBranchCommitLogDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get branch commit log default response
func (o *GetBranchCommitLogDefault) WithPayload(payload *models.Error) *GetBranchCommitLogDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get branch commit log default response
func (o *GetBranchCommitLogDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBranchCommitLogDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}