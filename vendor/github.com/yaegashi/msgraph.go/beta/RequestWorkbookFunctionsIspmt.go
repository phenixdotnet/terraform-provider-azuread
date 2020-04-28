// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsIspmtRequestBuilder struct{ BaseRequestBuilder }

// Ispmt action undocumented
func (b *WorkbookFunctionsRequestBuilder) Ispmt(reqObj *WorkbookFunctionsIspmtRequestParameter) *WorkbookFunctionsIspmtRequestBuilder {
	bb := &WorkbookFunctionsIspmtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ispmt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsIspmtRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsIspmtRequestBuilder) Request() *WorkbookFunctionsIspmtRequest {
	return &WorkbookFunctionsIspmtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsIspmtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
