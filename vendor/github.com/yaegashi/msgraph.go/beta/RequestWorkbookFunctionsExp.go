// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsExpRequestBuilder struct{ BaseRequestBuilder }

// Exp action undocumented
func (b *WorkbookFunctionsRequestBuilder) Exp(reqObj *WorkbookFunctionsExpRequestParameter) *WorkbookFunctionsExpRequestBuilder {
	bb := &WorkbookFunctionsExpRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/exp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsExpRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsExpRequestBuilder) Request() *WorkbookFunctionsExpRequest {
	return &WorkbookFunctionsExpRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsExpRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
