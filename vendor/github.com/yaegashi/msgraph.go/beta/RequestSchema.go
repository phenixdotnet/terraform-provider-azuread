// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SchemaRequestBuilder is request builder for Schema
type SchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns SchemaRequest
func (b *SchemaRequestBuilder) Request() *SchemaRequest {
	return &SchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SchemaRequest is request for Schema
type SchemaRequest struct{ BaseRequest }

// Get performs GET request for Schema
func (r *SchemaRequest) Get(ctx context.Context) (resObj *Schema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Schema
func (r *SchemaRequest) Update(ctx context.Context, reqObj *Schema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Schema
func (r *SchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SchemaExtensionRequestBuilder is request builder for SchemaExtension
type SchemaExtensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SchemaExtensionRequest
func (b *SchemaExtensionRequestBuilder) Request() *SchemaExtensionRequest {
	return &SchemaExtensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SchemaExtensionRequest is request for SchemaExtension
type SchemaExtensionRequest struct{ BaseRequest }

// Get performs GET request for SchemaExtension
func (r *SchemaExtensionRequest) Get(ctx context.Context) (resObj *SchemaExtension, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SchemaExtension
func (r *SchemaExtensionRequest) Update(ctx context.Context, reqObj *SchemaExtension) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SchemaExtension
func (r *SchemaExtensionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
