//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

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
)

// NewSchemaClusterStatusParams creates a new SchemaClusterStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchemaClusterStatusParams() *SchemaClusterStatusParams {
	return &SchemaClusterStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaClusterStatusParamsWithTimeout creates a new SchemaClusterStatusParams object
// with the ability to set a timeout on a request.
func NewSchemaClusterStatusParamsWithTimeout(timeout time.Duration) *SchemaClusterStatusParams {
	return &SchemaClusterStatusParams{
		timeout: timeout,
	}
}

// NewSchemaClusterStatusParamsWithContext creates a new SchemaClusterStatusParams object
// with the ability to set a context for a request.
func NewSchemaClusterStatusParamsWithContext(ctx context.Context) *SchemaClusterStatusParams {
	return &SchemaClusterStatusParams{
		Context: ctx,
	}
}

// NewSchemaClusterStatusParamsWithHTTPClient creates a new SchemaClusterStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchemaClusterStatusParamsWithHTTPClient(client *http.Client) *SchemaClusterStatusParams {
	return &SchemaClusterStatusParams{
		HTTPClient: client,
	}
}

/*
SchemaClusterStatusParams contains all the parameters to send to the API endpoint

	for the schema cluster status operation.

	Typically these are written to a http.Request.
*/
type SchemaClusterStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schema cluster status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaClusterStatusParams) WithDefaults() *SchemaClusterStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schema cluster status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaClusterStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schema cluster status params
func (o *SchemaClusterStatusParams) WithTimeout(timeout time.Duration) *SchemaClusterStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema cluster status params
func (o *SchemaClusterStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema cluster status params
func (o *SchemaClusterStatusParams) WithContext(ctx context.Context) *SchemaClusterStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema cluster status params
func (o *SchemaClusterStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema cluster status params
func (o *SchemaClusterStatusParams) WithHTTPClient(client *http.Client) *SchemaClusterStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema cluster status params
func (o *SchemaClusterStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaClusterStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
